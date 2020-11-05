package tea

// ErrCoalesce returns first non-nil error
//
func ErrCoalesce(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
