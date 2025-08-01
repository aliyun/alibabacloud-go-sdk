// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesEcsInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesEcsInfoListResponseBody
	GetCode() *string
	SetData(v []*ListInstancesEcsInfoListResponseBodyData) *ListInstancesEcsInfoListResponseBody
	GetData() []*ListInstancesEcsInfoListResponseBodyData
	SetMessage(v string) *ListInstancesEcsInfoListResponseBody
	GetMessage() *string
}

type ListInstancesEcsInfoListResponseBody struct {
	// example:
	//
	// SysomOpenAPI.InvalidParameter
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListInstancesEcsInfoListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListInstancesEcsInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesEcsInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesEcsInfoListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesEcsInfoListResponseBody) GetData() []*ListInstancesEcsInfoListResponseBodyData {
	return s.Data
}

func (s *ListInstancesEcsInfoListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesEcsInfoListResponseBody) SetCode(v string) *ListInstancesEcsInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesEcsInfoListResponseBody) SetData(v []*ListInstancesEcsInfoListResponseBodyData) *ListInstancesEcsInfoListResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesEcsInfoListResponseBody) SetMessage(v string) *ListInstancesEcsInfoListResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesEcsInfoListResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancesEcsInfoListResponseBodyData struct {
	// example:
	//
	// 11.193.52.91
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// example:
	//
	// sysom
	TagKey *string `json:"tag_key,omitempty" xml:"tag_key,omitempty"`
	// example:
	//
	// diagnosis
	TagValue *string `json:"tag_value,omitempty" xml:"tag_value,omitempty"`
	// example:
	//
	// public
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListInstancesEcsInfoListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesEcsInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesEcsInfoListResponseBodyData) GetIp() *string {
	return s.Ip
}

func (s *ListInstancesEcsInfoListResponseBodyData) GetTagKey() *string {
	return s.TagKey
}

func (s *ListInstancesEcsInfoListResponseBodyData) GetTagValue() *string {
	return s.TagValue
}

func (s *ListInstancesEcsInfoListResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListInstancesEcsInfoListResponseBodyData) SetIp(v string) *ListInstancesEcsInfoListResponseBodyData {
	s.Ip = &v
	return s
}

func (s *ListInstancesEcsInfoListResponseBodyData) SetTagKey(v string) *ListInstancesEcsInfoListResponseBodyData {
	s.TagKey = &v
	return s
}

func (s *ListInstancesEcsInfoListResponseBodyData) SetTagValue(v string) *ListInstancesEcsInfoListResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListInstancesEcsInfoListResponseBodyData) SetType(v string) *ListInstancesEcsInfoListResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListInstancesEcsInfoListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
