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
	SetRule(v string) *MerchandisePlacementDetectionRequest
	GetRule() *string
	SetType(v string) *MerchandisePlacementDetectionRequest
	GetType() *string
}

type MerchandisePlacementDetectionRequest struct {
	// Specify this parameter to use a custom API version. If you created a custom API during the trial phase, you can find the corresponding ApiId in the product console under Intelligent Inspection > API Management > My API.
	//
	// example:
	//
	// api_xxx
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The URL of the original shelf or floor stack image to be recognized (accessible over the Internet or through OSS).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/shelf.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ID of the customer-specific SKU vector library, which determines which library is used for retrieval. The library must be created in advance through the library creation process.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The custom rule. Enter a detection prompt as the workflow input parameter rule. When this parameter is specified, the type parameter is not required (a dedicated rule branch is used). If Rule is empty, you must specify Type to start detection.
	//
	// example:
	//
	// Please identify all visible beverage products in the image and return only a JSON array. Output format example: [{"bbox_2d":[100,200,250,600],"sku_name":"Coca-Cola"}]
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The business type (reserved for future routing by business line). The current release supports skincare & lotion.
	//
	// example:
	//
	// 水乳
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

func (s *MerchandisePlacementDetectionRequest) GetRule() *string {
	return s.Rule
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

func (s *MerchandisePlacementDetectionRequest) SetRule(v string) *MerchandisePlacementDetectionRequest {
	s.Rule = &v
	return s
}

func (s *MerchandisePlacementDetectionRequest) SetType(v string) *MerchandisePlacementDetectionRequest {
	s.Type = &v
	return s
}

func (s *MerchandisePlacementDetectionRequest) Validate() error {
	return dara.Validate(s)
}
