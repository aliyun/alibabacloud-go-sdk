// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePropertyResponseBody
	GetRequestId() *string
	SetUpdateResult(v *UpdatePropertyResponseBodyUpdateResult) *UpdatePropertyResponseBody
	GetUpdateResult() *UpdatePropertyResponseBodyUpdateResult
}

type UpdatePropertyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the modification.
	UpdateResult *UpdatePropertyResponseBodyUpdateResult `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty" type:"Struct"`
}

func (s UpdatePropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePropertyResponseBody) GetUpdateResult() *UpdatePropertyResponseBodyUpdateResult {
	return s.UpdateResult
}

func (s *UpdatePropertyResponseBody) SetRequestId(v string) *UpdatePropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePropertyResponseBody) SetUpdateResult(v *UpdatePropertyResponseBodyUpdateResult) *UpdatePropertyResponseBody {
	s.UpdateResult = v
	return s
}

func (s *UpdatePropertyResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdatePropertyResponseBodyUpdateResult struct {
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
	// testkey2
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The result of the property value modification.
	SavePropertyValueModel *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel `json:"SavePropertyValueModel,omitempty" xml:"SavePropertyValueModel,omitempty" type:"Struct"`
}

func (s UpdatePropertyResponseBodyUpdateResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResult) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResult) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *UpdatePropertyResponseBodyUpdateResult) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *UpdatePropertyResponseBodyUpdateResult) GetSavePropertyValueModel() *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel {
	return s.SavePropertyValueModel
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetPropertyId(v int64) *UpdatePropertyResponseBodyUpdateResult {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetPropertyKey(v string) *UpdatePropertyResponseBodyUpdateResult {
	s.PropertyKey = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetSavePropertyValueModel(v *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) *UpdatePropertyResponseBodyUpdateResult {
	s.SavePropertyValueModel = v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResult) Validate() error {
	return dara.Validate(s)
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel struct {
	// The property values that failed to be modified.
	FailedPropertyValues []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues `json:"FailedPropertyValues,omitempty" xml:"FailedPropertyValues,omitempty" type:"Repeated"`
	// The property values that were modified.
	SavePropertyValues []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues `json:"SavePropertyValues,omitempty" xml:"SavePropertyValues,omitempty" type:"Repeated"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) GetFailedPropertyValues() []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	return s.FailedPropertyValues
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) GetSavePropertyValues() []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues {
	return s.SavePropertyValues
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) SetFailedPropertyValues(v []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel {
	s.FailedPropertyValues = v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) SetSavePropertyValues(v []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel {
	s.SavePropertyValues = v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) Validate() error {
	return dara.Validate(s)
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues struct {
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
	// The ID of the property.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The value of the property.
	//
	// example:
	//
	// testvalue
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetErrorCode(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetErrorMessage(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetPropertyId(v int64) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetPropertyValue(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) Validate() error {
	return dara.Validate(s)
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues struct {
	// The value of the property.
	//
	// example:
	//
	// testvalue2
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) SetPropertyValue(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) SetPropertyValueId(v int64) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValueId = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) Validate() error {
	return dara.Validate(s)
}
