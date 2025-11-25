// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeInternetSlbResponseBodyDataList) *DescribeInternetSlbResponseBody
	GetDataList() []*DescribeInternetSlbResponseBodyDataList
	SetPageInfo(v *DescribeInternetSlbResponseBodyPageInfo) *DescribeInternetSlbResponseBody
	GetPageInfo() *DescribeInternetSlbResponseBodyPageInfo
	SetRequestId(v string) *DescribeInternetSlbResponseBody
	GetRequestId() *string
}

type DescribeInternetSlbResponseBody struct {
	DataList []*DescribeInternetSlbResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	PageInfo *DescribeInternetSlbResponseBodyPageInfo   `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 135BF83A-0416-5A11-96BB-FA7604D4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetSlbResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetSlbResponseBody) GetDataList() []*DescribeInternetSlbResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetSlbResponseBody) GetPageInfo() *DescribeInternetSlbResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInternetSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetSlbResponseBody) SetDataList(v []*DescribeInternetSlbResponseBodyDataList) *DescribeInternetSlbResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetSlbResponseBody) SetPageInfo(v *DescribeInternetSlbResponseBodyPageInfo) *DescribeInternetSlbResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInternetSlbResponseBody) SetRequestId(v string) *DescribeInternetSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetSlbResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInternetSlbResponseBodyDataList struct {
	// example:
	//
	// acl-uf66n6l9mf3fgq8xs****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// example:
	//
	// on
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// example:
	//
	// normal
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// example:
	//
	// lb-2ze8v2x5kd9qyvp2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// buyerpro1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// 1883
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 39.108.57.XXX
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo *string   `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Tags     []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeInternetSlbResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetSlbResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetSlbResponseBodyDataList) GetAclId() *string {
	return s.AclId
}

func (s *DescribeInternetSlbResponseBodyDataList) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeInternetSlbResponseBodyDataList) GetAclType() *string {
	return s.AclType
}

func (s *DescribeInternetSlbResponseBodyDataList) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeInternetSlbResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInternetSlbResponseBodyDataList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInternetSlbResponseBodyDataList) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeInternetSlbResponseBodyDataList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeInternetSlbResponseBodyDataList) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeInternetSlbResponseBodyDataList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInternetSlbResponseBodyDataList) GetTags() []*string {
	return s.Tags
}

func (s *DescribeInternetSlbResponseBodyDataList) SetAclId(v string) *DescribeInternetSlbResponseBodyDataList {
	s.AclId = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetAclStatus(v string) *DescribeInternetSlbResponseBodyDataList {
	s.AclStatus = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetAclType(v string) *DescribeInternetSlbResponseBodyDataList {
	s.AclType = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetHealthStatus(v string) *DescribeInternetSlbResponseBodyDataList {
	s.HealthStatus = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetInstanceId(v string) *DescribeInternetSlbResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetInstanceName(v string) *DescribeInternetSlbResponseBodyDataList {
	s.InstanceName = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetIpProtocol(v string) *DescribeInternetSlbResponseBodyDataList {
	s.IpProtocol = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetPort(v int32) *DescribeInternetSlbResponseBodyDataList {
	s.Port = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetPublicIp(v string) *DescribeInternetSlbResponseBodyDataList {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetRegionNo(v string) *DescribeInternetSlbResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) SetTags(v []*string) *DescribeInternetSlbResponseBodyDataList {
	s.Tags = v
	return s
}

func (s *DescribeInternetSlbResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeInternetSlbResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInternetSlbResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetSlbResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInternetSlbResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInternetSlbResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInternetSlbResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInternetSlbResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInternetSlbResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyPageInfo) SetPageSize(v int32) *DescribeInternetSlbResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInternetSlbResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInternetSlbResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
