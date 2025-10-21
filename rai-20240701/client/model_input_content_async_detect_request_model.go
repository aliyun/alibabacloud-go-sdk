// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentAsyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *ModelInputContentAsyncDetectRequestBodyData) *ModelInputContentAsyncDetectRequest
	GetBodyData() *ModelInputContentAsyncDetectRequestBodyData
	SetPolicyIdentifier(v string) *ModelInputContentAsyncDetectRequest
	GetPolicyIdentifier() *string
	SetRegionId(v string) *ModelInputContentAsyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *ModelInputContentAsyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *ModelInputContentAsyncDetectRequest
	GetServiceName() *string
}

type ModelInputContentAsyncDetectRequest struct {
	// Request object
	BodyData *ModelInputContentAsyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
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

func (s ModelInputContentAsyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectRequest) GetBodyData() *ModelInputContentAsyncDetectRequestBodyData {
	return s.BodyData
}

func (s *ModelInputContentAsyncDetectRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ModelInputContentAsyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModelInputContentAsyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ModelInputContentAsyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelInputContentAsyncDetectRequest) SetBodyData(v *ModelInputContentAsyncDetectRequestBodyData) *ModelInputContentAsyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetPolicyIdentifier(v string) *ModelInputContentAsyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetRegionId(v string) *ModelInputContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetSceneName(v string) *ModelInputContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetServiceName(v string) *ModelInputContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelInputContentAsyncDetectRequestBodyData struct {
	// 1. The text content to be reviewed, with a maximum limit of 10000 characters (including English and Chinese).
	//
	// 2. Or the URL address of the image to be reviewed.
	//
	// example:
	//
	// the content to be checked
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelInputContentAsyncDetectRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentAsyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectRequestBodyData) GetContent() *string {
	return s.Content
}

func (s *ModelInputContentAsyncDetectRequestBodyData) SetContent(v string) *ModelInputContentAsyncDetectRequestBodyData {
	s.Content = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequestBodyData) Validate() error {
	return dara.Validate(s)
}
