// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyId(v int64) *UpdatePropertyRequest
	GetPropertyId() *int64
	SetPropertyKey(v string) *UpdatePropertyRequest
	GetPropertyKey() *string
	SetPropertyValues(v []*UpdatePropertyRequestPropertyValues) *UpdatePropertyRequest
	GetPropertyValues() []*UpdatePropertyRequestPropertyValues
}

type UpdatePropertyRequest struct {
	// The ID of the property that you want to modify. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The new property name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testkey
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The values of property.
	PropertyValues []*UpdatePropertyRequestPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s UpdatePropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePropertyRequest) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *UpdatePropertyRequest) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *UpdatePropertyRequest) GetPropertyValues() []*UpdatePropertyRequestPropertyValues {
	return s.PropertyValues
}

func (s *UpdatePropertyRequest) SetPropertyId(v int64) *UpdatePropertyRequest {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyRequest) SetPropertyKey(v string) *UpdatePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *UpdatePropertyRequest) SetPropertyValues(v []*UpdatePropertyRequestPropertyValues) *UpdatePropertyRequest {
	s.PropertyValues = v
	return s
}

func (s *UpdatePropertyRequest) Validate() error {
	if s.PropertyValues != nil {
		for _, item := range s.PropertyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePropertyRequestPropertyValues struct {
	// The new property value.
	//
	// example:
	//
	// testvalue
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of property value that you want to modify. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property value ID.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s UpdatePropertyRequestPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyRequestPropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyRequestPropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *UpdatePropertyRequestPropertyValues) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *UpdatePropertyRequestPropertyValues) SetPropertyValue(v string) *UpdatePropertyRequestPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *UpdatePropertyRequestPropertyValues) SetPropertyValueId(v int64) *UpdatePropertyRequestPropertyValues {
	s.PropertyValueId = &v
	return s
}

func (s *UpdatePropertyRequestPropertyValues) Validate() error {
	return dara.Validate(s)
}
