// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchContentAsyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *BatchContentAsyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *BatchContentAsyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *BatchContentAsyncDetectRequest
	GetServiceName() *string
	SetServiceParameterList(v []*BatchContentAsyncDetectRequestServiceParameterList) *BatchContentAsyncDetectRequest
	GetServiceParameterList() []*BatchContentAsyncDetectRequestServiceParameterList
}

type BatchContentAsyncDetectRequest struct {
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
	ServiceName          *string                                               `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameterList []*BatchContentAsyncDetectRequestServiceParameterList `json:"serviceParameterList,omitempty" xml:"serviceParameterList,omitempty" type:"Repeated"`
}

func (s BatchContentAsyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchContentAsyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *BatchContentAsyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *BatchContentAsyncDetectRequest) GetServiceParameterList() []*BatchContentAsyncDetectRequestServiceParameterList {
	return s.ServiceParameterList
}

func (s *BatchContentAsyncDetectRequest) SetRegionId(v string) *BatchContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *BatchContentAsyncDetectRequest) SetSceneName(v string) *BatchContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *BatchContentAsyncDetectRequest) SetServiceName(v string) *BatchContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *BatchContentAsyncDetectRequest) SetServiceParameterList(v []*BatchContentAsyncDetectRequestServiceParameterList) *BatchContentAsyncDetectRequest {
	s.ServiceParameterList = v
	return s
}

func (s *BatchContentAsyncDetectRequest) Validate() error {
	if s.ServiceParameterList != nil {
		for _, item := range s.ServiceParameterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchContentAsyncDetectRequestServiceParameterList struct {
	// example:
	//
	// "XXXXXXXX"
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s BatchContentAsyncDetectRequestServiceParameterList) String() string {
	return dara.Prettify(s)
}

func (s BatchContentAsyncDetectRequestServiceParameterList) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectRequestServiceParameterList) GetContent() *string {
	return s.Content
}

func (s *BatchContentAsyncDetectRequestServiceParameterList) SetContent(v string) *BatchContentAsyncDetectRequestServiceParameterList {
	s.Content = &v
	return s
}

func (s *BatchContentAsyncDetectRequestServiceParameterList) Validate() error {
	return dara.Validate(s)
}
