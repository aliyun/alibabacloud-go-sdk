// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserVpcsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUserVpcsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserVpcsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetUserVpcsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetUserVpcsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetUserVpcsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserVpcsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetUserVpcsResponseBody
	GetTotalCount() *int64
	SetVpcs(v []*GetUserVpcsResponseBodyVpcs) *GetUserVpcsResponseBody
	GetVpcs() []*GetUserVpcsResponseBodyVpcs
}

type GetUserVpcsResponseBody struct {
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
	// CEE46204-E1CF-5F48-B094-67362DD4B73F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The VPCs.
	Vpcs []*GetUserVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
}

func (s GetUserVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserVpcsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserVpcsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserVpcsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserVpcsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetUserVpcsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUserVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserVpcsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserVpcsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetUserVpcsResponseBody) GetVpcs() []*GetUserVpcsResponseBodyVpcs {
	return s.Vpcs
}

func (s *GetUserVpcsResponseBody) SetCode(v string) *GetUserVpcsResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetHttpStatusCode(v int32) *GetUserVpcsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetMessage(v string) *GetUserVpcsResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetPageNumber(v int32) *GetUserVpcsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetPageSize(v int32) *GetUserVpcsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetRequestId(v string) *GetUserVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetSuccess(v bool) *GetUserVpcsResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetTotalCount(v int64) *GetUserVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetUserVpcsResponseBody) SetVpcs(v []*GetUserVpcsResponseBodyVpcs) *GetUserVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *GetUserVpcsResponseBody) Validate() error {
	if s.Vpcs != nil {
		for _, item := range s.Vpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserVpcsResponseBodyVpcs struct {
	// The IPv4 CIDR block of the VPC.
	//
	// example:
	//
	// 172.16.80.0/20
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the VPC.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the VPC belongs.
	//
	// example:
	//
	// rg-acfm3fzmgkehpewjertna
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the route tables.
	RouterTableIds []*string `json:"RouterTableIds,omitempty" xml:"RouterTableIds,omitempty" type:"Repeated"`
	// The vSwitches.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf6gc56wdjpafoiwej6adqb4qn72xtw
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// shanghai-vpc
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s GetUserVpcsResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *GetUserVpcsResponseBodyVpcs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *GetUserVpcsResponseBodyVpcs) GetDescription() *string {
	return s.Description
}

func (s *GetUserVpcsResponseBodyVpcs) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserVpcsResponseBodyVpcs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetUserVpcsResponseBodyVpcs) GetRouterTableIds() []*string {
	return s.RouterTableIds
}

func (s *GetUserVpcsResponseBodyVpcs) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetUserVpcsResponseBodyVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *GetUserVpcsResponseBodyVpcs) GetVpcName() *string {
	return s.VpcName
}

func (s *GetUserVpcsResponseBodyVpcs) SetCidrBlock(v string) *GetUserVpcsResponseBodyVpcs {
	s.CidrBlock = &v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) SetDescription(v string) *GetUserVpcsResponseBodyVpcs {
	s.Description = &v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) SetRegionId(v string) *GetUserVpcsResponseBodyVpcs {
	s.RegionId = &v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) SetResourceGroupId(v string) *GetUserVpcsResponseBodyVpcs {
	s.ResourceGroupId = &v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) SetRouterTableIds(v []*string) *GetUserVpcsResponseBodyVpcs {
	s.RouterTableIds = v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) SetVSwitchIds(v []*string) *GetUserVpcsResponseBodyVpcs {
	s.VSwitchIds = v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) SetVpcId(v string) *GetUserVpcsResponseBodyVpcs {
	s.VpcId = &v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) SetVpcName(v string) *GetUserVpcsResponseBodyVpcs {
	s.VpcName = &v
	return s
}

func (s *GetUserVpcsResponseBodyVpcs) Validate() error {
	return dara.Validate(s)
}
