// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLServerUpgradeVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeSQLServerUpgradeVersionsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeSQLServerUpgradeVersionsRequest
	GetDBInstanceId() *string
	SetEngineVersion(v string) *DescribeSQLServerUpgradeVersionsRequest
	GetEngineVersion() *string
	SetOwnerId(v int64) *DescribeSQLServerUpgradeVersionsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSQLServerUpgradeVersionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLServerUpgradeVersionsRequest
	GetResourceOwnerId() *int64
}

type DescribeSQLServerUpgradeVersionsRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// rm-xx1xxx2****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 2016_web
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQLServerUpgradeVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSQLServerUpgradeVersionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLServerUpgradeVersionsRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeSQLServerUpgradeVersionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLServerUpgradeVersionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLServerUpgradeVersionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLServerUpgradeVersionsRequest) SetClientToken(v string) *DescribeSQLServerUpgradeVersionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsRequest) SetDBInstanceId(v string) *DescribeSQLServerUpgradeVersionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsRequest) SetEngineVersion(v string) *DescribeSQLServerUpgradeVersionsRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsRequest) SetOwnerId(v int64) *DescribeSQLServerUpgradeVersionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsRequest) SetResourceOwnerAccount(v string) *DescribeSQLServerUpgradeVersionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsRequest) SetResourceOwnerId(v int64) *DescribeSQLServerUpgradeVersionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsRequest) Validate() error {
	return dara.Validate(s)
}
