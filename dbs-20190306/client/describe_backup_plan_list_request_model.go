// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanListRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type DescribeBackupPlanListRequest struct {
	// The ID of the backup schedule. You can query multiple backup schedule IDs. Separate multiple IDs with commas (,).
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
	// Backup plan status, the values are as follows:
	//
	// 	- **wait**: Not configured
	//
	// 	- **init**: Not started (pre-check failed)
	//
	// 	- **running**: Running
	//
	// 	- **stop**: Failed
	//
	// 	- **pause**: Paused
	//
	// 	- **locked**: Locked
	//
	// 	- **check_pass**: Pre-check passed
	//
	// example:
	//
	// wait
	BackupPlanStatus *string `json:"BackupPlanStatus,omitempty" xml:"BackupPlanStatus,omitempty"`
	// Used to ensure the idempotence of the request, preventing duplicate submissions.
	//
	// example:
	//
	// ASDASDASDSADASFCZXVZ
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number, must be greater than or equal to 0 and not exceed the maximum value of Integer. The default value is 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of records per page, the value should be between 1 and 100.
	//
	// > The default is **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// DBS region, you can view the supported DBS regions by calling the [DescribeRegions](https://help.aliyun.com/document_detail/2869853.html) interface.
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
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeBackupPlanListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanListRequest) GoString() string {
	return s.String()
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

func (s *DescribeBackupPlanListRequest) Validate() error {
	return dara.Validate(s)
}
