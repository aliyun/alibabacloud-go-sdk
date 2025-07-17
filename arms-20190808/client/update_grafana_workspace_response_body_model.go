// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGrafanaWorkspaceResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateGrafanaWorkspaceResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateGrafanaWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGrafanaWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGrafanaWorkspaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateGrafanaWorkspaceResponseBody
	GetTraceId() *string
}

type UpdateGrafanaWorkspaceResponseBody struct {
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
	// Indicates whether the update is successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned for the request.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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

func (s UpdateGrafanaWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGrafanaWorkspaceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGrafanaWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGrafanaWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGrafanaWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGrafanaWorkspaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateGrafanaWorkspaceResponseBody) SetCode(v int32) *UpdateGrafanaWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGrafanaWorkspaceResponseBody) SetData(v bool) *UpdateGrafanaWorkspaceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGrafanaWorkspaceResponseBody) SetMessage(v string) *UpdateGrafanaWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGrafanaWorkspaceResponseBody) SetRequestId(v string) *UpdateGrafanaWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceResponseBody) SetSuccess(v bool) *UpdateGrafanaWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGrafanaWorkspaceResponseBody) SetTraceId(v string) *UpdateGrafanaWorkspaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
