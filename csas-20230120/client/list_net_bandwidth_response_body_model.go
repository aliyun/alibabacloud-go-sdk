// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListNetBandwidthResponseBody
	GetCurrentPage() *int32
	SetNetBandwidthList(v []*ListNetBandwidthResponseBodyNetBandwidthList) *ListNetBandwidthResponseBody
	GetNetBandwidthList() []*ListNetBandwidthResponseBodyNetBandwidthList
	SetPageSize(v int32) *ListNetBandwidthResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNetBandwidthResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListNetBandwidthResponseBody
	GetTotalNum() *int32
}

type ListNetBandwidthResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of bandwidth configurations.
	NetBandwidthList []*ListNetBandwidthResponseBodyNetBandwidthList `json:"NetBandwidthList,omitempty" xml:"NetBandwidthList,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListNetBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetBandwidthResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListNetBandwidthResponseBody) GetNetBandwidthList() []*ListNetBandwidthResponseBodyNetBandwidthList {
	return s.NetBandwidthList
}

func (s *ListNetBandwidthResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetBandwidthResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListNetBandwidthResponseBody) SetCurrentPage(v int32) *ListNetBandwidthResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListNetBandwidthResponseBody) SetNetBandwidthList(v []*ListNetBandwidthResponseBodyNetBandwidthList) *ListNetBandwidthResponseBody {
	s.NetBandwidthList = v
	return s
}

func (s *ListNetBandwidthResponseBody) SetPageSize(v int32) *ListNetBandwidthResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNetBandwidthResponseBody) SetRequestId(v string) *ListNetBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetBandwidthResponseBody) SetTotalNum(v int32) *ListNetBandwidthResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListNetBandwidthResponseBody) Validate() error {
	if s.NetBandwidthList != nil {
		for _, item := range s.NetBandwidthList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetBandwidthResponseBodyNetBandwidthList struct {
	// The bandwidth value, in Mbps.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2026-08-01 10:20:30
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 2026-08-02 15:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// vpc-bp1234567890
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// production-vpc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type.
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListNetBandwidthResponseBodyNetBandwidthList) String() string {
	return dara.Prettify(s)
}

func (s ListNetBandwidthResponseBodyNetBandwidthList) GoString() string {
	return s.String()
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) GetNetType() *string {
	return s.NetType
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) GetRegion() *string {
	return s.Region
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) SetBandwidth(v int32) *ListNetBandwidthResponseBodyNetBandwidthList {
	s.Bandwidth = &v
	return s
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) SetGmtCreate(v string) *ListNetBandwidthResponseBodyNetBandwidthList {
	s.GmtCreate = &v
	return s
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) SetGmtModified(v string) *ListNetBandwidthResponseBodyNetBandwidthList {
	s.GmtModified = &v
	return s
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) SetInstanceId(v string) *ListNetBandwidthResponseBodyNetBandwidthList {
	s.InstanceId = &v
	return s
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) SetInstanceName(v string) *ListNetBandwidthResponseBodyNetBandwidthList {
	s.InstanceName = &v
	return s
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) SetNetType(v string) *ListNetBandwidthResponseBodyNetBandwidthList {
	s.NetType = &v
	return s
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) SetRegion(v string) *ListNetBandwidthResponseBodyNetBandwidthList {
	s.Region = &v
	return s
}

func (s *ListNetBandwidthResponseBodyNetBandwidthList) Validate() error {
	return dara.Validate(s)
}
