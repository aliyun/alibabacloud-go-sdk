// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcHoneyPotListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeVpcHoneyPotListResponseBodyPageInfo) *DescribeVpcHoneyPotListResponseBody
	GetPageInfo() *DescribeVpcHoneyPotListResponseBodyPageInfo
	SetRequestId(v string) *DescribeVpcHoneyPotListResponseBody
	GetRequestId() *string
	SetVpcHoneyPotDTOList(v []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) *DescribeVpcHoneyPotListResponseBody
	GetVpcHoneyPotDTOList() []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList
}

type DescribeVpcHoneyPotListResponseBody struct {
	// The pagination information.
	PageInfo *DescribeVpcHoneyPotListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4FEC7F58-FCDA-415F-AE25-CD8BC0931DF2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the honeypots.
	VpcHoneyPotDTOList []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList `json:"VpcHoneyPotDTOList,omitempty" xml:"VpcHoneyPotDTOList,omitempty" type:"Repeated"`
}

func (s DescribeVpcHoneyPotListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponseBody) GetPageInfo() *DescribeVpcHoneyPotListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeVpcHoneyPotListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcHoneyPotListResponseBody) GetVpcHoneyPotDTOList() []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	return s.VpcHoneyPotDTOList
}

func (s *DescribeVpcHoneyPotListResponseBody) SetPageInfo(v *DescribeVpcHoneyPotListResponseBodyPageInfo) *DescribeVpcHoneyPotListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBody) SetRequestId(v string) *DescribeVpcHoneyPotListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBody) SetVpcHoneyPotDTOList(v []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) *DescribeVpcHoneyPotListResponseBody {
	s.VpcHoneyPotDTOList = v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VpcHoneyPotDTOList != nil {
		for _, item := range s.VpcHoneyPotDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcHoneyPotListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcHoneyPotListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetCount(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetPageSize(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList struct {
	// The CIDR block of the VPC.
	//
	// example:
	//
	// 192.168.XX.XX/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time at which the VPC was created. Unit: milliseconds.
	//
	// example:
	//
	// 1607365213000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The status of the server on which the honeypot is deployed. Valid values:
	//
	// 	- **Pending**: The server is being created.
	//
	// 	- **Running**: The server is running.
	//
	// 	- **Starting**: The server is being started.
	//
	// 	- **Stopping**: The server is being stopped.
	//
	// 	- **Stopped**: The server is stopped.
	//
	// example:
	//
	// Running
	HoneyPotEcsInstanceStatus *string `json:"HoneyPotEcsInstanceStatus,omitempty" xml:"HoneyPotEcsInstanceStatus,omitempty"`
	// The ID of the elastic network interface (ENI) used by the honeypot in the VPC.
	//
	// example:
	//
	// eni-p0whwgg7bing8b80****
	HoneyPotEniInstanceId *string `json:"HoneyPotEniInstanceId,omitempty" xml:"HoneyPotEniInstanceId,omitempty"`
	// Indicates whether the cloud honeypot feature is enabled for the VPC. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	HoneyPotExistence *bool `json:"HoneyPotExistence,omitempty" xml:"HoneyPotExistence,omitempty"`
	// The status of the honeypot. Valid values:
	//
	// 	- **pending**: The honeypot is being created.
	//
	// 	- **deleting**: The honeypot is being deleted.
	//
	// 	- **off**: The honeypot is disabled.
	//
	// 	- **suspending**: The honeypot is suspended.
	//
	// 	- **on**: The honeypot is enabled.
	//
	// example:
	//
	// on
	HoneyPotInstanceStatus *string `json:"HoneyPotInstanceStatus,omitempty" xml:"HoneyPotInstanceStatus,omitempty"`
	// The ID of the vSwitch to which the ENI used by the honeypot is connected.
	//
	// example:
	//
	// vsw-p0w7gdcfvn20tvdul****
	HoneyPotVpcSwitchId *string `json:"HoneyPotVpcSwitchId,omitempty" xml:"HoneyPotVpcSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-p0w223apdl49sr5zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// abc-vpcname
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The region ID of the VPC.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// ap-southeast-2
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// The status of the VPC. Valid values:
	//
	// 	- **Available**: The VPC is normal and available.
	//
	// 	- **Pending**: The VPC is being configured.
	//
	// example:
	//
	// Available
	VpcStatus *string `json:"VpcStatus,omitempty" xml:"VpcStatus,omitempty"`
	// An array that consists of the vSwitches in the VPC.
	VpcSwitchIdList []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList `json:"VpcSwitchIdList,omitempty" xml:"VpcSwitchIdList,omitempty" type:"Repeated"`
}

func (s DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetHoneyPotEcsInstanceStatus() *string {
	return s.HoneyPotEcsInstanceStatus
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetHoneyPotEniInstanceId() *string {
	return s.HoneyPotEniInstanceId
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetHoneyPotExistence() *bool {
	return s.HoneyPotExistence
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetHoneyPotInstanceStatus() *string {
	return s.HoneyPotInstanceStatus
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetHoneyPotVpcSwitchId() *string {
	return s.HoneyPotVpcSwitchId
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetVpcStatus() *string {
	return s.VpcStatus
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GetVpcSwitchIdList() []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList {
	return s.VpcSwitchIdList
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetCidrBlock(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetCreateTime(v int64) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotEcsInstanceStatus(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotEcsInstanceStatus = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotEniInstanceId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotEniInstanceId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotExistence(v bool) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotExistence = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotInstanceStatus(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotInstanceStatus = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotVpcSwitchId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotVpcSwitchId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcName(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcRegionId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcStatus(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcStatus = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcSwitchIdList(v []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcSwitchIdList = v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) Validate() error {
	if s.VpcSwitchIdList != nil {
		for _, item := range s.VpcSwitchIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList struct {
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-p0wdnyv4wzp6jkuu4****
	VpcSwitchId *string `json:"VpcSwitchId,omitempty" xml:"VpcSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// abc01
	VpcSwitchName *string `json:"VpcSwitchName,omitempty" xml:"VpcSwitchName,omitempty"`
	// The zone ID of the vSwitch.
	//
	// example:
	//
	// ap-southeast-2b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) GetVpcSwitchId() *string {
	return s.VpcSwitchId
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) GetVpcSwitchName() *string {
	return s.VpcSwitchName
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) SetVpcSwitchId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList {
	s.VpcSwitchId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) SetVpcSwitchName(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList {
	s.VpcSwitchName = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) SetZoneId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOListVpcSwitchIdList) Validate() error {
	return dara.Validate(s)
}
