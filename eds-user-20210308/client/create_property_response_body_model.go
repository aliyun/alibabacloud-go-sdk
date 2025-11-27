// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateResult(v *CreatePropertyResponseBodyCreateResult) *CreatePropertyResponseBody
	GetCreateResult() *CreatePropertyResponseBodyCreateResult
	SetRequestId(v string) *CreatePropertyResponseBody
	GetRequestId() *string
}

type CreatePropertyResponseBody struct {
	// The result of creating the property.
	CreateResult *CreatePropertyResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBody) GetCreateResult() *CreatePropertyResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreatePropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePropertyResponseBody) SetCreateResult(v *CreatePropertyResponseBodyCreateResult) *CreatePropertyResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreatePropertyResponseBody) SetRequestId(v string) *CreatePropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePropertyResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePropertyResponseBodyCreateResult struct {
	// The ID of the property.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The name of the property.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The result of creating the property value.
	SavePropertyValueModel *CreatePropertyResponseBodyCreateResultSavePropertyValueModel `json:"SavePropertyValueModel,omitempty" xml:"SavePropertyValueModel,omitempty" type:"Struct"`
}

func (s CreatePropertyResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResult) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *CreatePropertyResponseBodyCreateResult) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *CreatePropertyResponseBodyCreateResult) GetSavePropertyValueModel() *CreatePropertyResponseBodyCreateResultSavePropertyValueModel {
	return s.SavePropertyValueModel
}

func (s *CreatePropertyResponseBodyCreateResult) SetPropertyId(v int64) *CreatePropertyResponseBodyCreateResult {
	s.PropertyId = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResult) SetPropertyKey(v string) *CreatePropertyResponseBodyCreateResult {
	s.PropertyKey = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResult) SetSavePropertyValueModel(v *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) *CreatePropertyResponseBodyCreateResult {
	s.SavePropertyValueModel = v
	return s
}

func (s *CreatePropertyResponseBodyCreateResult) Validate() error {
	if s.SavePropertyValueModel != nil {
		if err := s.SavePropertyValueModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModel struct {
	// The property values that failed to be created.
	FailedPropertyValues []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues `json:"FailedPropertyValues,omitempty" xml:"FailedPropertyValues,omitempty" type:"Repeated"`
	// Details of the property values that were created.
	SavePropertyValues []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues `json:"SavePropertyValues,omitempty" xml:"SavePropertyValues,omitempty" type:"Repeated"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModel) String() string {
	return dara.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModel) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) GetFailedPropertyValues() []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	return s.FailedPropertyValues
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) GetSavePropertyValues() []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues {
	return s.SavePropertyValues
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) SetFailedPropertyValues(v []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) *CreatePropertyResponseBodyCreateResultSavePropertyValueModel {
	s.FailedPropertyValues = v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) SetSavePropertyValues(v []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) *CreatePropertyResponseBodyCreateResultSavePropertyValueModel {
	s.SavePropertyValues = v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) Validate() error {
	if s.FailedPropertyValues != nil {
		for _, item := range s.FailedPropertyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SavePropertyValues != nil {
		for _, item := range s.SavePropertyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues struct {
	// The error code.
	//
	// example:
	//
	// ExistedPropertyValue
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The property value is used by another property.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The value of the property.
	//
	// example:
	//
	// HR
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetErrorCode(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorCode = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetErrorMessage(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetPropertyId(v int64) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyId = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetPropertyValue(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) Validate() error {
	return dara.Validate(s)
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues struct {
	// The value of the property.
	//
	// example:
	//
	// HR
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) String() string {
	return dara.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) SetPropertyValue(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) SetPropertyValueId(v int64) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValueId = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) Validate() error {
	return dara.Validate(s)
}
