// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTagOptionResponseBody
	GetRequestId() *string
	SetTagOptionDetail(v *GetTagOptionResponseBodyTagOptionDetail) *GetTagOptionResponseBody
	GetTagOptionDetail() *GetTagOptionResponseBodyTagOptionDetail
}

type GetTagOptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C1509725-055D-5C7B-9420-8B737DBD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the tag option.
	TagOptionDetail *GetTagOptionResponseBodyTagOptionDetail `json:"TagOptionDetail,omitempty" xml:"TagOptionDetail,omitempty" type:"Struct"`
}

func (s GetTagOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTagOptionResponseBody) GetTagOptionDetail() *GetTagOptionResponseBodyTagOptionDetail {
	return s.TagOptionDetail
}

func (s *GetTagOptionResponseBody) SetRequestId(v string) *GetTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagOptionResponseBody) SetTagOptionDetail(v *GetTagOptionResponseBodyTagOptionDetail) *GetTagOptionResponseBody {
	s.TagOptionDetail = v
	return s
}

func (s *GetTagOptionResponseBody) Validate() error {
	if s.TagOptionDetail != nil {
		if err := s.TagOptionDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTagOptionResponseBodyTagOptionDetail struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
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
	// tag-bp15qmwz3r****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTagOptionResponseBodyTagOptionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetTagOptionResponseBodyTagOptionDetail) GoString() string {
	return s.String()
}

func (s *GetTagOptionResponseBodyTagOptionDetail) GetActive() *bool {
	return s.Active
}

func (s *GetTagOptionResponseBodyTagOptionDetail) GetKey() *string {
	return s.Key
}

func (s *GetTagOptionResponseBodyTagOptionDetail) GetOwner() *string {
	return s.Owner
}

func (s *GetTagOptionResponseBodyTagOptionDetail) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *GetTagOptionResponseBodyTagOptionDetail) GetValue() *string {
	return s.Value
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetActive(v bool) *GetTagOptionResponseBodyTagOptionDetail {
	s.Active = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetKey(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.Key = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetOwner(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.Owner = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetTagOptionId(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.TagOptionId = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetValue(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.Value = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) Validate() error {
	return dara.Validate(s)
}
