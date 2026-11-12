// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShareImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeShareImageResponseBody
	GetCode() *string
	SetData(v *DescribeShareImageResponseBodyData) *DescribeShareImageResponseBody
	GetData() *DescribeShareImageResponseBodyData
	SetMessage(v string) *DescribeShareImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeShareImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeShareImageResponseBody
	GetSuccess() *bool
}

type DescribeShareImageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Image.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned object.
	//
	// example:
	//
	// None
	Data *DescribeShareImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A0C2ED-95F2-1A37-9FC6-4A395179****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeShareImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeShareImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShareImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeShareImageResponseBody) GetData() *DescribeShareImageResponseBodyData {
	return s.Data
}

func (s *DescribeShareImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeShareImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeShareImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeShareImageResponseBody) SetCode(v string) *DescribeShareImageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeShareImageResponseBody) SetData(v *DescribeShareImageResponseBodyData) *DescribeShareImageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeShareImageResponseBody) SetMessage(v string) *DescribeShareImageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeShareImageResponseBody) SetRequestId(v string) *DescribeShareImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShareImageResponseBody) SetSuccess(v bool) *DescribeShareImageResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeShareImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeShareImageResponseBodyData struct {
	// The list of Alibaba Cloud account IDs to which the image is shared.
	AliUidList []*int64 `json:"AliUidList,omitempty" xml:"AliUidList,omitempty" type:"Repeated"`
}

func (s DescribeShareImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeShareImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeShareImageResponseBodyData) GetAliUidList() []*int64 {
	return s.AliUidList
}

func (s *DescribeShareImageResponseBodyData) SetAliUidList(v []*int64) *DescribeShareImageResponseBodyData {
	s.AliUidList = v
	return s
}

func (s *DescribeShareImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
