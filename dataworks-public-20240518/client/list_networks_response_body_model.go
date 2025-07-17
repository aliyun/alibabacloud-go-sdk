// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListNetworksResponseBodyPagingInfo) *ListNetworksResponseBody
	GetPagingInfo() *ListNetworksResponseBodyPagingInfo
	SetRequestId(v string) *ListNetworksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNetworksResponseBody
	GetSuccess() *bool
}

type ListNetworksResponseBody struct {
	// The pagination information.
	PagingInfo *ListNetworksResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The ID of the request. It is used to locate logs and troubleshoot problems.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNetworksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworksResponseBody) GetPagingInfo() *ListNetworksResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListNetworksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNetworksResponseBody) SetPagingInfo(v *ListNetworksResponseBodyPagingInfo) *ListNetworksResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNetworksResponseBody) SetRequestId(v string) *ListNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworksResponseBody) SetSuccess(v bool) *ListNetworksResponseBody {
	s.Success = &v
	return s
}

func (s *ListNetworksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNetworksResponseBodyPagingInfo struct {
	// The network resources of the serverless resource group.
	NetworkList []*ListNetworksResponseBodyPagingInfoNetworkList `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
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
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworksResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListNetworksResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNetworksResponseBodyPagingInfo) GetNetworkList() []*ListNetworksResponseBodyPagingInfoNetworkList {
	return s.NetworkList
}

func (s *ListNetworksResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNetworksResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetworksResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNetworksResponseBodyPagingInfo) SetNetworkList(v []*ListNetworksResponseBodyPagingInfoNetworkList) *ListNetworksResponseBodyPagingInfo {
	s.NetworkList = v
	return s
}

func (s *ListNetworksResponseBodyPagingInfo) SetPageNumber(v int32) *ListNetworksResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfo) SetPageSize(v int32) *ListNetworksResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfo) SetTotalCount(v int32) *ListNetworksResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListNetworksResponseBodyPagingInfoNetworkList struct {
	// The time when the network resource was created. The value is a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who creates the network resource.
	//
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The network ID.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the serverless resource group.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze13vamugr7jenXXXXX
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the network resource. Valid values: Pending, Creating, Running, Deleting, and Deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VSwitch ID.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListNetworksResponseBodyPagingInfoNetworkList) String() string {
	return dara.Prettify(s)
}

func (s ListNetworksResponseBodyPagingInfoNetworkList) GoString() string {
	return s.String()
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetId() *int64 {
	return s.Id
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetStatus() *string {
	return s.Status
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetCreateTime(v int64) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.CreateTime = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetCreateUser(v string) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.CreateUser = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetId(v int64) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.Id = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetResourceGroupId(v string) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetSecurityGroupId(v string) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.SecurityGroupId = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetStatus(v string) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.Status = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetVpcId(v string) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.VpcId = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) SetVswitchId(v string) *ListNetworksResponseBodyPagingInfoNetworkList {
	s.VswitchId = &v
	return s
}

func (s *ListNetworksResponseBodyPagingInfoNetworkList) Validate() error {
	return dara.Validate(s)
}
