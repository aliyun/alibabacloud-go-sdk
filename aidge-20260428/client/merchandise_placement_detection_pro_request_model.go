// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *MerchandisePlacementDetectionProRequest
	GetImageUrl() *string
	SetRule(v string) *MerchandisePlacementDetectionProRequest
	GetRule() *string
	SetType(v string) *MerchandisePlacementDetectionProRequest
	GetType() *string
}

type MerchandisePlacementDetectionProRequest struct {
	// The HTTPS URL of the display image to detect.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The detection rule. When non-empty, this value takes priority as the model prompt.
	//
	// example:
	//
	// Identify all Genki Forest beverages on the shelf and mark their positions
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The product type. This parameter must be set to Genki Forest when Rule is empty.
	//
	// example:
	//
	// 元气森林
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MerchandisePlacementDetectionProRequest) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionProRequest) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionProRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MerchandisePlacementDetectionProRequest) GetRule() *string {
	return s.Rule
}

func (s *MerchandisePlacementDetectionProRequest) GetType() *string {
	return s.Type
}

func (s *MerchandisePlacementDetectionProRequest) SetImageUrl(v string) *MerchandisePlacementDetectionProRequest {
	s.ImageUrl = &v
	return s
}

func (s *MerchandisePlacementDetectionProRequest) SetRule(v string) *MerchandisePlacementDetectionProRequest {
	s.Rule = &v
	return s
}

func (s *MerchandisePlacementDetectionProRequest) SetType(v string) *MerchandisePlacementDetectionProRequest {
	s.Type = &v
	return s
}

func (s *MerchandisePlacementDetectionProRequest) Validate() error {
	return dara.Validate(s)
}
