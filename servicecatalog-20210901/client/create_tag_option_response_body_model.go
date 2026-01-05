// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTagOptionResponseBody
	GetRequestId() *string
	SetTagOptionDetail(v *CreateTagOptionResponseBodyTagOptionDetail) *CreateTagOptionResponseBody
	GetTagOptionDetail() *CreateTagOptionResponseBodyTagOptionDetail
}

type CreateTagOptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A48A5F12-6098-54A1-A389-6834AF27****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the tag option.
	TagOptionDetail *CreateTagOptionResponseBodyTagOptionDetail `json:"TagOptionDetail,omitempty" xml:"TagOptionDetail,omitempty" type:"Struct"`
}

func (s CreateTagOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagOptionResponseBody) GetTagOptionDetail() *CreateTagOptionResponseBodyTagOptionDetail {
	return s.TagOptionDetail
}

func (s *CreateTagOptionResponseBody) SetRequestId(v string) *CreateTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagOptionResponseBody) SetTagOptionDetail(v *CreateTagOptionResponseBodyTagOptionDetail) *CreateTagOptionResponseBody {
	s.TagOptionDetail = v
	return s
}

func (s *CreateTagOptionResponseBody) Validate() error {
	if s.TagOptionDetail != nil {
		if err := s.TagOptionDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTagOptionResponseBodyTagOptionDetail struct {
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
	// The key must be 1 to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
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
	// tag-bp1r3mxq3t****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// The value must be 1 to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagOptionResponseBodyTagOptionDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateTagOptionResponseBodyTagOptionDetail) GoString() string {
	return s.String()
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) GetActive() *bool {
	return s.Active
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) GetKey() *string {
	return s.Key
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) GetOwner() *string {
	return s.Owner
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) GetValue() *string {
	return s.Value
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetActive(v bool) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Active = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetKey(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Key = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetOwner(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Owner = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetTagOptionId(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.TagOptionId = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetValue(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Value = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) Validate() error {
	return dara.Validate(s)
}
