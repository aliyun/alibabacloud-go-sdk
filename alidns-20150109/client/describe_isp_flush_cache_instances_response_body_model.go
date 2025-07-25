// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIspFlushCacheInstances(v []*DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) *DescribeIspFlushCacheInstancesResponseBody
	GetIspFlushCacheInstances() []*DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances
	SetPageNumber(v int32) *DescribeIspFlushCacheInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIspFlushCacheInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIspFlushCacheInstancesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeIspFlushCacheInstancesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeIspFlushCacheInstancesResponseBody
	GetTotalPages() *int32
}

type DescribeIspFlushCacheInstancesResponseBody struct {
	IspFlushCacheInstances []*DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances `json:"IspFlushCacheInstances,omitempty" xml:"IspFlushCacheInstances,omitempty" type:"Repeated"`
	PageNumber             *int32                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItems             *int32                                                              `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages             *int32                                                              `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeIspFlushCacheInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheInstancesResponseBody) GetIspFlushCacheInstances() []*DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	return s.IspFlushCacheInstances
}

func (s *DescribeIspFlushCacheInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIspFlushCacheInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIspFlushCacheInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIspFlushCacheInstancesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeIspFlushCacheInstancesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeIspFlushCacheInstancesResponseBody) SetIspFlushCacheInstances(v []*DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) *DescribeIspFlushCacheInstancesResponseBody {
	s.IspFlushCacheInstances = v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBody) SetPageNumber(v int32) *DescribeIspFlushCacheInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBody) SetPageSize(v int32) *DescribeIspFlushCacheInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBody) SetRequestId(v string) *DescribeIspFlushCacheInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBody) SetTotalItems(v int32) *DescribeIspFlushCacheInstancesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBody) SetTotalPages(v int32) *DescribeIspFlushCacheInstancesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances struct {
	ExpireTime      *string                                                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpireTimestamp *int64                                                                     `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	InstanceId      *string                                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string                                                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Isp             *string                                                                    `json:"Isp,omitempty" xml:"Isp,omitempty"`
	QuotaInfo       *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo `json:"QuotaInfo,omitempty" xml:"QuotaInfo,omitempty" type:"Struct"`
	Status          *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionCode     *string                                                                    `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetIsp() *string {
	return s.Isp
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetQuotaInfo() *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo {
	return s.QuotaInfo
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetExpireTime(v string) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetExpireTimestamp(v int64) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetInstanceId(v string) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetInstanceName(v string) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetIsp(v string) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.Isp = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetQuotaInfo(v *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.QuotaInfo = v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetStatus(v string) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.Status = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) SetVersionCode(v string) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances {
	s.VersionCode = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo struct {
	InstanceQuota     *int32 `json:"InstanceQuota,omitempty" xml:"InstanceQuota,omitempty"`
	InstanceQuotaUsed *int32 `json:"InstanceQuotaUsed,omitempty" xml:"InstanceQuotaUsed,omitempty"`
}

func (s DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) GetInstanceQuota() *int32 {
	return s.InstanceQuota
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) GetInstanceQuotaUsed() *int32 {
	return s.InstanceQuotaUsed
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) SetInstanceQuota(v int32) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo {
	s.InstanceQuota = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) SetInstanceQuotaUsed(v int32) *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo {
	s.InstanceQuotaUsed = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesResponseBodyIspFlushCacheInstancesQuotaInfo) Validate() error {
	return dara.Validate(s)
}
