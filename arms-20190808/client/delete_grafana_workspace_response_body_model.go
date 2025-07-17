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
	// The status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the workspace was deleted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
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
