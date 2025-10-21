// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchContentSyncDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *BatchContentSyncDetectRequest
	GetRegionId() *string
	SetSceneName(v string) *BatchContentSyncDetectRequest
	GetSceneName() *string
	SetServiceName(v string) *BatchContentSyncDetectRequest
	GetServiceName() *string
	SetServiceParameterList(v []*BatchContentSyncDetectRequestServiceParameterList) *BatchContentSyncDetectRequest
	GetServiceParameterList() []*BatchContentSyncDetectRequestServiceParameterList
}

type BatchContentSyncDetectRequest struct {
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
	//
	// imageDetection
	ServiceName          *string                                              `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameterList []*BatchContentSyncDetectRequestServiceParameterList `json:"serviceParameterList,omitempty" xml:"serviceParameterList,omitempty" type:"Repeated"`
}

func (s BatchContentSyncDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchContentSyncDetectRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *BatchContentSyncDetectRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *BatchContentSyncDetectRequest) GetServiceParameterList() []*BatchContentSyncDetectRequestServiceParameterList {
	return s.ServiceParameterList
}

func (s *BatchContentSyncDetectRequest) SetRegionId(v string) *BatchContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *BatchContentSyncDetectRequest) SetSceneName(v string) *BatchContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *BatchContentSyncDetectRequest) SetServiceName(v string) *BatchContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *BatchContentSyncDetectRequest) SetServiceParameterList(v []*BatchContentSyncDetectRequestServiceParameterList) *BatchContentSyncDetectRequest {
	s.ServiceParameterList = v
	return s
}

func (s *BatchContentSyncDetectRequest) Validate() error {
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

type BatchContentSyncDetectRequestServiceParameterList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s BatchContentSyncDetectRequestServiceParameterList) String() string {
	return dara.Prettify(s)
}

func (s BatchContentSyncDetectRequestServiceParameterList) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectRequestServiceParameterList) GetContent() *string {
	return s.Content
}

func (s *BatchContentSyncDetectRequestServiceParameterList) SetContent(v string) *BatchContentSyncDetectRequestServiceParameterList {
	s.Content = &v
	return s
}

func (s *BatchContentSyncDetectRequestServiceParameterList) Validate() error {
	return dara.Validate(s)
}
