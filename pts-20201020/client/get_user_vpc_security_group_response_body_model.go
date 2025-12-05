// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserVpcSecurityGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUserVpcSecurityGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserVpcSecurityGroupResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetUserVpcSecurityGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetUserVpcSecurityGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetUserVpcSecurityGroupResponseBody
	GetRequestId() *string
	SetSecurityGroupCount(v int32) *GetUserVpcSecurityGroupResponseBody
	GetSecurityGroupCount() *int32
	SetSecurityGroupList(v []*GetUserVpcSecurityGroupResponseBodySecurityGroupList) *GetUserVpcSecurityGroupResponseBody
	GetSecurityGroupList() []*GetUserVpcSecurityGroupResponseBodySecurityGroupList
	SetSuccess(v bool) *GetUserVpcSecurityGroupResponseBody
	GetSuccess() *bool
}

type GetUserVpcSecurityGroupResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 61B15017-1A68-5C47-834F-87E2BBC44F2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of security groups.
	//
	// example:
	//
	// 9
	SecurityGroupCount *int32 `json:"SecurityGroupCount,omitempty" xml:"SecurityGroupCount,omitempty"`
	// The security groups.
	SecurityGroupList []*GetUserVpcSecurityGroupResponseBodySecurityGroupList `json:"SecurityGroupList,omitempty" xml:"SecurityGroupList,omitempty" type:"Repeated"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserVpcSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserVpcSecurityGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserVpcSecurityGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserVpcSecurityGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserVpcSecurityGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetUserVpcSecurityGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUserVpcSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserVpcSecurityGroupResponseBody) GetSecurityGroupCount() *int32 {
	return s.SecurityGroupCount
}

func (s *GetUserVpcSecurityGroupResponseBody) GetSecurityGroupList() []*GetUserVpcSecurityGroupResponseBodySecurityGroupList {
	return s.SecurityGroupList
}

func (s *GetUserVpcSecurityGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserVpcSecurityGroupResponseBody) SetCode(v string) *GetUserVpcSecurityGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetHttpStatusCode(v int32) *GetUserVpcSecurityGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetMessage(v string) *GetUserVpcSecurityGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetPageNumber(v int32) *GetUserVpcSecurityGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetPageSize(v int32) *GetUserVpcSecurityGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetRequestId(v string) *GetUserVpcSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetSecurityGroupCount(v int32) *GetUserVpcSecurityGroupResponseBody {
	s.SecurityGroupCount = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetSecurityGroupList(v []*GetUserVpcSecurityGroupResponseBodySecurityGroupList) *GetUserVpcSecurityGroupResponseBody {
	s.SecurityGroupList = v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) SetSuccess(v bool) *GetUserVpcSecurityGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBody) Validate() error {
	if s.SecurityGroupList != nil {
		for _, item := range s.SecurityGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserVpcSecurityGroupResponseBodySecurityGroupList struct {
	// The description of the security group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp16bt3zuugxpfjkasdfvthxth8
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the security group.
	//
	// example:
	//
	// my-security-group
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf6tar2ohlasdhsatjln37h30bv
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetUserVpcSecurityGroupResponseBodySecurityGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcSecurityGroupResponseBodySecurityGroupList) GoString() string {
	return s.String()
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) GetDescription() *string {
	return s.Description
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) GetVpcId() *string {
	return s.VpcId
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) SetDescription(v string) *GetUserVpcSecurityGroupResponseBodySecurityGroupList {
	s.Description = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) SetSecurityGroupId(v string) *GetUserVpcSecurityGroupResponseBodySecurityGroupList {
	s.SecurityGroupId = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) SetSecurityGroupName(v string) *GetUserVpcSecurityGroupResponseBodySecurityGroupList {
	s.SecurityGroupName = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) SetVpcId(v string) *GetUserVpcSecurityGroupResponseBodySecurityGroupList {
	s.VpcId = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponseBodySecurityGroupList) Validate() error {
	return dara.Validate(s)
}
