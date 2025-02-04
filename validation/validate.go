package validation

import "regexp"

type ErrorValidation struct {
	Message string
	Field   string
	Tag     string
}

func (ve *ErrorValidation) Error() string {
	return ve.Message
}

func ValidatePassword(password string) error {
	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return &ErrorValidation{
			Message: "Password should contain atleast one lower case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}
	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return &ErrorValidation{
			Message: "Password should contain atleast one upper case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}
	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return err
	}
	if !done {
		return &ErrorValidation{
			Message: "Password should contain atleast one digit",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	//done, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
	//if err != nil {
	//	return err
	//}
	//if !done {
	//	return &ErrorValidation{
	//		Message: "Password should contain atleast one special character",
	//		Field:   "password",
	//		Tag:     "strong_password",
	//	}
	//}
	return nil
}
