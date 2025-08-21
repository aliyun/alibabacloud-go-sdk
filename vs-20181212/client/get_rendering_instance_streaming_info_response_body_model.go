// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingInstanceStreamingInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *GetRenderingInstanceStreamingInfoResponseBody
	GetFlowId() *string
	SetGateway(v string) *GetRenderingInstanceStreamingInfoResponseBody
	GetGateway() *string
	SetHostname(v string) *GetRenderingInstanceStreamingInfoResponseBody
	GetHostname() *string
	SetPort(v string) *GetRenderingInstanceStreamingInfoResponseBody
	GetPort() *string
	SetRenderingInstanceId(v string) *GetRenderingInstanceStreamingInfoResponseBody
	GetRenderingInstanceId() *string
	SetRequestId(v string) *GetRenderingInstanceStreamingInfoResponseBody
	GetRequestId() *string
}

type GetRenderingInstanceStreamingInfoResponseBody struct {
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
	// 10.18.20.2
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 10003
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRenderingInstanceStreamingInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingInstanceStreamingInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) GetFlowId() *string {
	return s.FlowId
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) GetGateway() *string {
	return s.Gateway
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) GetPort() *string {
	return s.Port
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) SetFlowId(v string) *GetRenderingInstanceStreamingInfoResponseBody {
	s.FlowId = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) SetGateway(v string) *GetRenderingInstanceStreamingInfoResponseBody {
	s.Gateway = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) SetHostname(v string) *GetRenderingInstanceStreamingInfoResponseBody {
	s.Hostname = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) SetPort(v string) *GetRenderingInstanceStreamingInfoResponseBody {
	s.Port = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) SetRenderingInstanceId(v string) *GetRenderingInstanceStreamingInfoResponseBody {
	s.RenderingInstanceId = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) SetRequestId(v string) *GetRenderingInstanceStreamingInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
