// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeParameterTemplatesRequest
	GetCategory() *string
	SetClientToken(v string) *DescribeParameterTemplatesRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeParameterTemplatesRequest
	GetDBInstanceId() *string
	SetEngine(v string) *DescribeParameterTemplatesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeParameterTemplatesRequest
	GetEngineVersion() *string
	SetOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeParameterTemplatesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetResourceOwnerId() *int64
}

type DescribeParameterTemplatesRequest struct {
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// Basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-bp1imnm**********
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- **mysql**: MySQL
	//
	// 	- **mssql**: SQL Server
	//
	// 	- **PostgreSQL**: PostgreSQL
	//
	// 	- **MariaDB**: MariaDB
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine. Valid values:
	//
	// 	- Valid values when you set the Engine parameter to mysql: **5.5, 5.6, 5.7, and 8.0**.
	//
	// 	- Valid values when you set the Engine parameter to mssql: **2008r2**.
	//
	// 	- Valid values when you set the Engine parameter to PostgreSQL: **10.0, 11.0, 12.0, 13.0, 14.0, and 15.0**.
	//
	// 	- Valid values when you set the Engine parameter to MariaDB: **10.3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeParameterTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeParameterTemplatesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeParameterTemplatesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeParameterTemplatesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterTemplatesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterTemplatesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterTemplatesRequest) SetCategory(v string) *DescribeParameterTemplatesRequest {
	s.Category = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetClientToken(v string) *DescribeParameterTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetDBInstanceId(v string) *DescribeParameterTemplatesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngine(v string) *DescribeParameterTemplatesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngineVersion(v string) *DescribeParameterTemplatesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetRegionId(v string) *DescribeParameterTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
