// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsedPropertyValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyId(v int64) *CheckUsedPropertyValueRequest
	GetPropertyId() *int64
	SetPropertyValueId(v int64) *CheckUsedPropertyValueRequest
	GetPropertyValueId() *int64
}

type CheckUsedPropertyValueRequest struct {
	// The property ID. You can call the [ListProperty](~~ListProperty~~) operation to query property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 380
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The ID of the property value. You can call the [ListProperty](~~ListProperty~~) operation to query the ID of the property value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s CheckUsedPropertyValueRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUsedPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueRequest) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *CheckUsedPropertyValueRequest) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *CheckUsedPropertyValueRequest) SetPropertyId(v int64) *CheckUsedPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *CheckUsedPropertyValueRequest) SetPropertyValueId(v int64) *CheckUsedPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

func (s *CheckUsedPropertyValueRequest) Validate() error {
	return dara.Validate(s)
}
