// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshRenderingInstanceStreamingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *RefreshRenderingInstanceStreamingResponseBody
	GetFlowId() *string
	SetGateway(v string) *RefreshRenderingInstanceStreamingResponseBody
	GetGateway() *string
	SetHostname(v string) *RefreshRenderingInstanceStreamingResponseBody
	GetHostname() *string
	SetPort(v string) *RefreshRenderingInstanceStreamingResponseBody
	GetPort() *string
	SetRenderingInstanceId(v string) *RefreshRenderingInstanceStreamingResponseBody
	GetRenderingInstanceId() *string
	SetRequestId(v string) *RefreshRenderingInstanceStreamingResponseBody
	GetRequestId() *string
}

type RefreshRenderingInstanceStreamingResponseBody struct {
	// example:
	//
	// 792fy125-594c-4dde-ab35-9ff8hrf0a86f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// 10.178.208.22
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// example:
	//
	// cn-xxx.ecr.aliyuncs.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshRenderingInstanceStreamingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshRenderingInstanceStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshRenderingInstanceStreamingResponseBody) GetFlowId() *string {
	return s.FlowId
}

func (s *RefreshRenderingInstanceStreamingResponseBody) GetGateway() *string {
	return s.Gateway
}

func (s *RefreshRenderingInstanceStreamingResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *RefreshRenderingInstanceStreamingResponseBody) GetPort() *string {
	return s.Port
}

func (s *RefreshRenderingInstanceStreamingResponseBody) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RefreshRenderingInstanceStreamingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshRenderingInstanceStreamingResponseBody) SetFlowId(v string) *RefreshRenderingInstanceStreamingResponseBody {
	s.FlowId = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponseBody) SetGateway(v string) *RefreshRenderingInstanceStreamingResponseBody {
	s.Gateway = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponseBody) SetHostname(v string) *RefreshRenderingInstanceStreamingResponseBody {
	s.Hostname = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponseBody) SetPort(v string) *RefreshRenderingInstanceStreamingResponseBody {
	s.Port = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponseBody) SetRenderingInstanceId(v string) *RefreshRenderingInstanceStreamingResponseBody {
	s.RenderingInstanceId = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponseBody) SetRequestId(v string) *RefreshRenderingInstanceStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponseBody) Validate() error {
	return dara.Validate(s)
}
