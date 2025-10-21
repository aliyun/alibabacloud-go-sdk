// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlStatementValidationResult interface {
	dara.Model
	String() string
	GoString() string
	SetErrorDetails(v *ErrorDetails) *SqlStatementValidationResult
	GetErrorDetails() *ErrorDetails
	SetMessage(v string) *SqlStatementValidationResult
	GetMessage() *string
	SetSuccess(v bool) *SqlStatementValidationResult
	GetSuccess() *bool
	SetValidationResult(v string) *SqlStatementValidationResult
	GetValidationResult() *string
}

type SqlStatementValidationResult struct {
	ErrorDetails     *ErrorDetails `json:"errorDetails,omitempty" xml:"errorDetails,omitempty"`
	Message          *string       `json:"message,omitempty" xml:"message,omitempty"`
	Success          *bool         `json:"success,omitempty" xml:"success,omitempty"`
	ValidationResult *string       `json:"validationResult,omitempty" xml:"validationResult,omitempty"`
}

func (s SqlStatementValidationResult) String() string {
	return dara.Prettify(s)
}

func (s SqlStatementValidationResult) GoString() string {
	return s.String()
}

func (s *SqlStatementValidationResult) GetErrorDetails() *ErrorDetails {
	return s.ErrorDetails
}

func (s *SqlStatementValidationResult) GetMessage() *string {
	return s.Message
}

func (s *SqlStatementValidationResult) GetSuccess() *bool {
	return s.Success
}

func (s *SqlStatementValidationResult) GetValidationResult() *string {
	return s.ValidationResult
}

func (s *SqlStatementValidationResult) SetErrorDetails(v *ErrorDetails) *SqlStatementValidationResult {
	s.ErrorDetails = v
	return s
}

func (s *SqlStatementValidationResult) SetMessage(v string) *SqlStatementValidationResult {
	s.Message = &v
	return s
}

func (s *SqlStatementValidationResult) SetSuccess(v bool) *SqlStatementValidationResult {
	s.Success = &v
	return s
}

func (s *SqlStatementValidationResult) SetValidationResult(v string) *SqlStatementValidationResult {
	s.ValidationResult = &v
	return s
}

func (s *SqlStatementValidationResult) Validate() error {
	if s.ErrorDetails != nil {
		if err := s.ErrorDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}
