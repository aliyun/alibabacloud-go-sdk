// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int32) *DescribeDataLimitsRequest
	GetAuditStatus() *int32
	SetCheckStatus(v int32) *DescribeDataLimitsRequest
	GetCheckStatus() *int32
	SetCurrentPage(v int32) *DescribeDataLimitsRequest
	GetCurrentPage() *int32
	SetDatamaskStatus(v int32) *DescribeDataLimitsRequest
	GetDatamaskStatus() *int32
	SetEnable(v int32) *DescribeDataLimitsRequest
	GetEnable() *int32
	SetEndTime(v int64) *DescribeDataLimitsRequest
	GetEndTime() *int64
	SetEngineType(v string) *DescribeDataLimitsRequest
	GetEngineType() *string
	SetFeatureType(v int32) *DescribeDataLimitsRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeDataLimitsRequest
	GetLang() *string
	SetMemberAccount(v int64) *DescribeDataLimitsRequest
	GetMemberAccount() *int64
	SetPageSize(v int32) *DescribeDataLimitsRequest
	GetPageSize() *int32
	SetParentId(v string) *DescribeDataLimitsRequest
	GetParentId() *string
	SetResourceType(v int32) *DescribeDataLimitsRequest
	GetResourceType() *int32
	SetServiceRegionId(v string) *DescribeDataLimitsRequest
	GetServiceRegionId() *string
	SetStartTime(v int64) *DescribeDataLimitsRequest
	GetStartTime() *int64
}

type DescribeDataLimitsRequest struct {
	// The audit status. Valid values:
	//
	// - **1**: audit enabled.
	//
	// - **0**: audit disabled.
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The data detection status. Valid values:
	//
	// - **0**: ready.
	//
	// - **1**: running.
	//
	// - **2**: connectivity test in progress.
	//
	// - **3**: connectivity test passed.
	//
	// - **4**: connectivity test failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// Settings the page number of the current page when you perform a paged query. For more information about paging, refer to the paging parameters.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data masking permission status. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// example:
	//
	// 1
	DatamaskStatus *int32 `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	// The detection permission status. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end of the creation time range. Format: timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1616068534877
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database type. Valid values: **MySQL**, **SQLServer**, **Oracle**, **PostgreSQL**, **MongoDB**, and others.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the member accounts.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// The maximum number of entries to return on each page when you perform a paged query. For more information about paging, refer to the paging parameters.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent asset to which the data asset belongs. Valid values:
	//
	// - The project name or ID for MaxCompute.
	//
	// - The bucket name or ID for OSS.
	//
	// - The instance name or ID, or the database name or ID for RDS.
	//
	// example:
	//
	// 1112
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// Required. The type of the product to which the data asset belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADS
	//
	// - **4**: OTS
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The region where the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The start of the creation time range. Format: timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1616068534877
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataLimitsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsRequest) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *DescribeDataLimitsRequest) GetCheckStatus() *int32 {
	return s.CheckStatus
}

func (s *DescribeDataLimitsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataLimitsRequest) GetDatamaskStatus() *int32 {
	return s.DatamaskStatus
}

func (s *DescribeDataLimitsRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *DescribeDataLimitsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDataLimitsRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeDataLimitsRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeDataLimitsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataLimitsRequest) GetMemberAccount() *int64 {
	return s.MemberAccount
}

func (s *DescribeDataLimitsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataLimitsRequest) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDataLimitsRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *DescribeDataLimitsRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeDataLimitsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDataLimitsRequest) SetAuditStatus(v int32) *DescribeDataLimitsRequest {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetCheckStatus(v int32) *DescribeDataLimitsRequest {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetCurrentPage(v int32) *DescribeDataLimitsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetDatamaskStatus(v int32) *DescribeDataLimitsRequest {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetEnable(v int32) *DescribeDataLimitsRequest {
	s.Enable = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetEndTime(v int64) *DescribeDataLimitsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetEngineType(v string) *DescribeDataLimitsRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetFeatureType(v int32) *DescribeDataLimitsRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetLang(v string) *DescribeDataLimitsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetMemberAccount(v int64) *DescribeDataLimitsRequest {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetPageSize(v int32) *DescribeDataLimitsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetParentId(v string) *DescribeDataLimitsRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetResourceType(v int32) *DescribeDataLimitsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetServiceRegionId(v string) *DescribeDataLimitsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetStartTime(v int64) *DescribeDataLimitsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataLimitsRequest) Validate() error {
	return dara.Validate(s)
}
