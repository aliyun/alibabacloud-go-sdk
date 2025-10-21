// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentAsyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *ModelOutputContentAsyncDetectRequestBodyData) *ModelOutputContentAsyncDetectRequest
	GetBodyData() *ModelOutputContentAsyncDetectRequestBodyData
	SetPolicyIdentifier(v string) *ModelOutputContentAsyncDetectRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelOutputContentAsyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelOutputContentAsyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelOutputContentAsyncDetectRequest
	GetServiceName() *string
}

type ModelOutputContentAsyncDetectRequest struct {
	// Request object
	BodyData *ModelOutputContentAsyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
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

func (s ModelOutputContentAsyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectRequest) GetBodyData() *ModelOutputContentAsyncDetectRequestBodyData {
	return s.BodyData
}

func (s *ModelOutputContentAsyncDetectRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelOutputContentAsyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelOutputContentAsyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelOutputContentAsyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelOutputContentAsyncDetectRequest) SetBodyData(v *ModelOutputContentAsyncDetectRequestBodyData) *ModelOutputContentAsyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetPolicyIdentifier(v string) *ModelOutputContentAsyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetRegionId(v string) *ModelOutputContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetSceneName(v string) *ModelOutputContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetServiceName(v string) *ModelOutputContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelOutputContentAsyncDetectRequestBodyData struct {
	// 1. The text content to be reviewed, with a maximum limit of 10000 characters (including English and Chinese).
	//
	// 2. Or the URL address of the image to be reviewed.
	//
	// example:
	//
	// the content to be checked
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelOutputContentAsyncDetectRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentAsyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectRequestBodyData) GetContent() *string {
	return s.Content
}

func (s *ModelOutputContentAsyncDetectRequestBodyData) SetContent(v string) *ModelOutputContentAsyncDetectRequestBodyData {
	s.Content = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequestBodyData) Validate() error {
	return dara.Validate(s)
}
