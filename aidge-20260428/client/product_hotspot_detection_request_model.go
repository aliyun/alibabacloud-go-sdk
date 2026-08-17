// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductHotspotDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReferenceImageUrls(v []*string) *ProductHotspotDetectionRequest
	GetReferenceImageUrls() []*string
	SetReqId(v string) *ProductHotspotDetectionRequest
	GetReqId() *string
	SetTargetImageUrl(v string) *ProductHotspotDetectionRequest
	GetTargetImageUrl() *string
}

type ProductHotspotDetectionRequest struct {
	// The HTTPS URLs of reference images that define the SKU whitelist. A maximum of 20 images are supported.
	//
	// This parameter is required.
	ReferenceImageUrls []*string `json:"ReferenceImageUrls,omitempty" xml:"ReferenceImageUrls,omitempty" type:"Repeated"`
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

func (s ProductHotspotDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ProductHotspotDetectionRequest) GoString() string {
	return s.String()
}

func (s *ProductHotspotDetectionRequest) GetReferenceImageUrls() []*string {
	return s.ReferenceImageUrls
}

func (s *ProductHotspotDetectionRequest) GetReqId() *string {
	return s.ReqId
}

func (s *ProductHotspotDetectionRequest) GetTargetImageUrl() *string {
	return s.TargetImageUrl
}

func (s *ProductHotspotDetectionRequest) SetReferenceImageUrls(v []*string) *ProductHotspotDetectionRequest {
	s.ReferenceImageUrls = v
	return s
}

func (s *ProductHotspotDetectionRequest) SetReqId(v string) *ProductHotspotDetectionRequest {
	s.ReqId = &v
	return s
}

func (s *ProductHotspotDetectionRequest) SetTargetImageUrl(v string) *ProductHotspotDetectionRequest {
	s.TargetImageUrl = &v
	return s
}

func (s *ProductHotspotDetectionRequest) Validate() error {
	return dara.Validate(s)
}
