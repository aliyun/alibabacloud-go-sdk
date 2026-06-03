// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonStatisticPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QueryCommonStatisticPreviewRequest
	GetEndDate() *string
	SetOwnerId(v int64) *QueryCommonStatisticPreviewRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryCommonStatisticPreviewRequest
	GetProdCode() *string
	SetProductType(v int32) *QueryCommonStatisticPreviewRequest
	GetProductType() *int32
	SetResourceOwnerAccount(v string) *QueryCommonStatisticPreviewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCommonStatisticPreviewRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *QueryCommonStatisticPreviewRequest
	GetStartDate() *string
}

type QueryCommonStatisticPreviewRequest struct {
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

func (s QueryCommonStatisticPreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonStatisticPreviewRequest) GoString() string {
	return s.String()
}

func (s *QueryCommonStatisticPreviewRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryCommonStatisticPreviewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCommonStatisticPreviewRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryCommonStatisticPreviewRequest) GetProductType() *int32 {
	return s.ProductType
}

func (s *QueryCommonStatisticPreviewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCommonStatisticPreviewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCommonStatisticPreviewRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryCommonStatisticPreviewRequest) SetEndDate(v string) *QueryCommonStatisticPreviewRequest {
	s.EndDate = &v
	return s
}

func (s *QueryCommonStatisticPreviewRequest) SetOwnerId(v int64) *QueryCommonStatisticPreviewRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCommonStatisticPreviewRequest) SetProdCode(v string) *QueryCommonStatisticPreviewRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCommonStatisticPreviewRequest) SetProductType(v int32) *QueryCommonStatisticPreviewRequest {
	s.ProductType = &v
	return s
}

func (s *QueryCommonStatisticPreviewRequest) SetResourceOwnerAccount(v string) *QueryCommonStatisticPreviewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCommonStatisticPreviewRequest) SetResourceOwnerId(v int64) *QueryCommonStatisticPreviewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCommonStatisticPreviewRequest) SetStartDate(v string) *QueryCommonStatisticPreviewRequest {
	s.StartDate = &v
	return s
}

func (s *QueryCommonStatisticPreviewRequest) Validate() error {
	return dara.Validate(s)
}
