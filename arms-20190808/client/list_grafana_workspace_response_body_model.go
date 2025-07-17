// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrafanaWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGrafanaWorkspaceResponseBody
	GetCode() *int32
	SetData(v []*GrafanaWorkspace) *ListGrafanaWorkspaceResponseBody
	GetData() []*GrafanaWorkspace
	SetMessage(v string) *ListGrafanaWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGrafanaWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGrafanaWorkspaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListGrafanaWorkspaceResponseBody
	GetTraceId() *string
}

type ListGrafanaWorkspaceResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data []*GrafanaWorkspace `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message returned if the request parameters are invalid.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0080BE65-167F-5CB6-A691-14E2EFD474BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// eac0a8048716731735000007137d000b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListGrafanaWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGrafanaWorkspaceResponseBody) GetData() []*GrafanaWorkspace {
	return s.Data
}

func (s *ListGrafanaWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGrafanaWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGrafanaWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGrafanaWorkspaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListGrafanaWorkspaceResponseBody) SetCode(v int32) *ListGrafanaWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *ListGrafanaWorkspaceResponseBody) SetData(v []*GrafanaWorkspace) *ListGrafanaWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *ListGrafanaWorkspaceResponseBody) SetMessage(v string) *ListGrafanaWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *ListGrafanaWorkspaceResponseBody) SetRequestId(v string) *ListGrafanaWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGrafanaWorkspaceResponseBody) SetSuccess(v bool) *ListGrafanaWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *ListGrafanaWorkspaceResponseBody) SetTraceId(v string) *ListGrafanaWorkspaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListGrafanaWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
