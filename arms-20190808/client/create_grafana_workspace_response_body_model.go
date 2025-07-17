// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGrafanaWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateGrafanaWorkspaceResponseBody
	GetCode() *int32
	SetData(v *GrafanaWorkspace) *CreateGrafanaWorkspaceResponseBody
	GetData() *GrafanaWorkspace
	SetMessage(v string) *CreateGrafanaWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGrafanaWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGrafanaWorkspaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateGrafanaWorkspaceResponseBody
	GetTraceId() *string
}

type CreateGrafanaWorkspaceResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the Grafana workspace.
	Data *GrafanaWorkspace `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D80ADAAC-8C32-5479-BD14-C28CF832****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// eac0a8048716731735000007137d000b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateGrafanaWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateGrafanaWorkspaceResponseBody) GetData() *GrafanaWorkspace {
	return s.Data
}

func (s *CreateGrafanaWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGrafanaWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGrafanaWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGrafanaWorkspaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateGrafanaWorkspaceResponseBody) SetCode(v int32) *CreateGrafanaWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGrafanaWorkspaceResponseBody) SetData(v *GrafanaWorkspace) *CreateGrafanaWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateGrafanaWorkspaceResponseBody) SetMessage(v string) *CreateGrafanaWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGrafanaWorkspaceResponseBody) SetRequestId(v string) *CreateGrafanaWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGrafanaWorkspaceResponseBody) SetSuccess(v bool) *CreateGrafanaWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGrafanaWorkspaceResponseBody) SetTraceId(v string) *CreateGrafanaWorkspaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateGrafanaWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
