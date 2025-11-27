// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaskingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteMaskingRulesRequest
	GetDBInstanceName() *string
	SetDBName(v string) *DeleteMaskingRulesRequest
	GetDBName() *string
	SetOwnerId(v string) *DeleteMaskingRulesRequest
	GetOwnerId() *string
	SetRegionId(v string) *DeleteMaskingRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteMaskingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMaskingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleName(v string) *DeleteMaskingRulesRequest
	GetRuleName() *string
}

type DeleteMaskingRulesRequest struct {
	// This parameter is required.
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteMaskingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaskingRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaskingRulesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteMaskingRulesRequest) GetDBName() *string {
	return s.DBName
}

func (s *DeleteMaskingRulesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteMaskingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMaskingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMaskingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMaskingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteMaskingRulesRequest) SetDBInstanceName(v string) *DeleteMaskingRulesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetDBName(v string) *DeleteMaskingRulesRequest {
	s.DBName = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetOwnerId(v string) *DeleteMaskingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetRegionId(v string) *DeleteMaskingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetResourceOwnerAccount(v string) *DeleteMaskingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetResourceOwnerId(v int64) *DeleteMaskingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetRuleName(v string) *DeleteMaskingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteMaskingRulesRequest) Validate() error {
	return dara.Validate(s)
}
