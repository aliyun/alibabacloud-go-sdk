// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListPropertyResponseBody
	GetNextToken() *string
	SetProperties(v []*ListPropertyResponseBodyProperties) *ListPropertyResponseBody
	GetProperties() []*ListPropertyResponseBodyProperties
	SetRequestId(v string) *ListPropertyResponseBody
	GetRequestId() *string
}

type ListPropertyResponseBody struct {
	// The token used to start the next query. An empty NextToken indicates that no more results are available.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The property information.
	Properties []*ListPropertyResponseBodyProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPropertyResponseBody) GetProperties() []*ListPropertyResponseBodyProperties {
	return s.Properties
}

func (s *ListPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPropertyResponseBody) SetNextToken(v string) *ListPropertyResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPropertyResponseBody) SetProperties(v []*ListPropertyResponseBodyProperties) *ListPropertyResponseBody {
	s.Properties = v
	return s
}

func (s *ListPropertyResponseBody) SetRequestId(v string) *ListPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPropertyResponseBody) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPropertyResponseBodyProperties struct {
	// The property ID.
	//
	// example:
	//
	// 30
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property name.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property value information.
	PropertyValues []*ListPropertyResponseBodyPropertiesPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s ListPropertyResponseBodyProperties) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyResponseBodyProperties) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBodyProperties) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *ListPropertyResponseBodyProperties) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *ListPropertyResponseBodyProperties) GetPropertyValues() []*ListPropertyResponseBodyPropertiesPropertyValues {
	return s.PropertyValues
}

func (s *ListPropertyResponseBodyProperties) SetPropertyId(v int64) *ListPropertyResponseBodyProperties {
	s.PropertyId = &v
	return s
}

func (s *ListPropertyResponseBodyProperties) SetPropertyKey(v string) *ListPropertyResponseBodyProperties {
	s.PropertyKey = &v
	return s
}

func (s *ListPropertyResponseBodyProperties) SetPropertyValues(v []*ListPropertyResponseBodyPropertiesPropertyValues) *ListPropertyResponseBodyProperties {
	s.PropertyValues = v
	return s
}

func (s *ListPropertyResponseBodyProperties) Validate() error {
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

type ListPropertyResponseBodyPropertiesPropertyValues struct {
	// The property value.
	//
	// example:
	//
	// A
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The property value ID.
	//
	// example:
	//
	// 42
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s ListPropertyResponseBodyPropertiesPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyResponseBodyPropertiesPropertyValues) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) SetPropertyValue(v string) *ListPropertyResponseBodyPropertiesPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) SetPropertyValueId(v int64) *ListPropertyResponseBodyPropertiesPropertyValues {
	s.PropertyValueId = &v
	return s
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) Validate() error {
	return dara.Validate(s)
}
