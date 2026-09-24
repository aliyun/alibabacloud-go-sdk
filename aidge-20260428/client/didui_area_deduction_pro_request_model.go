// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiduiAreaDeductionProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *DiduiAreaDeductionProRequest
	GetImageUrl() *string
	SetRagId(v string) *DiduiAreaDeductionProRequest
	GetRagId() *string
}

type DiduiAreaDeductionProRequest struct {
	// The HTTP(S) URL of the overall floor display image.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ID of the SKU asset knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// rag_9c1f2b7d4e5a4c8fa1b2c3d4e5f60718
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
}

func (s DiduiAreaDeductionProRequest) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionProRequest) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionProRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DiduiAreaDeductionProRequest) GetRagId() *string {
	return s.RagId
}

func (s *DiduiAreaDeductionProRequest) SetImageUrl(v string) *DiduiAreaDeductionProRequest {
	s.ImageUrl = &v
	return s
}

func (s *DiduiAreaDeductionProRequest) SetRagId(v string) *DiduiAreaDeductionProRequest {
	s.RagId = &v
	return s
}

func (s *DiduiAreaDeductionProRequest) Validate() error {
	return dara.Validate(s)
}
