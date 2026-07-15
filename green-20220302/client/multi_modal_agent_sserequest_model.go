// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalAgentSSERequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppID(v string) *MultiModalAgentSSERequest
	GetAppID() *string
	SetServiceParameters(v string) *MultiModalAgentSSERequest
	GetServiceParameters() *string
	SetStream(v string) *MultiModalAgentSSERequest
	GetStream() *string
}

type MultiModalAgentSSERequest struct {
	// The unique identifier of the whiteboard application. To obtain the whiteboard application ID, see [CreateApp](https://help.aliyun.com/document_detail/204234.html).
	//
	// example:
	//
	// txt_check_pro_agent_01
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// The parameter set required by the moderation service, in JSON string format. The input parameter for text content is content (String), the custom data ID is DataId (String), and the cache type is CacheType (String, valid value: ephemeral).
	//
	// example:
	//
	// {
	//
	//     "content": "这里待审核的文本内容",
	//
	//     "DataId": "data123***",
	//
	//     "CacheType":"ephemeral"
	//
	//     }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	// Specifies whether to use streaming output.
	//
	// example:
	//
	// true
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s MultiModalAgentSSERequest) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentSSERequest) GoString() string {
	return s.String()
}

func (s *MultiModalAgentSSERequest) GetAppID() *string {
	return s.AppID
}

func (s *MultiModalAgentSSERequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalAgentSSERequest) GetStream() *string {
	return s.Stream
}

func (s *MultiModalAgentSSERequest) SetAppID(v string) *MultiModalAgentSSERequest {
	s.AppID = &v
	return s
}

func (s *MultiModalAgentSSERequest) SetServiceParameters(v string) *MultiModalAgentSSERequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalAgentSSERequest) SetStream(v string) *MultiModalAgentSSERequest {
	s.Stream = &v
	return s
}

func (s *MultiModalAgentSSERequest) Validate() error {
	return dara.Validate(s)
}
