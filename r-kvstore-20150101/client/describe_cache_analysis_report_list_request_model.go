// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDays(v int32) *DescribeCacheAnalysisReportListRequest
	GetDays() *int32
	SetInstanceId(v string) *DescribeCacheAnalysisReportListRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeCacheAnalysisReportListRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeCacheAnalysisReportListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCacheAnalysisReportListRequest
	GetOwnerId() *int64
	SetPageNumbers(v int32) *DescribeCacheAnalysisReportListRequest
	GetPageNumbers() *int32
	SetPageSize(v int32) *DescribeCacheAnalysisReportListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCacheAnalysisReportListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCacheAnalysisReportListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeCacheAnalysisReportListRequest
	GetSecurityToken() *string
}

type DescribeCacheAnalysisReportListRequest struct {
	// The time range to query. Default value: 7. Unit: days.
	//
	// > If daily automatic analysis has not started and manual analysis is not performed, no records are returned.
	//
	// example:
	//
	// 7
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the child node in the cluster instance.
	//
	// >  If you do not specify this parameter, the analysis results of all child nodes in the instance are returned.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumbers *int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**.
	//
	// > The default value is **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCacheAnalysisReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListRequest) GetDays() *int32 {
	return s.Days
}

func (s *DescribeCacheAnalysisReportListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheAnalysisReportListRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisReportListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCacheAnalysisReportListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCacheAnalysisReportListRequest) GetPageNumbers() *int32 {
	return s.PageNumbers
}

func (s *DescribeCacheAnalysisReportListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCacheAnalysisReportListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCacheAnalysisReportListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCacheAnalysisReportListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCacheAnalysisReportListRequest) SetDays(v int32) *DescribeCacheAnalysisReportListRequest {
	s.Days = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetInstanceId(v string) *DescribeCacheAnalysisReportListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetNodeId(v string) *DescribeCacheAnalysisReportListRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetOwnerAccount(v string) *DescribeCacheAnalysisReportListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetOwnerId(v int64) *DescribeCacheAnalysisReportListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetPageNumbers(v int32) *DescribeCacheAnalysisReportListRequest {
	s.PageNumbers = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetPageSize(v int32) *DescribeCacheAnalysisReportListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetResourceOwnerAccount(v string) *DescribeCacheAnalysisReportListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetResourceOwnerId(v int64) *DescribeCacheAnalysisReportListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) SetSecurityToken(v string) *DescribeCacheAnalysisReportListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCacheAnalysisReportListRequest) Validate() error {
	return dara.Validate(s)
}
