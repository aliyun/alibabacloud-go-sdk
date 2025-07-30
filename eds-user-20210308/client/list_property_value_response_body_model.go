// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPropertyValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyValueInfos(v []*ListPropertyValueResponseBodyPropertyValueInfos) *ListPropertyValueResponseBody
	GetPropertyValueInfos() []*ListPropertyValueResponseBodyPropertyValueInfos
	SetRequestId(v string) *ListPropertyValueResponseBody
	GetRequestId() *string
}

type ListPropertyValueResponseBody struct {
	// Details about property values.
	PropertyValueInfos []*ListPropertyValueResponseBodyPropertyValueInfos `json:"PropertyValueInfos,omitempty" xml:"PropertyValueInfos,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C52013A5-3422-5D1F-B22C-A57110972AD9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPropertyValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponseBody) GetPropertyValueInfos() []*ListPropertyValueResponseBodyPropertyValueInfos {
	return s.PropertyValueInfos
}

func (s *ListPropertyValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPropertyValueResponseBody) SetPropertyValueInfos(v []*ListPropertyValueResponseBodyPropertyValueInfos) *ListPropertyValueResponseBody {
	s.PropertyValueInfos = v
	return s
}

func (s *ListPropertyValueResponseBody) SetRequestId(v string) *ListPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPropertyValueResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPropertyValueResponseBodyPropertyValueInfos struct {
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

func (s ListPropertyValueResponseBodyPropertyValueInfos) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyValueResponseBodyPropertyValueInfos) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) SetPropertyValue(v string) *ListPropertyValueResponseBodyPropertyValueInfos {
	s.PropertyValue = &v
	return s
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) SetPropertyValueId(v int64) *ListPropertyValueResponseBodyPropertyValueInfos {
	s.PropertyValueId = &v
	return s
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) Validate() error {
	return dara.Validate(s)
}
