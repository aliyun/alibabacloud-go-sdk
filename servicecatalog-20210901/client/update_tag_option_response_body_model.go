// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTagOptionResponseBody
	GetRequestId() *string
	SetTagOptionDetail(v *UpdateTagOptionResponseBodyTagOptionDetail) *UpdateTagOptionResponseBody
	GetTagOptionDetail() *UpdateTagOptionResponseBodyTagOptionDetail
}

type UpdateTagOptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the tag option.
	TagOptionDetail *UpdateTagOptionResponseBodyTagOptionDetail `json:"TagOptionDetail,omitempty" xml:"TagOptionDetail,omitempty" type:"Struct"`
}

func (s UpdateTagOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTagOptionResponseBody) GetTagOptionDetail() *UpdateTagOptionResponseBodyTagOptionDetail {
	return s.TagOptionDetail
}

func (s *UpdateTagOptionResponseBody) SetRequestId(v string) *UpdateTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTagOptionResponseBody) SetTagOptionDetail(v *UpdateTagOptionResponseBodyTagOptionDetail) *UpdateTagOptionResponseBody {
	s.TagOptionDetail = v
	return s
}

func (s *UpdateTagOptionResponseBody) Validate() error {
	if s.TagOptionDetail != nil {
		if err := s.TagOptionDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTagOptionResponseBodyTagOptionDetail struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	//
	// The key must be 1 to 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the Alibaba Cloud account to which the tag option belongs.
	//
	// example:
	//
	// 133413081827****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	//
	// example:
	//
	// tag-bp1u6mdf3d****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// The value must be 1 to 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTagOptionResponseBodyTagOptionDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagOptionResponseBodyTagOptionDetail) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) GetActive() *bool {
	return s.Active
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) GetKey() *string {
	return s.Key
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) GetOwner() *string {
	return s.Owner
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) GetValue() *string {
	return s.Value
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetActive(v bool) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Active = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetKey(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Key = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetOwner(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Owner = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetTagOptionId(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.TagOptionId = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetValue(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Value = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) Validate() error {
	return dara.Validate(s)
}
