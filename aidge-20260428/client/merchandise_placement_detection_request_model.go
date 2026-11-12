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
	// Specifies the ID of a custom API. Pass this parameter if you need to specify a custom version of the API. If you created a custom API during the trial phase, you can query the ApiId of the corresponding feature in the product console by choosing Intelligent Inspection > API Management > My APIs.
	//
	// example:
	//
	// api_xxx
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The URL of the original shelf or floor display image to be recognized. The URL must be accessible over the Internet or from OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/shelf.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ID of the customer-specific SKU vector database. This parameter determines which database is used for retrieval. The database must be created through the database creation process in advance.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The custom rule. You can enter a detection prompt, which is passed as the rule input parameter to the workflow. If this parameter is specified, the type parameter is not required, and the independent rule branch is used. If this parameter is left empty, you must specify the type parameter to start detection.
	//
	// example:
	//
	// Recognize all visible beverage products in the image and return only a JSON array. Output format example: [{"bbox_2d":[100,200,250,600],"sku_name":"Coca-Cola"}]
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The business type. This parameter is reserved for future routing based on business lines. Currently, only skincare products are supported.
	//
	// example:
	//
	// Skincare products
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
