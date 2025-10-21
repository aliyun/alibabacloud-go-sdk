// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentSyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *ModelInputContentSyncDetectRequestBodyData) *ModelInputContentSyncDetectRequest
	GetBodyData() *ModelInputContentSyncDetectRequestBodyData
	SetPolicyIdentifier(v string) *ModelInputContentSyncDetectRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelInputContentSyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelInputContentSyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelInputContentSyncDetectRequest
	GetServiceName() *string
}

type ModelInputContentSyncDetectRequest struct {
	// Request object
	BodyData *ModelInputContentSyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
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

func (s ModelInputContentSyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectRequest) GetBodyData() *ModelInputContentSyncDetectRequestBodyData {
	return s.BodyData
}

func (s *ModelInputContentSyncDetectRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelInputContentSyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelInputContentSyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelInputContentSyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelInputContentSyncDetectRequest) SetBodyData(v *ModelInputContentSyncDetectRequestBodyData) *ModelInputContentSyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetPolicyIdentifier(v string) *ModelInputContentSyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetRegionId(v string) *ModelInputContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetSceneName(v string) *ModelInputContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetServiceName(v string) *ModelInputContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelInputContentSyncDetectRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelInputContentSyncDetectRequestBodyData struct {
	// 1. The text content to be checked, with a maximum limit of 10,000 characters (including English and Chinese).
	//
	// 2. Or the URL address of the image to be checked.
	//
	// example:
	//
	// 要检测的内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelInputContentSyncDetectRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectRequestBodyData) GetContent() *string {
	return s.Content
}

func (s *ModelInputContentSyncDetectRequestBodyData) SetContent(v string) *ModelInputContentSyncDetectRequestBodyData {
	s.Content = &v
	return s
}

func (s *ModelInputContentSyncDetectRequestBodyData) Validate() error {
	return dara.Validate(s)
}
