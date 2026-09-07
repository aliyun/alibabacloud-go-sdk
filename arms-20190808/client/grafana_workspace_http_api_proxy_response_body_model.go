// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceHttpApiProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GrafanaWorkspaceHttpApiProxyResponseBody
	GetCode() *int32
	SetData(v interface{}) *GrafanaWorkspaceHttpApiProxyResponseBody
	GetData() interface{}
	SetMessage(v string) *GrafanaWorkspaceHttpApiProxyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrafanaWorkspaceHttpApiProxyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GrafanaWorkspaceHttpApiProxyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GrafanaWorkspaceHttpApiProxyResponseBody
	GetTraceId() *string
}

type GrafanaWorkspaceHttpApiProxyResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// eac0a8048716731735000007137d000b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GrafanaWorkspaceHttpApiProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceHttpApiProxyResponseBody) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) SetCode(v int32) *GrafanaWorkspaceHttpApiProxyResponseBody {
	s.Code = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) SetData(v interface{}) *GrafanaWorkspaceHttpApiProxyResponseBody {
	s.Data = v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) SetMessage(v string) *GrafanaWorkspaceHttpApiProxyResponseBody {
	s.Message = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) SetRequestId(v string) *GrafanaWorkspaceHttpApiProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) SetSuccess(v bool) *GrafanaWorkspaceHttpApiProxyResponseBody {
	s.Success = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) SetTraceId(v string) *GrafanaWorkspaceHttpApiProxyResponseBody {
	s.TraceId = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
