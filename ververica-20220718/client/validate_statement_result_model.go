// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateStatementResult interface {
	dara.Model
	String() string
	GoString() string
	SetErrorDetails(v *ValidationErrorDetails) *ValidateStatementResult
	GetErrorDetails() *ValidationErrorDetails
	SetValidationResult(v string) *ValidateStatementResult
	GetValidationResult() *string
}

type ValidateStatementResult struct {
	ErrorDetails *ValidationErrorDetails `json:"errorDetails,omitempty" xml:"errorDetails,omitempty"`
	// example:
	//
	// "there have some errors""
	ValidationResult *string `json:"validationResult,omitempty" xml:"validationResult,omitempty"`
}

func (s ValidateStatementResult) String() string {
	return dara.Prettify(s)
}

func (s ValidateStatementResult) GoString() string {
	return s.String()
}

func (s *ValidateStatementResult) GetErrorDetails() *ValidationErrorDetails {
	return s.ErrorDetails
}

func (s *ValidateStatementResult) GetValidationResult() *string {
	return s.ValidationResult
}

func (s *ValidateStatementResult) SetErrorDetails(v *ValidationErrorDetails) *ValidateStatementResult {
	s.ErrorDetails = v
	return s
}

func (s *ValidateStatementResult) SetValidationResult(v string) *ValidateStatementResult {
	s.ValidationResult = &v
	return s
}

func (s *ValidateStatementResult) Validate() error {
	if s.ErrorDetails != nil {
		if err := s.ErrorDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}
