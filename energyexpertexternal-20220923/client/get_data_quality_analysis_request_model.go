// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataQualityAnalysisRequest
	GetCode() *string
	SetDataQualityEvaluationType(v int64) *GetDataQualityAnalysisRequest
	GetDataQualityEvaluationType() *int64
	SetProductId(v int64) *GetDataQualityAnalysisRequest
	GetProductId() *int64
	SetProductType(v int64) *GetDataQualityAnalysisRequest
	GetProductType() *int64
}

type GetDataQualityAnalysisRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Data quality assessment type: 1 is DQI and 2 is DQR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataQualityEvaluationType *int64 `json:"dataQualityEvaluationType,omitempty" xml:"dataQualityEvaluationType,omitempty"`
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

func (s GetDataQualityAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisRequest) GetCode() *string {
	return s.Code
}

func (s *GetDataQualityAnalysisRequest) GetDataQualityEvaluationType() *int64 {
	return s.DataQualityEvaluationType
}

func (s *GetDataQualityAnalysisRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetDataQualityAnalysisRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetDataQualityAnalysisRequest) SetCode(v string) *GetDataQualityAnalysisRequest {
	s.Code = &v
	return s
}

func (s *GetDataQualityAnalysisRequest) SetDataQualityEvaluationType(v int64) *GetDataQualityAnalysisRequest {
	s.DataQualityEvaluationType = &v
	return s
}

func (s *GetDataQualityAnalysisRequest) SetProductId(v int64) *GetDataQualityAnalysisRequest {
	s.ProductId = &v
	return s
}

func (s *GetDataQualityAnalysisRequest) SetProductType(v int64) *GetDataQualityAnalysisRequest {
	s.ProductType = &v
	return s
}

func (s *GetDataQualityAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
