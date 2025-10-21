// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentSyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ContentSyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *ContentSyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *ContentSyncDetectRequest
	GetServiceName() *string
	SetServiceParameter(v *ContentSyncDetectRequestServiceParameter) *ContentSyncDetectRequest
	GetServiceParameter() *ContentSyncDetectRequestServiceParameter
}

type ContentSyncDetectRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName      *string                                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameter *ContentSyncDetectRequestServiceParameter `json:"serviceParameter,omitempty" xml:"serviceParameter,omitempty" type:"Struct"`
}

func (s ContentSyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s ContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ContentSyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ContentSyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ContentSyncDetectRequest) GetServiceParameter() *ContentSyncDetectRequestServiceParameter {
	return s.ServiceParameter
}

func (s *ContentSyncDetectRequest) SetRegionId(v string) *ContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ContentSyncDetectRequest) SetSceneName(v string) *ContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ContentSyncDetectRequest) SetServiceName(v string) *ContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ContentSyncDetectRequest) SetServiceParameter(v *ContentSyncDetectRequestServiceParameter) *ContentSyncDetectRequest {
	s.ServiceParameter = v
	return s
}

func (s *ContentSyncDetectRequest) Validate() error {
	if s.ServiceParameter != nil {
		if err := s.ServiceParameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ContentSyncDetectRequestServiceParameter struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ContentSyncDetectRequestServiceParameter) String() string {
	return dara.Prettify(s)
}

func (s ContentSyncDetectRequestServiceParameter) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectRequestServiceParameter) GetContent() *string {
	return s.Content
}

func (s *ContentSyncDetectRequestServiceParameter) SetContent(v string) *ContentSyncDetectRequestServiceParameter {
	s.Content = &v
	return s
}

func (s *ContentSyncDetectRequestServiceParameter) Validate() error {
	return dara.Validate(s)
}
