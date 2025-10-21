// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentAsyncDetectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *ModelInputContentAsyncDetectShrinkRequest
	GetBodyDataShrink() *string
	SetPolicyIdentifier(v string) *ModelInputContentAsyncDetectShrinkRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelInputContentAsyncDetectShrinkRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelInputContentAsyncDetectShrinkRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelInputContentAsyncDetectShrinkRequest
	GetServiceName() *string
}

type ModelInputContentAsyncDetectShrinkRequest struct {
	// Request object
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	// Policy ID
	//
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Scene name.
	//
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Service name
	//
	// example:
	//
	// textDetection
	//
	// imageDetection
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelInputContentAsyncDetectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentAsyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *ModelInputContentAsyncDetectShrinkRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelInputContentAsyncDetectShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelInputContentAsyncDetectShrinkRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelInputContentAsyncDetectShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetRegionId(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetSceneName(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetServiceName(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
