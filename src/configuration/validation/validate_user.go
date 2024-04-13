package validation

import (
	"encoding/json"
	"errors"

	resterr "github.com/Flavio-coutinho/CRUD-GoLang/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)


var (
	Validate = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validate_err error,
) *resterr.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validate_err, &jsonErr) {
		return resterr.NewBadRequestError("Inavalid field type")
	} else if errors.As(validate_err, &jsonValidationError) {
		errorsCauses := []resterr.Causes{}

		for _, e := range validate_err.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return resterr.NewBadRequestValidationError("Some fields are inavlid", errorsCauses)
	} else {
		return resterr.NewBadRequestError("Error trying to convert fields")
	}

} 