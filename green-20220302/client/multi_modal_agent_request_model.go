// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppID(v string) *MultiModalAgentRequest
	GetAppID() *string
	SetServiceParameters(v string) *MultiModalAgentRequest
	GetServiceParameters() *string
}

type MultiModalAgentRequest struct {
	// The unique identifier of the whiteboard application. To get the whiteboard application ID, see [CreateApp](https://help.aliyun.com/document_detail/204234.html).
	//
	// example:
	//
	// txt_check_agent_01
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// The set of parameters for the auditing service. This includes the taskId of the detection task to query. You can specify only one taskId at a time.
	//
	// example:
	//
	// {"content":"测试文本","dataId":"img1234567"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentRequest) GoString() string {
	return s.String()
}

func (s *MultiModalAgentRequest) GetAppID() *string {
	return s.AppID
}

func (s *MultiModalAgentRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalAgentRequest) SetAppID(v string) *MultiModalAgentRequest {
	s.AppID = &v
	return s
}

func (s *MultiModalAgentRequest) SetServiceParameters(v string) *MultiModalAgentRequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalAgentRequest) Validate() error {
	return dara.Validate(s)
}
