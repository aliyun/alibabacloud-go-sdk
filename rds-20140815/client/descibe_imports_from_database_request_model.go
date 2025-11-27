// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescibeImportsFromDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescibeImportsFromDatabaseRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescibeImportsFromDatabaseRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescibeImportsFromDatabaseRequest
	GetEndTime() *string
	SetEngine(v string) *DescibeImportsFromDatabaseRequest
	GetEngine() *string
	SetImportId(v int32) *DescibeImportsFromDatabaseRequest
	GetImportId() *int32
	SetOwnerAccount(v string) *DescibeImportsFromDatabaseRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescibeImportsFromDatabaseRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescibeImportsFromDatabaseRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescibeImportsFromDatabaseRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescibeImportsFromDatabaseRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescibeImportsFromDatabaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescibeImportsFromDatabaseRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescibeImportsFromDatabaseRequest
	GetStartTime() *string
}

type DescibeImportsFromDatabaseRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bpxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-06-11T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine of the instance. Set the value to **MySQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the migration task.
	//
	// example:
	//
	// 123
	ImportId     *int32  `json:"ImportId,omitempty" xml:"ImportId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: any non-zero positive integer.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. You can call the DescribeDBInstanceAttribute operation to obtain the ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-06-11T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescibeImportsFromDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescibeImportsFromDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DescibeImportsFromDatabaseRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescibeImportsFromDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescibeImportsFromDatabaseRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescibeImportsFromDatabaseRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescibeImportsFromDatabaseRequest) GetImportId() *int32 {
	return s.ImportId
}

func (s *DescibeImportsFromDatabaseRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescibeImportsFromDatabaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescibeImportsFromDatabaseRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescibeImportsFromDatabaseRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescibeImportsFromDatabaseRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescibeImportsFromDatabaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescibeImportsFromDatabaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescibeImportsFromDatabaseRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescibeImportsFromDatabaseRequest) SetClientToken(v string) *DescibeImportsFromDatabaseRequest {
	s.ClientToken = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetDBInstanceId(v string) *DescibeImportsFromDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetEndTime(v string) *DescibeImportsFromDatabaseRequest {
	s.EndTime = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetEngine(v string) *DescibeImportsFromDatabaseRequest {
	s.Engine = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetImportId(v int32) *DescibeImportsFromDatabaseRequest {
	s.ImportId = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetOwnerAccount(v string) *DescibeImportsFromDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetOwnerId(v int64) *DescibeImportsFromDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetPageNumber(v int32) *DescibeImportsFromDatabaseRequest {
	s.PageNumber = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetPageSize(v int32) *DescibeImportsFromDatabaseRequest {
	s.PageSize = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetResourceGroupId(v string) *DescibeImportsFromDatabaseRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetResourceOwnerAccount(v string) *DescibeImportsFromDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetResourceOwnerId(v int64) *DescibeImportsFromDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) SetStartTime(v string) *DescibeImportsFromDatabaseRequest {
	s.StartTime = &v
	return s
}

func (s *DescibeImportsFromDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
