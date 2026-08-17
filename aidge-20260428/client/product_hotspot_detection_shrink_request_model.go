// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductHotspotDetectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReferenceImageUrlsShrink(v string) *ProductHotspotDetectionShrinkRequest
	GetReferenceImageUrlsShrink() *string
	SetReqId(v string) *ProductHotspotDetectionShrinkRequest
	GetReqId() *string
	SetTargetImageUrl(v string) *ProductHotspotDetectionShrinkRequest
	GetTargetImageUrl() *string
}

type ProductHotspotDetectionShrinkRequest struct {
	// The HTTPS URLs of reference images that define the SKU whitelist. A maximum of 20 images are supported.
	//
	// This parameter is required.
	ReferenceImageUrlsShrink *string `json:"ReferenceImageUrls,omitempty" xml:"ReferenceImageUrls,omitempty"`
	// The unique business ID for this single-scene call.
	//
	// example:
	//
	// hotspot-request-001
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The HTTPS OSS or CDN URL of the target image to be annotated with bounding boxes.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/scene.jpg
	TargetImageUrl *string `json:"TargetImageUrl,omitempty" xml:"TargetImageUrl,omitempty"`
}

func (s ProductHotspotDetectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ProductHotspotDetectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ProductHotspotDetectionShrinkRequest) GetReferenceImageUrlsShrink() *string {
	return s.ReferenceImageUrlsShrink
}

func (s *ProductHotspotDetectionShrinkRequest) GetReqId() *string {
	return s.ReqId
}

func (s *ProductHotspotDetectionShrinkRequest) GetTargetImageUrl() *string {
	return s.TargetImageUrl
}

func (s *ProductHotspotDetectionShrinkRequest) SetReferenceImageUrlsShrink(v string) *ProductHotspotDetectionShrinkRequest {
	s.ReferenceImageUrlsShrink = &v
	return s
}

func (s *ProductHotspotDetectionShrinkRequest) SetReqId(v string) *ProductHotspotDetectionShrinkRequest {
	s.ReqId = &v
	return s
}

func (s *ProductHotspotDetectionShrinkRequest) SetTargetImageUrl(v string) *ProductHotspotDetectionShrinkRequest {
	s.TargetImageUrl = &v
	return s
}

func (s *ProductHotspotDetectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
