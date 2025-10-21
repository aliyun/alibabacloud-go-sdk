// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentSyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *ModelOutputContentSyncDetectRequestBodyData) *ModelOutputContentSyncDetectRequest
	GetBodyData() *ModelOutputContentSyncDetectRequestBodyData
	SetPolicyIdentifier(v string) *ModelOutputContentSyncDetectRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelOutputContentSyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelOutputContentSyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelOutputContentSyncDetectRequest
	GetServiceName() *string
}

type ModelOutputContentSyncDetectRequest struct {
	// Request object
	BodyData *ModelOutputContentSyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
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

func (s ModelOutputContentSyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectRequest) GetBodyData() *ModelOutputContentSyncDetectRequestBodyData {
	return s.BodyData
}

func (s *ModelOutputContentSyncDetectRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelOutputContentSyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelOutputContentSyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelOutputContentSyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelOutputContentSyncDetectRequest) SetBodyData(v *ModelOutputContentSyncDetectRequestBodyData) *ModelOutputContentSyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetPolicyIdentifier(v string) *ModelOutputContentSyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetRegionId(v string) *ModelOutputContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetSceneName(v string) *ModelOutputContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetServiceName(v string) *ModelOutputContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelOutputContentSyncDetectRequestBodyData struct {
	// 1. The text content to be reviewed, with a maximum limit of 10,000 characters (including English and Chinese).
	//
	// 2. Or the URL address of the image to be reviewed.
	//
	// example:
	//
	// the content to be checked
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelOutputContentSyncDetectRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectRequestBodyData) GetContent() *string {
	return s.Content
}

func (s *ModelOutputContentSyncDetectRequestBodyData) SetContent(v string) *ModelOutputContentSyncDetectRequestBodyData {
	s.Content = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequestBodyData) Validate() error {
	return dara.Validate(s)
}
