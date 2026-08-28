// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPluginPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunPluginPipelineResponseBody
	GetCode() *string
	SetMessage(v string) *RunPluginPipelineResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunPluginPipelineResponseBody
	GetRequestId() *string
}

type RunPluginPipelineResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E9E6192C-E3D6-5176-9109-340E9DA7CADD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunPluginPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunPluginPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *RunPluginPipelineResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunPluginPipelineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunPluginPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunPluginPipelineResponseBody) SetCode(v string) *RunPluginPipelineResponseBody {
	s.Code = &v
	return s
}

func (s *RunPluginPipelineResponseBody) SetMessage(v string) *RunPluginPipelineResponseBody {
	s.Message = &v
	return s
}

func (s *RunPluginPipelineResponseBody) SetRequestId(v string) *RunPluginPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPluginPipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
