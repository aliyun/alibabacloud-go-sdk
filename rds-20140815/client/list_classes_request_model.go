// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClassesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListClassesRequest
	GetClientToken() *string
	SetCommodityCode(v string) *ListClassesRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *ListClassesRequest
	GetDBInstanceId() *string
	SetEngine(v string) *ListClassesRequest
	GetEngine() *string
	SetOrderType(v string) *ListClassesRequest
	GetOrderType() *string
	SetOwnerId(v int64) *ListClassesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListClassesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListClassesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListClassesRequest
	GetResourceOwnerId() *int64
}

type ListClassesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListClassesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClassesRequest) GoString() string {
	return s.String()
}

func (s *ListClassesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListClassesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListClassesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListClassesRequest) GetEngine() *string {
	return s.Engine
}

func (s *ListClassesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListClassesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListClassesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClassesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListClassesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListClassesRequest) SetClientToken(v string) *ListClassesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListClassesRequest) SetCommodityCode(v string) *ListClassesRequest {
	s.CommodityCode = &v
	return s
}

func (s *ListClassesRequest) SetDBInstanceId(v string) *ListClassesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListClassesRequest) SetEngine(v string) *ListClassesRequest {
	s.Engine = &v
	return s
}

func (s *ListClassesRequest) SetOrderType(v string) *ListClassesRequest {
	s.OrderType = &v
	return s
}

func (s *ListClassesRequest) SetOwnerId(v int64) *ListClassesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListClassesRequest) SetRegionId(v string) *ListClassesRequest {
	s.RegionId = &v
	return s
}

func (s *ListClassesRequest) SetResourceOwnerAccount(v string) *ListClassesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListClassesRequest) SetResourceOwnerId(v int64) *ListClassesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClassesRequest) Validate() error {
	return dara.Validate(s)
}
