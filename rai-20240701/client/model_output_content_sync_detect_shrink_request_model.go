// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentSyncDetectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *ModelOutputContentSyncDetectShrinkRequest
	GetBodyDataShrink() *string
	SetPolicyIdentifier(v string) *ModelOutputContentSyncDetectShrinkRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelOutputContentSyncDetectShrinkRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelOutputContentSyncDetectShrinkRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelOutputContentSyncDetectShrinkRequest
	GetServiceName() *string
}

type ModelOutputContentSyncDetectShrinkRequest struct {
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

func (s ModelOutputContentSyncDetectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *ModelOutputContentSyncDetectShrinkRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelOutputContentSyncDetectShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelOutputContentSyncDetectShrinkRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelOutputContentSyncDetectShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetRegionId(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetSceneName(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetServiceName(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
