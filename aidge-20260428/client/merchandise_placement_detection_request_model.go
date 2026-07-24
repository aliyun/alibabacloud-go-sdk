// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *MerchandisePlacementDetectionRequest
	GetApiId() *string
	SetImageUrl(v string) *MerchandisePlacementDetectionRequest
	GetImageUrl() *string
	SetRagId(v string) *MerchandisePlacementDetectionRequest
	GetRagId() *string
	SetType(v string) *MerchandisePlacementDetectionRequest
	GetType() *string
}

type MerchandisePlacementDetectionRequest struct {
	// Specifies a custom API version. If you created a "My API" during the trial phase, you can find the corresponding ApiId in the product console under "Intelligent Inspection > API Management > My API".
	//
	// example:
	//
	// api_xxx
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The URL of the shelf or floor-stack image to be recognized (accessible via the public network or OSS).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/shelf.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The customer-specific SKU vector library ID that determines which library to retrieve from. The library must be created in advance through the library creation process.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The business type (reserved for future routing by business line). The current release supports skincare & lotion.
	//
	// example:
	//
	// Lotion.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MerchandisePlacementDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionRequest) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionRequest) GetApiId() *string {
	return s.ApiId
}

func (s *MerchandisePlacementDetectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MerchandisePlacementDetectionRequest) GetRagId() *string {
	return s.RagId
}

func (s *MerchandisePlacementDetectionRequest) GetType() *string {
	return s.Type
}

func (s *MerchandisePlacementDetectionRequest) SetApiId(v string) *MerchandisePlacementDetectionRequest {
	s.ApiId = &v
	return s
}

func (s *MerchandisePlacementDetectionRequest) SetImageUrl(v string) *MerchandisePlacementDetectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *MerchandisePlacementDetectionRequest) SetRagId(v string) *MerchandisePlacementDetectionRequest {
	s.RagId = &v
	return s
}

func (s *MerchandisePlacementDetectionRequest) SetType(v string) *MerchandisePlacementDetectionRequest {
	s.Type = &v
	return s
}

func (s *MerchandisePlacementDetectionRequest) Validate() error {
	return dara.Validate(s)
}
