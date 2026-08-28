// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaterialInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *MaterialInspectionRequest
	GetApiId() *string
	SetImageRefer(v string) *MaterialInspectionRequest
	GetImageRefer() *string
	SetImageUrl(v string) *MaterialInspectionRequest
	GetImageUrl() *string
	SetReqId(v string) *MaterialInspectionRequest
	GetReqId() *string
	SetRules(v string) *MaterialInspectionRequest
	GetRules() *string
}

type MaterialInspectionRequest struct {
	// apiId
	//
	// example:
	//
	// fb0012f49b004f889207a3c5e6ef8da9
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The URL of the reference image. Optional. When this parameter is not empty, the system performs multi-image differential comparison.
	//
	// example:
	//
	// https://example.com/reference.jpg
	ImageRefer *string `json:"ImageRefer,omitempty" xml:"ImageRefer,omitempty"`
	// The URL of the target image to be inspected.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/store.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The business request identity. Optional. This value is returned as-is in the response for Tracing Analysis purposes.
	//
	// example:
	//
	// req-2026-06-04-001
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The rule prompt in natural language that describes the inspection criteria.
	//
	// example:
	//
	// The door-shaped display stand must be placed in the entrance area; the functional display card must be placed on the counter surface
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s MaterialInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s MaterialInspectionRequest) GoString() string {
	return s.String()
}

func (s *MaterialInspectionRequest) GetApiId() *string {
	return s.ApiId
}

func (s *MaterialInspectionRequest) GetImageRefer() *string {
	return s.ImageRefer
}

func (s *MaterialInspectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MaterialInspectionRequest) GetReqId() *string {
	return s.ReqId
}

func (s *MaterialInspectionRequest) GetRules() *string {
	return s.Rules
}

func (s *MaterialInspectionRequest) SetApiId(v string) *MaterialInspectionRequest {
	s.ApiId = &v
	return s
}

func (s *MaterialInspectionRequest) SetImageRefer(v string) *MaterialInspectionRequest {
	s.ImageRefer = &v
	return s
}

func (s *MaterialInspectionRequest) SetImageUrl(v string) *MaterialInspectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *MaterialInspectionRequest) SetReqId(v string) *MaterialInspectionRequest {
	s.ReqId = &v
	return s
}

func (s *MaterialInspectionRequest) SetRules(v string) *MaterialInspectionRequest {
	s.Rules = &v
	return s
}

func (s *MaterialInspectionRequest) Validate() error {
	return dara.Validate(s)
}
