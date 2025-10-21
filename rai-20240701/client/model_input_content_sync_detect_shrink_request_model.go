// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentSyncDetectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *ModelInputContentSyncDetectShrinkRequest
	GetBodyDataShrink() *string
	SetPolicyIdentifier(v string) *ModelInputContentSyncDetectShrinkRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelInputContentSyncDetectShrinkRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelInputContentSyncDetectShrinkRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelInputContentSyncDetectShrinkRequest
	GetServiceName() *string
}

type ModelInputContentSyncDetectShrinkRequest struct {
	// Request object
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	// Policy Identifier
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

func (s ModelInputContentSyncDetectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *ModelInputContentSyncDetectShrinkRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelInputContentSyncDetectShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelInputContentSyncDetectShrinkRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelInputContentSyncDetectShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetRegionId(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetSceneName(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetServiceName(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
