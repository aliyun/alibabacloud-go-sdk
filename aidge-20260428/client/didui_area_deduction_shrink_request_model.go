// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiduiAreaDeductionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductsShrink(v string) *DiduiAreaDeductionShrinkRequest
	GetProductsShrink() *string
	SetRagId(v string) *DiduiAreaDeductionShrinkRequest
	GetRagId() *string
	SetReqId(v string) *DiduiAreaDeductionShrinkRequest
	GetReqId() *string
	SetTargetImageUrl(v string) *DiduiAreaDeductionShrinkRequest
	GetTargetImageUrl() *string
}

type DiduiAreaDeductionShrinkRequest struct {
	// The list of products and their detection boxes.
	//
	// This parameter is required.
	ProductsShrink *string `json:"Products,omitempty" xml:"Products,omitempty"`
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

func (s DiduiAreaDeductionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionShrinkRequest) GetProductsShrink() *string {
	return s.ProductsShrink
}

func (s *DiduiAreaDeductionShrinkRequest) GetRagId() *string {
	return s.RagId
}

func (s *DiduiAreaDeductionShrinkRequest) GetReqId() *string {
	return s.ReqId
}

func (s *DiduiAreaDeductionShrinkRequest) GetTargetImageUrl() *string {
	return s.TargetImageUrl
}

func (s *DiduiAreaDeductionShrinkRequest) SetProductsShrink(v string) *DiduiAreaDeductionShrinkRequest {
	s.ProductsShrink = &v
	return s
}

func (s *DiduiAreaDeductionShrinkRequest) SetRagId(v string) *DiduiAreaDeductionShrinkRequest {
	s.RagId = &v
	return s
}

func (s *DiduiAreaDeductionShrinkRequest) SetReqId(v string) *DiduiAreaDeductionShrinkRequest {
	s.ReqId = &v
	return s
}

func (s *DiduiAreaDeductionShrinkRequest) SetTargetImageUrl(v string) *DiduiAreaDeductionShrinkRequest {
	s.TargetImageUrl = &v
	return s
}

func (s *DiduiAreaDeductionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
