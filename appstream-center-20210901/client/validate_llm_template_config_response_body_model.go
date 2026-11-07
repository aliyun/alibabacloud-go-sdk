// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateLlmTemplateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ValidateLlmTemplateConfigResponseBodyData) *ValidateLlmTemplateConfigResponseBody
	GetData() *ValidateLlmTemplateConfigResponseBodyData
	SetRequestId(v string) *ValidateLlmTemplateConfigResponseBody
	GetRequestId() *string
}

type ValidateLlmTemplateConfigResponseBody struct {
	Data *ValidateLlmTemplateConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateLlmTemplateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateLlmTemplateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateLlmTemplateConfigResponseBody) GetData() *ValidateLlmTemplateConfigResponseBodyData {
	return s.Data
}

func (s *ValidateLlmTemplateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateLlmTemplateConfigResponseBody) SetData(v *ValidateLlmTemplateConfigResponseBodyData) *ValidateLlmTemplateConfigResponseBody {
	s.Data = v
	return s
}

func (s *ValidateLlmTemplateConfigResponseBody) SetRequestId(v string) *ValidateLlmTemplateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateLlmTemplateConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateLlmTemplateConfigResponseBodyData struct {
	Valid            *bool                                                        `json:"Valid,omitempty" xml:"Valid,omitempty"`
	ValidationErrors []*ValidateLlmTemplateConfigResponseBodyDataValidationErrors `json:"ValidationErrors,omitempty" xml:"ValidationErrors,omitempty" type:"Repeated"`
}

func (s ValidateLlmTemplateConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ValidateLlmTemplateConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidateLlmTemplateConfigResponseBodyData) GetValid() *bool {
	return s.Valid
}

func (s *ValidateLlmTemplateConfigResponseBodyData) GetValidationErrors() []*ValidateLlmTemplateConfigResponseBodyDataValidationErrors {
	return s.ValidationErrors
}

func (s *ValidateLlmTemplateConfigResponseBodyData) SetValid(v bool) *ValidateLlmTemplateConfigResponseBodyData {
	s.Valid = &v
	return s
}

func (s *ValidateLlmTemplateConfigResponseBodyData) SetValidationErrors(v []*ValidateLlmTemplateConfigResponseBodyDataValidationErrors) *ValidateLlmTemplateConfigResponseBodyData {
	s.ValidationErrors = v
	return s
}

func (s *ValidateLlmTemplateConfigResponseBodyData) Validate() error {
	if s.ValidationErrors != nil {
		for _, item := range s.ValidationErrors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ValidateLlmTemplateConfigResponseBodyDataValidationErrors struct {
	// example:
	//
	// UNKNOWN_FIELD
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// unknownField
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// Unknown field "unknownField" is not recognized by Agent and will be ignored.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ValidateLlmTemplateConfigResponseBodyDataValidationErrors) String() string {
	return dara.Prettify(s)
}

func (s ValidateLlmTemplateConfigResponseBodyDataValidationErrors) GoString() string {
	return s.String()
}

func (s *ValidateLlmTemplateConfigResponseBodyDataValidationErrors) GetCode() *string {
	return s.Code
}

func (s *ValidateLlmTemplateConfigResponseBodyDataValidationErrors) GetField() *string {
	return s.Field
}

func (s *ValidateLlmTemplateConfigResponseBodyDataValidationErrors) GetMessage() *string {
	return s.Message
}

func (s *ValidateLlmTemplateConfigResponseBodyDataValidationErrors) SetCode(v string) *ValidateLlmTemplateConfigResponseBodyDataValidationErrors {
	s.Code = &v
	return s
}

func (s *ValidateLlmTemplateConfigResponseBodyDataValidationErrors) SetField(v string) *ValidateLlmTemplateConfigResponseBodyDataValidationErrors {
	s.Field = &v
	return s
}

func (s *ValidateLlmTemplateConfigResponseBodyDataValidationErrors) SetMessage(v string) *ValidateLlmTemplateConfigResponseBodyDataValidationErrors {
	s.Message = &v
	return s
}

func (s *ValidateLlmTemplateConfigResponseBodyDataValidationErrors) Validate() error {
	return dara.Validate(s)
}
