// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *TransferVersionRequest
	GetDBClusterId() *string
	SetDisableWriteWindows(v string) *TransferVersionRequest
	GetDisableWriteWindows() *string
	SetOwnerAccount(v string) *TransferVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransferVersionRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *TransferVersionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *TransferVersionRequest
	GetPageSize() *int32
	SetRegionId(v string) *TransferVersionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *TransferVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransferVersionRequest
	GetResourceOwnerId() *int64
	SetSourceAccount(v string) *TransferVersionRequest
	GetSourceAccount() *string
	SetSourceClusterName(v string) *TransferVersionRequest
	GetSourceClusterName() *string
	SetSourcePassword(v string) *TransferVersionRequest
	GetSourcePassword() *string
	SetSourceShards(v string) *TransferVersionRequest
	GetSourceShards() *string
	SetTargetAccount(v string) *TransferVersionRequest
	GetTargetAccount() *string
	SetTargetDbClusterId(v string) *TransferVersionRequest
	GetTargetDbClusterId() *string
	SetTargetPassword(v string) *TransferVersionRequest
	GetTargetPassword() *string
}

type TransferVersionRequest struct {
	// The ID of the source ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1tm8zf130ew****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The time window during which write operations are stopped.
	//
	// example:
	//
	// 2025-02-08T00:00:00+08:00,2025-02-12T00:00:00+08:00
	DisableWriteWindows *string `json:"DisableWriteWindows,omitempty" xml:"DisableWriteWindows,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database account that is used to log on to the database in the source ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	SourceAccount *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	// The name of the self-managed ClickHouse cluster. You can execute the `SELECT 	- FROM system.clusters` statement to query the cluster name.
	//
	// example:
	//
	// default
	SourceClusterName *string `json:"SourceClusterName,omitempty" xml:"SourceClusterName,omitempty"`
	// The password that corresponds to the database account for logging on to the database in the source ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	SourcePassword *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	// The endpoint and TCP port of the self-managed ClickHouse cluster.
	//
	// example:
	//
	// 192.168.0.4:9000,192.168.0.5:9000,192.168.0.6:9000
	SourceShards *string `json:"SourceShards,omitempty" xml:"SourceShards,omitempty"`
	// The database account that is used to log on to the database in the destination ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	TargetAccount *string `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty"`
	// The ID of the destination ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp14b39djx7zg****
	TargetDbClusterId *string `json:"TargetDbClusterId,omitempty" xml:"TargetDbClusterId,omitempty"`
	// The password that corresponds to the database account for logging on to the database in the destination ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Ff
	TargetPassword *string `json:"TargetPassword,omitempty" xml:"TargetPassword,omitempty"`
}

func (s TransferVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferVersionRequest) GoString() string {
	return s.String()
}

func (s *TransferVersionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *TransferVersionRequest) GetDisableWriteWindows() *string {
	return s.DisableWriteWindows
}

func (s *TransferVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransferVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransferVersionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *TransferVersionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TransferVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TransferVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransferVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransferVersionRequest) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *TransferVersionRequest) GetSourceClusterName() *string {
	return s.SourceClusterName
}

func (s *TransferVersionRequest) GetSourcePassword() *string {
	return s.SourcePassword
}

func (s *TransferVersionRequest) GetSourceShards() *string {
	return s.SourceShards
}

func (s *TransferVersionRequest) GetTargetAccount() *string {
	return s.TargetAccount
}

func (s *TransferVersionRequest) GetTargetDbClusterId() *string {
	return s.TargetDbClusterId
}

func (s *TransferVersionRequest) GetTargetPassword() *string {
	return s.TargetPassword
}

func (s *TransferVersionRequest) SetDBClusterId(v string) *TransferVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *TransferVersionRequest) SetDisableWriteWindows(v string) *TransferVersionRequest {
	s.DisableWriteWindows = &v
	return s
}

func (s *TransferVersionRequest) SetOwnerAccount(v string) *TransferVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransferVersionRequest) SetOwnerId(v int64) *TransferVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *TransferVersionRequest) SetPageNumber(v int32) *TransferVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *TransferVersionRequest) SetPageSize(v int32) *TransferVersionRequest {
	s.PageSize = &v
	return s
}

func (s *TransferVersionRequest) SetRegionId(v string) *TransferVersionRequest {
	s.RegionId = &v
	return s
}

func (s *TransferVersionRequest) SetResourceOwnerAccount(v string) *TransferVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransferVersionRequest) SetResourceOwnerId(v int64) *TransferVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransferVersionRequest) SetSourceAccount(v string) *TransferVersionRequest {
	s.SourceAccount = &v
	return s
}

func (s *TransferVersionRequest) SetSourceClusterName(v string) *TransferVersionRequest {
	s.SourceClusterName = &v
	return s
}

func (s *TransferVersionRequest) SetSourcePassword(v string) *TransferVersionRequest {
	s.SourcePassword = &v
	return s
}

func (s *TransferVersionRequest) SetSourceShards(v string) *TransferVersionRequest {
	s.SourceShards = &v
	return s
}

func (s *TransferVersionRequest) SetTargetAccount(v string) *TransferVersionRequest {
	s.TargetAccount = &v
	return s
}

func (s *TransferVersionRequest) SetTargetDbClusterId(v string) *TransferVersionRequest {
	s.TargetDbClusterId = &v
	return s
}

func (s *TransferVersionRequest) SetTargetPassword(v string) *TransferVersionRequest {
	s.TargetPassword = &v
	return s
}

func (s *TransferVersionRequest) Validate() error {
	return dara.Validate(s)
}
