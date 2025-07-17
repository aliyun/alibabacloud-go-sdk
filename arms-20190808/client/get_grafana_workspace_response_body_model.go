// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrafanaWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGrafanaWorkspaceResponseBody
	GetCode() *int32
	SetData(v *GrafanaWorkspace) *GetGrafanaWorkspaceResponseBody
	GetData() *GrafanaWorkspace
	SetMessage(v string) *GetGrafanaWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGrafanaWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGrafanaWorkspaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GetGrafanaWorkspaceResponseBody
	GetTraceId() *string
}

type GetGrafanaWorkspaceResponseBody struct {
	// The HTTP status code returned for the request. Valid values:
	//
	// 	- `2XX`: The request is successful.
	//
	// 	- `3XX`: A redirection message is returned.
	//
	// 	- `4XX`: The request is invalid.
	//
	// 	- `5XX`: A server error occurs.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the Grafana workspace.
	Data *GrafanaWorkspace `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when the request parameters are invalid.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2C3F217B-9AAE-5D51-974D-48********
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
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// eac0a8048716731735000007137d000b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetGrafanaWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGrafanaWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetGrafanaWorkspaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGrafanaWorkspaceResponseBody) GetData() *GrafanaWorkspace {
	return s.Data
}

func (s *GetGrafanaWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGrafanaWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGrafanaWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGrafanaWorkspaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetGrafanaWorkspaceResponseBody) SetCode(v int32) *GetGrafanaWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetGrafanaWorkspaceResponseBody) SetData(v *GrafanaWorkspace) *GetGrafanaWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *GetGrafanaWorkspaceResponseBody) SetMessage(v string) *GetGrafanaWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetGrafanaWorkspaceResponseBody) SetRequestId(v string) *GetGrafanaWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGrafanaWorkspaceResponseBody) SetSuccess(v bool) *GetGrafanaWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetGrafanaWorkspaceResponseBody) SetTraceId(v string) *GetGrafanaWorkspaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetGrafanaWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
