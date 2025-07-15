// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclResourceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAclResourceNameResponseBody
	GetCode() *int32
	SetData(v *DescribeAclResourceNameResponseBodyData) *DescribeAclResourceNameResponseBody
	GetData() *DescribeAclResourceNameResponseBodyData
	SetMessage(v string) *DescribeAclResourceNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAclResourceNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAclResourceNameResponseBody
	GetSuccess() *bool
}

type DescribeAclResourceNameResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeAclResourceNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAclResourceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclResourceNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAclResourceNameResponseBody) GetData() *DescribeAclResourceNameResponseBodyData {
	return s.Data
}

func (s *DescribeAclResourceNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAclResourceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclResourceNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAclResourceNameResponseBody) SetCode(v int32) *DescribeAclResourceNameResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetData(v *DescribeAclResourceNameResponseBodyData) *DescribeAclResourceNameResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetMessage(v string) *DescribeAclResourceNameResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetRequestId(v string) *DescribeAclResourceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetSuccess(v bool) *DescribeAclResourceNameResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAclResourceNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAclResourceNameResponseBodyData struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeAclResourceNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclResourceNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameResponseBodyData) GetData() []*string {
	return s.Data
}

func (s *DescribeAclResourceNameResponseBodyData) SetData(v []*string) *DescribeAclResourceNameResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeAclResourceNameResponseBodyData) Validate() error {
	return dara.Validate(s)
}
