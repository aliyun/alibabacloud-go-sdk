// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDDoSInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceInfo(v []*ListDDoSInstancesResponseBodyInstanceInfo) *ListDDoSInstancesResponseBody
	GetInstanceInfo() []*ListDDoSInstancesResponseBodyInstanceInfo
	SetPageNumber(v int32) *ListDDoSInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDDoSInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDDoSInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDDoSInstancesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListDDoSInstancesResponseBody
	GetTotalPage() *int32
}

type ListDDoSInstancesResponseBody struct {
	// Instance details
	InstanceInfo []*ListDDoSInstancesResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
	// Page number. Default: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 9F1DC265-BF10-5C9C-B607-760265C5F365
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of entries
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListDDoSInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDDoSInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDDoSInstancesResponseBody) GetInstanceInfo() []*ListDDoSInstancesResponseBodyInstanceInfo {
	return s.InstanceInfo
}

func (s *ListDDoSInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDDoSInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDDoSInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDDoSInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDDoSInstancesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListDDoSInstancesResponseBody) SetInstanceInfo(v []*ListDDoSInstancesResponseBodyInstanceInfo) *ListDDoSInstancesResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *ListDDoSInstancesResponseBody) SetPageNumber(v int32) *ListDDoSInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDDoSInstancesResponseBody) SetPageSize(v int32) *ListDDoSInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDDoSInstancesResponseBody) SetRequestId(v string) *ListDDoSInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDDoSInstancesResponseBody) SetTotalCount(v int32) *ListDDoSInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDDoSInstancesResponseBody) SetTotalPage(v int32) *ListDDoSInstancesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListDDoSInstancesResponseBody) Validate() error {
	if s.InstanceInfo != nil {
		for _, item := range s.InstanceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDDoSInstancesResponseBodyInstanceInfo struct {
	// Time when the instance was purchased
	//
	// example:
	//
	// 2025-07-01T07:59:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// DDoS protection specification for the Chinese mainland
	//
	// example:
	//
	// cn_300
	DDoSBurstableDomesticProtection *string `json:"DDoSBurstableDomesticProtection,omitempty" xml:"DDoSBurstableDomesticProtection,omitempty"`
	// DDoS protection specification for regions outside the Chinese mainland
	//
	// example:
	//
	// overseas_300
	DDoSBurstableOverseasProtection *string `json:"DDoSBurstableOverseasProtection,omitempty" xml:"DDoSBurstableOverseasProtection,omitempty"`
	// Instance ID
	//
	// example:
	//
	// sp-ddddxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scheduled release time
	//
	// example:
	//
	// 2026-03-25T16:00:00Z
	ReserveReleaseTime *string `json:"ReserveReleaseTime,omitempty" xml:"ReserveReleaseTime,omitempty"`
	// Associated site package instance ID
	//
	// example:
	//
	// esa-site-b0s6kmx0r0n4
	SiteInstanceId *string `json:"SiteInstanceId,omitempty" xml:"SiteInstanceId,omitempty"`
	// Instance status. Valid values:
	//
	// - **online**: Normal service status.
	//
	// - **offline**: Expired but not overdue. Instance is unavailable.
	//
	// - **disable**: Released.
	//
	// - **overdue**: Service suspended due to overdue payment.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDDoSInstancesResponseBodyInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDDoSInstancesResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) GetDDoSBurstableDomesticProtection() *string {
	return s.DDoSBurstableDomesticProtection
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) GetDDoSBurstableOverseasProtection() *string {
	return s.DDoSBurstableOverseasProtection
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) GetReserveReleaseTime() *string {
	return s.ReserveReleaseTime
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) GetSiteInstanceId() *string {
	return s.SiteInstanceId
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) SetCreateTime(v string) *ListDDoSInstancesResponseBodyInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) SetDDoSBurstableDomesticProtection(v string) *ListDDoSInstancesResponseBodyInstanceInfo {
	s.DDoSBurstableDomesticProtection = &v
	return s
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) SetDDoSBurstableOverseasProtection(v string) *ListDDoSInstancesResponseBodyInstanceInfo {
	s.DDoSBurstableOverseasProtection = &v
	return s
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) SetInstanceId(v string) *ListDDoSInstancesResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) SetReserveReleaseTime(v string) *ListDDoSInstancesResponseBodyInstanceInfo {
	s.ReserveReleaseTime = &v
	return s
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) SetSiteInstanceId(v string) *ListDDoSInstancesResponseBodyInstanceInfo {
	s.SiteInstanceId = &v
	return s
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) SetStatus(v string) *ListDDoSInstancesResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *ListDDoSInstancesResponseBodyInstanceInfo) Validate() error {
	return dara.Validate(s)
}
