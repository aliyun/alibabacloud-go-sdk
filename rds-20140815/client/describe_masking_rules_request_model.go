// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaskingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeMaskingRulesRequest
	GetDBInstanceName() *string
	SetDBName(v string) *DescribeMaskingRulesRequest
	GetDBName() *string
	SetOwnerId(v string) *DescribeMaskingRulesRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeMaskingRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeMaskingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMaskingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleName(v string) *DescribeMaskingRulesRequest
	GetRuleName() *string
}

type DescribeMaskingRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n8t18o******5
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// myDB
	DBName  *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// test1,test2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeMaskingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeMaskingRulesRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeMaskingRulesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeMaskingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMaskingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMaskingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMaskingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeMaskingRulesRequest) SetDBInstanceName(v string) *DescribeMaskingRulesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetDBName(v string) *DescribeMaskingRulesRequest {
	s.DBName = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetOwnerId(v string) *DescribeMaskingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetRegionId(v string) *DescribeMaskingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetResourceOwnerAccount(v string) *DescribeMaskingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetResourceOwnerId(v int64) *DescribeMaskingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetRuleName(v string) *DescribeMaskingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeMaskingRulesRequest) Validate() error {
	return dara.Validate(s)
}
