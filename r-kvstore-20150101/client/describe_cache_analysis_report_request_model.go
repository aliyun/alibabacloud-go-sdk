// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisType(v string) *DescribeCacheAnalysisReportRequest
	GetAnalysisType() *string
	SetDate(v string) *DescribeCacheAnalysisReportRequest
	GetDate() *string
	SetInstanceId(v string) *DescribeCacheAnalysisReportRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeCacheAnalysisReportRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeCacheAnalysisReportRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCacheAnalysisReportRequest
	GetOwnerId() *int64
	SetPageNumbers(v int32) *DescribeCacheAnalysisReportRequest
	GetPageNumbers() *int32
	SetPageSize(v int32) *DescribeCacheAnalysisReportRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCacheAnalysisReportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCacheAnalysisReportRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeCacheAnalysisReportRequest
	GetSecurityToken() *string
}

type DescribeCacheAnalysisReportRequest struct {
	// The type of analytics. Set the value to **BigKey**.
	//
	// This parameter is required.
	//
	// example:
	//
	// BigKey
	AnalysisType *string `json:"AnalysisType,omitempty" xml:"AnalysisType,omitempty"`
	// The date to query. You can query the report for one day each time. Specify the date in the *yyyy-MM-dd*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-08-05Z
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// -bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the child node in the cluster instance.
	//
	// >  If you do not specify this parameter, the analysis results of all child nodes in the instance are returned.
	//
	// example:
	//
	// -bp1zxszhcgatnx****-db-0
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// > If the parameter value exceeds the maximum number of the returned pages, an empty large key list is returned.
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

func (s DescribeCacheAnalysisReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportRequest) GetAnalysisType() *string {
	return s.AnalysisType
}

func (s *DescribeCacheAnalysisReportRequest) GetDate() *string {
	return s.Date
}

func (s *DescribeCacheAnalysisReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheAnalysisReportRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisReportRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCacheAnalysisReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCacheAnalysisReportRequest) GetPageNumbers() *int32 {
	return s.PageNumbers
}

func (s *DescribeCacheAnalysisReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCacheAnalysisReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCacheAnalysisReportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCacheAnalysisReportRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCacheAnalysisReportRequest) SetAnalysisType(v string) *DescribeCacheAnalysisReportRequest {
	s.AnalysisType = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetDate(v string) *DescribeCacheAnalysisReportRequest {
	s.Date = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetInstanceId(v string) *DescribeCacheAnalysisReportRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetNodeId(v string) *DescribeCacheAnalysisReportRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetOwnerAccount(v string) *DescribeCacheAnalysisReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetOwnerId(v int64) *DescribeCacheAnalysisReportRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetPageNumbers(v int32) *DescribeCacheAnalysisReportRequest {
	s.PageNumbers = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetPageSize(v int32) *DescribeCacheAnalysisReportRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetResourceOwnerAccount(v string) *DescribeCacheAnalysisReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetResourceOwnerId(v int64) *DescribeCacheAnalysisReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) SetSecurityToken(v string) *DescribeCacheAnalysisReportRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCacheAnalysisReportRequest) Validate() error {
	return dara.Validate(s)
}
