// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReductionProposalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetReductionProposalRequest
	GetCode() *string
	SetDataQualityEvaluationType(v int32) *GetReductionProposalRequest
	GetDataQualityEvaluationType() *int32
	SetProductId(v int64) *GetReductionProposalRequest
	GetProductId() *int64
	SetProductType(v int64) *GetReductionProposalRequest
	GetProductType() *int64
}

type GetReductionProposalRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The type of the data quality evaluation. 1 is DQI and 2 is DQR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataQualityEvaluationType *int32 `json:"dataQualityEvaluationType,omitempty" xml:"dataQualityEvaluationType,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetReductionProposalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReductionProposalRequest) GoString() string {
	return s.String()
}

func (s *GetReductionProposalRequest) GetCode() *string {
	return s.Code
}

func (s *GetReductionProposalRequest) GetDataQualityEvaluationType() *int32 {
	return s.DataQualityEvaluationType
}

func (s *GetReductionProposalRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetReductionProposalRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetReductionProposalRequest) SetCode(v string) *GetReductionProposalRequest {
	s.Code = &v
	return s
}

func (s *GetReductionProposalRequest) SetDataQualityEvaluationType(v int32) *GetReductionProposalRequest {
	s.DataQualityEvaluationType = &v
	return s
}

func (s *GetReductionProposalRequest) SetProductId(v int64) *GetReductionProposalRequest {
	s.ProductId = &v
	return s
}

func (s *GetReductionProposalRequest) SetProductType(v int64) *GetReductionProposalRequest {
	s.ProductType = &v
	return s
}

func (s *GetReductionProposalRequest) Validate() error {
	return dara.Validate(s)
}
