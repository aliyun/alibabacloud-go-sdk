// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentAsyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ContentAsyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *ContentAsyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *ContentAsyncDetectRequest
	GetServiceName() *string
	SetServiceParameter(v *ContentAsyncDetectRequestServiceParameter) *ContentAsyncDetectRequest
	GetServiceParameter() *ContentAsyncDetectRequestServiceParameter
}

type ContentAsyncDetectRequest struct {
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
	ServiceName      *string                                    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameter *ContentAsyncDetectRequestServiceParameter `json:"serviceParameter,omitempty" xml:"serviceParameter,omitempty" type:"Struct"`
}

func (s ContentAsyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s ContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ContentAsyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ContentAsyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ContentAsyncDetectRequest) GetServiceParameter() *ContentAsyncDetectRequestServiceParameter {
	return s.ServiceParameter
}

func (s *ContentAsyncDetectRequest) SetRegionId(v string) *ContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ContentAsyncDetectRequest) SetSceneName(v string) *ContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ContentAsyncDetectRequest) SetServiceName(v string) *ContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ContentAsyncDetectRequest) SetServiceParameter(v *ContentAsyncDetectRequestServiceParameter) *ContentAsyncDetectRequest {
	s.ServiceParameter = v
	return s
}

func (s *ContentAsyncDetectRequest) Validate() error {
	if s.ServiceParameter != nil {
		if err := s.ServiceParameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ContentAsyncDetectRequestServiceParameter struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ContentAsyncDetectRequestServiceParameter) String() string {
	return dara.Prettify(s)
}

func (s ContentAsyncDetectRequestServiceParameter) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectRequestServiceParameter) GetContent() *string {
	return s.Content
}

func (s *ContentAsyncDetectRequestServiceParameter) SetContent(v string) *ContentAsyncDetectRequestServiceParameter {
	s.Content = &v
	return s
}

func (s *ContentAsyncDetectRequestServiceParameter) Validate() error {
	return dara.Validate(s)
}
