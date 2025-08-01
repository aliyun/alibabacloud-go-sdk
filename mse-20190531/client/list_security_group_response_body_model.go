// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSecurityGroupResponseBody
	GetCode() *int32
	SetData(v []*ListSecurityGroupResponseBodyData) *ListSecurityGroupResponseBody
	GetData() []*ListSecurityGroupResponseBodyData
	SetHttpStatusCode(v int32) *ListSecurityGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSecurityGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSecurityGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSecurityGroupResponseBody
	GetSuccess() *bool
}

type ListSecurityGroupResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	Data []*ListSecurityGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSecurityGroupResponseBody) GetData() []*ListSecurityGroupResponseBodyData {
	return s.Data
}

func (s *ListSecurityGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSecurityGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSecurityGroupResponseBody) SetCode(v int32) *ListSecurityGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecurityGroupResponseBody) SetData(v []*ListSecurityGroupResponseBodyData) *ListSecurityGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListSecurityGroupResponseBody) SetHttpStatusCode(v int32) *ListSecurityGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSecurityGroupResponseBody) SetMessage(v string) *ListSecurityGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecurityGroupResponseBody) SetRequestId(v string) *ListSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityGroupResponseBody) SetSuccess(v bool) *ListSecurityGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSecurityGroupResponseBodyData struct {
	// The ID of the security group.
	//
	// example:
	//
	// sg-8vb8gsmrqyc35k645rk6
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the security group.
	//
	// example:
	//
	// sg2
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// The type of the security group. Valid values:
	//
	// 	- normal: basic security group
	//
	// 	- enterprise: advanced security group For more information, see [Advanced security groups](https://help.aliyun.com/document_detail/120621.html).
	//
	// example:
	//
	// enterprise
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1b
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListSecurityGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListSecurityGroupResponseBodyData) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *ListSecurityGroupResponseBodyData) GetSecurityGroupType() *string {
	return s.SecurityGroupType
}

func (s *ListSecurityGroupResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *ListSecurityGroupResponseBodyData) SetSecurityGroupId(v string) *ListSecurityGroupResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *ListSecurityGroupResponseBodyData) SetSecurityGroupName(v string) *ListSecurityGroupResponseBodyData {
	s.SecurityGroupName = &v
	return s
}

func (s *ListSecurityGroupResponseBodyData) SetSecurityGroupType(v string) *ListSecurityGroupResponseBodyData {
	s.SecurityGroupType = &v
	return s
}

func (s *ListSecurityGroupResponseBodyData) SetVpcId(v string) *ListSecurityGroupResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *ListSecurityGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
