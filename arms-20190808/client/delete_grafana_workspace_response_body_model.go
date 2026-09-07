// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGrafanaWorkspaceResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteGrafanaWorkspaceResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteGrafanaWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGrafanaWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGrafanaWorkspaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteGrafanaWorkspaceResponseBody
	GetTraceId() *string
}

type DeleteGrafanaWorkspaceResponseBody struct {
	// The status code. A value of 200 indicates success. Other values indicate errors.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the Grafana workspace was deleted. Valid values:
	//
	// - true: The workspace was deleted.
	//
	// - false: The workspace failed to be deleted.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The operation was successful.
	//
	// - `false`: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, which is used to query the details of the call.
	//
	// example:
	//
	// eac0a8048716731735000007137d000b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteGrafanaWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGrafanaWorkspaceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGrafanaWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGrafanaWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGrafanaWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGrafanaWorkspaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteGrafanaWorkspaceResponseBody) SetCode(v int32) *DeleteGrafanaWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGrafanaWorkspaceResponseBody) SetData(v bool) *DeleteGrafanaWorkspaceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGrafanaWorkspaceResponseBody) SetMessage(v string) *DeleteGrafanaWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGrafanaWorkspaceResponseBody) SetRequestId(v string) *DeleteGrafanaWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceResponseBody) SetSuccess(v bool) *DeleteGrafanaWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGrafanaWorkspaceResponseBody) SetTraceId(v string) *DeleteGrafanaWorkspaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
