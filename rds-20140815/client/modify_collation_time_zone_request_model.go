// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCollationTimeZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollation(v string) *ModifyCollationTimeZoneRequest
	GetCollation() *string
	SetDBInstanceId(v string) *ModifyCollationTimeZoneRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyCollationTimeZoneRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyCollationTimeZoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCollationTimeZoneRequest
	GetResourceOwnerId() *int64
	SetTimezone(v string) *ModifyCollationTimeZoneRequest
	GetTimezone() *string
}

type ModifyCollationTimeZoneRequest struct {
	Collation *string `json:"Collation,omitempty" xml:"Collation,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Timezone             *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s ModifyCollationTimeZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCollationTimeZoneRequest) GoString() string {
	return s.String()
}

func (s *ModifyCollationTimeZoneRequest) GetCollation() *string {
	return s.Collation
}

func (s *ModifyCollationTimeZoneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyCollationTimeZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCollationTimeZoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCollationTimeZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCollationTimeZoneRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *ModifyCollationTimeZoneRequest) SetCollation(v string) *ModifyCollationTimeZoneRequest {
	s.Collation = &v
	return s
}

func (s *ModifyCollationTimeZoneRequest) SetDBInstanceId(v string) *ModifyCollationTimeZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyCollationTimeZoneRequest) SetOwnerId(v int64) *ModifyCollationTimeZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCollationTimeZoneRequest) SetResourceOwnerAccount(v string) *ModifyCollationTimeZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCollationTimeZoneRequest) SetResourceOwnerId(v int64) *ModifyCollationTimeZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCollationTimeZoneRequest) SetTimezone(v string) *ModifyCollationTimeZoneRequest {
	s.Timezone = &v
	return s
}

func (s *ModifyCollationTimeZoneRequest) Validate() error {
	return dara.Validate(s)
}
