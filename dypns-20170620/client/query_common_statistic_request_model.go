// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QueryCommonStatisticRequest
	GetEndDate() *string
	SetOwnerId(v int64) *QueryCommonStatisticRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryCommonStatisticRequest
	GetProdCode() *string
	SetProductType(v int32) *QueryCommonStatisticRequest
	GetProductType() *int32
	SetResourceOwnerAccount(v string) *QueryCommonStatisticRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCommonStatisticRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *QueryCommonStatisticRequest
	GetStartDate() *string
}

type QueryCommonStatisticRequest struct {
	// This parameter is required.
	EndDate  *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// This parameter is required.
	ProductType          *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryCommonStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonStatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryCommonStatisticRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryCommonStatisticRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCommonStatisticRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryCommonStatisticRequest) GetProductType() *int32 {
	return s.ProductType
}

func (s *QueryCommonStatisticRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCommonStatisticRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCommonStatisticRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryCommonStatisticRequest) SetEndDate(v string) *QueryCommonStatisticRequest {
	s.EndDate = &v
	return s
}

func (s *QueryCommonStatisticRequest) SetOwnerId(v int64) *QueryCommonStatisticRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCommonStatisticRequest) SetProdCode(v string) *QueryCommonStatisticRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCommonStatisticRequest) SetProductType(v int32) *QueryCommonStatisticRequest {
	s.ProductType = &v
	return s
}

func (s *QueryCommonStatisticRequest) SetResourceOwnerAccount(v string) *QueryCommonStatisticRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCommonStatisticRequest) SetResourceOwnerId(v int64) *QueryCommonStatisticRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCommonStatisticRequest) SetStartDate(v string) *QueryCommonStatisticRequest {
	s.StartDate = &v
	return s
}

func (s *QueryCommonStatisticRequest) Validate() error {
	return dara.Validate(s)
}
