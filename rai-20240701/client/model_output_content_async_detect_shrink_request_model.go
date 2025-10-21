// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentAsyncDetectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *ModelOutputContentAsyncDetectShrinkRequest
	GetBodyDataShrink() *string
	SetPolicyIdentifier(v string) *ModelOutputContentAsyncDetectShrinkRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelOutputContentAsyncDetectShrinkRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelOutputContentAsyncDetectShrinkRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelOutputContentAsyncDetectShrinkRequest
	GetServiceName() *string
}

type ModelOutputContentAsyncDetectShrinkRequest struct {
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

func (s ModelOutputContentAsyncDetectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentAsyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetRegionId(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetSceneName(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetServiceName(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
