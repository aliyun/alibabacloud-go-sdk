// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiduiAreaDeductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProducts(v []*DiduiAreaDeductionRequestProducts) *DiduiAreaDeductionRequest
	GetProducts() []*DiduiAreaDeductionRequestProducts
	SetRagId(v string) *DiduiAreaDeductionRequest
	GetRagId() *string
	SetReqId(v string) *DiduiAreaDeductionRequest
	GetReqId() *string
	SetTargetImageUrl(v string) *DiduiAreaDeductionRequest
	GetTargetImageUrl() *string
}

type DiduiAreaDeductionRequest struct {
	// The list of products and their detection boxes.
	//
	// This parameter is required.
	Products []*DiduiAreaDeductionRequestProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// The ID of the customer-specific SKU vector database that determines which database is used for retrieval. The database must be created in advance through the database creation process.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The optional business request ID used for Tracing Analysis.
	//
	// example:
	//
	// didui-request-001
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The HTTPS URL of the overall floor display image.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/didui.jpg
	TargetImageUrl *string `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
}

func (s DiduiAreaDeductionRequest) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionRequest) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionRequest) GetProducts() []*DiduiAreaDeductionRequestProducts {
	return s.Products
}

func (s *DiduiAreaDeductionRequest) GetRagId() *string {
	return s.RagId
}

func (s *DiduiAreaDeductionRequest) GetReqId() *string {
	return s.ReqId
}

func (s *DiduiAreaDeductionRequest) GetTargetImageUrl() *string {
	return s.TargetImageUrl
}

func (s *DiduiAreaDeductionRequest) SetProducts(v []*DiduiAreaDeductionRequestProducts) *DiduiAreaDeductionRequest {
	s.Products = v
	return s
}

func (s *DiduiAreaDeductionRequest) SetRagId(v string) *DiduiAreaDeductionRequest {
	s.RagId = &v
	return s
}

func (s *DiduiAreaDeductionRequest) SetReqId(v string) *DiduiAreaDeductionRequest {
	s.ReqId = &v
	return s
}

func (s *DiduiAreaDeductionRequest) SetTargetImageUrl(v string) *DiduiAreaDeductionRequest {
	s.TargetImageUrl = &v
	return s
}

func (s *DiduiAreaDeductionRequest) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DiduiAreaDeductionRequestProducts struct {
	// The detection boxes for the current SKU. Coordinate values range from 0 to 1000.
	//
	// This parameter is required.
	Boxes []*DiduiAreaDeductionRequestProductsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// The unique ID of the SKU.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6901234579
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s DiduiAreaDeductionRequestProducts) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionRequestProducts) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionRequestProducts) GetBoxes() []*DiduiAreaDeductionRequestProductsBoxes {
	return s.Boxes
}

func (s *DiduiAreaDeductionRequestProducts) GetSkuId() *string {
	return s.SkuId
}

func (s *DiduiAreaDeductionRequestProducts) SetBoxes(v []*DiduiAreaDeductionRequestProductsBoxes) *DiduiAreaDeductionRequestProducts {
	s.Boxes = v
	return s
}

func (s *DiduiAreaDeductionRequestProducts) SetSkuId(v string) *DiduiAreaDeductionRequestProducts {
	s.SkuId = &v
	return s
}

func (s *DiduiAreaDeductionRequestProducts) Validate() error {
	if s.Boxes != nil {
		for _, item := range s.Boxes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DiduiAreaDeductionRequestProductsBoxes struct {
	// The bottom boundary of the detection box.
	//
	// This parameter is required.
	//
	// example:
	//
	// 689
	Bottom *float64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// The left boundary of the detection box.
	//
	// This parameter is required.
	//
	// example:
	//
	// 763
	Left *float64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The right boundary of the detection box.
	//
	// This parameter is required.
	//
	// example:
	//
	// 904
	Right *float64 `json:"Right,omitempty" xml:"Right,omitempty"`
	// The top boundary of the detection box.
	//
	// This parameter is required.
	//
	// example:
	//
	// 606
	Top *float64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DiduiAreaDeductionRequestProductsBoxes) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionRequestProductsBoxes) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionRequestProductsBoxes) GetBottom() *float64 {
	return s.Bottom
}

func (s *DiduiAreaDeductionRequestProductsBoxes) GetLeft() *float64 {
	return s.Left
}

func (s *DiduiAreaDeductionRequestProductsBoxes) GetRight() *float64 {
	return s.Right
}

func (s *DiduiAreaDeductionRequestProductsBoxes) GetTop() *float64 {
	return s.Top
}

func (s *DiduiAreaDeductionRequestProductsBoxes) SetBottom(v float64) *DiduiAreaDeductionRequestProductsBoxes {
	s.Bottom = &v
	return s
}

func (s *DiduiAreaDeductionRequestProductsBoxes) SetLeft(v float64) *DiduiAreaDeductionRequestProductsBoxes {
	s.Left = &v
	return s
}

func (s *DiduiAreaDeductionRequestProductsBoxes) SetRight(v float64) *DiduiAreaDeductionRequestProductsBoxes {
	s.Right = &v
	return s
}

func (s *DiduiAreaDeductionRequestProductsBoxes) SetTop(v float64) *DiduiAreaDeductionRequestProductsBoxes {
	s.Top = &v
	return s
}

func (s *DiduiAreaDeductionRequestProductsBoxes) Validate() error {
	return dara.Validate(s)
}
