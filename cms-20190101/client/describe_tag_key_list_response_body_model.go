// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTagKeyListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeTagKeyListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTagKeyListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTagKeyListResponseBody
	GetSuccess() *bool
	SetTagKeys(v *DescribeTagKeyListResponseBodyTagKeys) *DescribeTagKeyListResponseBody
	GetTagKeys() *DescribeTagKeyListResponseBodyTagKeys
}

type DescribeTagKeyListResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Specified parameter PageSize is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B04B8CF3-4489-432D-83BA-6F128E5F2293
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TagKeys *DescribeTagKeyListResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s DescribeTagKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTagKeyListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTagKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagKeyListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTagKeyListResponseBody) GetTagKeys() *DescribeTagKeyListResponseBodyTagKeys {
	return s.TagKeys
}

func (s *DescribeTagKeyListResponseBody) SetCode(v string) *DescribeTagKeyListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetMessage(v string) *DescribeTagKeyListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetRequestId(v string) *DescribeTagKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetSuccess(v bool) *DescribeTagKeyListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetTagKeys(v *DescribeTagKeyListResponseBodyTagKeys) *DescribeTagKeyListResponseBody {
	s.TagKeys = v
	return s
}

func (s *DescribeTagKeyListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagKeyListResponseBodyTagKeys struct {
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DescribeTagKeyListResponseBodyTagKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeyListResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListResponseBodyTagKeys) GetTagKey() []*string {
	return s.TagKey
}

func (s *DescribeTagKeyListResponseBodyTagKeys) SetTagKey(v []*string) *DescribeTagKeyListResponseBodyTagKeys {
	s.TagKey = v
	return s
}

func (s *DescribeTagKeyListResponseBodyTagKeys) Validate() error {
	return dara.Validate(s)
}
