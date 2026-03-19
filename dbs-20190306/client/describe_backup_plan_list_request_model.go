// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMethod(v string) *DescribeBackupPlanListRequest
	GetBackupMethod() *string
	SetBackupPlanId(v string) *DescribeBackupPlanListRequest
	GetBackupPlanId() *string
	SetBackupPlanName(v string) *DescribeBackupPlanListRequest
	GetBackupPlanName() *string
	SetBackupPlanStatus(v string) *DescribeBackupPlanListRequest
	GetBackupPlanStatus() *string
	SetClientToken(v string) *DescribeBackupPlanListRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeBackupPlanListRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeBackupPlanListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeBackupPlanListRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeBackupPlanListRequest
	GetRegion() *string
	SetResourceGroupId(v string) *DescribeBackupPlanListRequest
	GetResourceGroupId() *string
	SetShowBackupStrategyInfo(v bool) *DescribeBackupPlanListRequest
	GetShowBackupStrategyInfo() *bool
	SetShowRecoverTimeRange(v bool) *DescribeBackupPlanListRequest
	GetShowRecoverTimeRange() *bool
	SetShowStorageStrategyInfo(v bool) *DescribeBackupPlanListRequest
	GetShowStorageStrategyInfo() *bool
}

type DescribeBackupPlanListRequest struct {
	// example:
	//
	// logical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// Backup plan ID. To list multiple backup plans, separate IDs with commas.
	//
	// example:
	//
	// dbstooi01exXXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// Backup plan name.
	//
	// example:
	//
	// test123
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// Backup plan status. Valid values:
	//
	// - **wait**: Not configured.
	//
	// - **init**: Not started (precheck failed).
	//
	// - **running**: Running.
	//
	// - **stop**: Failed.
	//
	// - **pause**: Paused.
	//
	// - **locked**: Locked.
	//
	// - **check_pass**: Precheck passed.
	//
	// example:
	//
	// wait
	BackupPlanStatus *string `json:"BackupPlanStatus,omitempty" xml:"BackupPlanStatus,omitempty"`
	// A client token used to ensure idempotence and prevent duplicate requests.
	//
	// example:
	//
	// ASDASDASDSADASFCZXVZ
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number. Valid values: integers greater than or equal to 0 and less than or equal to the maximum integer value. Default value: 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of records per page. Valid values: 1 to 100.
	//
	// > Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// DBS region. Call [DescribeRegions](https://help.aliyun.com/document_detail/2869853.html) to view supported regions.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekzecovzti****
	ResourceGroupId         *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ShowBackupStrategyInfo  *bool   `json:"ShowBackupStrategyInfo,omitempty" xml:"ShowBackupStrategyInfo,omitempty"`
	ShowRecoverTimeRange    *bool   `json:"ShowRecoverTimeRange,omitempty" xml:"ShowRecoverTimeRange,omitempty"`
	ShowStorageStrategyInfo *bool   `json:"ShowStorageStrategyInfo,omitempty" xml:"ShowStorageStrategyInfo,omitempty"`
}

func (s DescribeBackupPlanListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupPlanListRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeBackupPlanListRequest) GetBackupPlanName() *string {
	return s.BackupPlanName
}

func (s *DescribeBackupPlanListRequest) GetBackupPlanStatus() *string {
	return s.BackupPlanStatus
}

func (s *DescribeBackupPlanListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeBackupPlanListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeBackupPlanListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeBackupPlanListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupPlanListRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeBackupPlanListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBackupPlanListRequest) GetShowBackupStrategyInfo() *bool {
	return s.ShowBackupStrategyInfo
}

func (s *DescribeBackupPlanListRequest) GetShowRecoverTimeRange() *bool {
	return s.ShowRecoverTimeRange
}

func (s *DescribeBackupPlanListRequest) GetShowStorageStrategyInfo() *bool {
	return s.ShowStorageStrategyInfo
}

func (s *DescribeBackupPlanListRequest) SetBackupMethod(v string) *DescribeBackupPlanListRequest {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetBackupPlanId(v string) *DescribeBackupPlanListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetBackupPlanName(v string) *DescribeBackupPlanListRequest {
	s.BackupPlanName = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetBackupPlanStatus(v string) *DescribeBackupPlanListRequest {
	s.BackupPlanStatus = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetClientToken(v string) *DescribeBackupPlanListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetOwnerId(v string) *DescribeBackupPlanListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetPageNum(v int32) *DescribeBackupPlanListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetPageSize(v int32) *DescribeBackupPlanListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetRegion(v string) *DescribeBackupPlanListRequest {
	s.Region = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetResourceGroupId(v string) *DescribeBackupPlanListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetShowBackupStrategyInfo(v bool) *DescribeBackupPlanListRequest {
	s.ShowBackupStrategyInfo = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetShowRecoverTimeRange(v bool) *DescribeBackupPlanListRequest {
	s.ShowRecoverTimeRange = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetShowStorageStrategyInfo(v bool) *DescribeBackupPlanListRequest {
	s.ShowStorageStrategyInfo = &v
	return s
}

func (s *DescribeBackupPlanListRequest) Validate() error {
	return dara.Validate(s)
}
