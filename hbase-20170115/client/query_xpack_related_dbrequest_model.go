// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryXpackRelatedDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *QueryXpackRelatedDBRequest
	GetClusterId() *string
	SetOwnerId(v int64) *QueryXpackRelatedDBRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryXpackRelatedDBRequest
	GetRegionId() *string
	SetRelateDbType(v string) *QueryXpackRelatedDBRequest
	GetRelateDbType() *string
	SetResourceOwnerAccount(v string) *QueryXpackRelatedDBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryXpackRelatedDBRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *QueryXpackRelatedDBRequest
	GetZoneId() *string
}

type QueryXpackRelatedDBRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RelateDbType         *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryXpackRelatedDBRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelatedDBRequest) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryXpackRelatedDBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryXpackRelatedDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryXpackRelatedDBRequest) GetRelateDbType() *string {
	return s.RelateDbType
}

func (s *QueryXpackRelatedDBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryXpackRelatedDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryXpackRelatedDBRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *QueryXpackRelatedDBRequest) SetClusterId(v string) *QueryXpackRelatedDBRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetOwnerId(v int64) *QueryXpackRelatedDBRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetRegionId(v string) *QueryXpackRelatedDBRequest {
	s.RegionId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetRelateDbType(v string) *QueryXpackRelatedDBRequest {
	s.RelateDbType = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetResourceOwnerAccount(v string) *QueryXpackRelatedDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetResourceOwnerId(v int64) *QueryXpackRelatedDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetZoneId(v string) *QueryXpackRelatedDBRequest {
	s.ZoneId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) Validate() error {
	return dara.Validate(s)
}
