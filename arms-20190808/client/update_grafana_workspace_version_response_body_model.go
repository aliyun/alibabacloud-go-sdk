// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGrafanaWorkspaceVersionResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateGrafanaWorkspaceVersionResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateGrafanaWorkspaceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGrafanaWorkspaceVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGrafanaWorkspaceVersionResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateGrafanaWorkspaceVersionResponseBody
	GetTraceId() *string
}

type UpdateGrafanaWorkspaceVersionResponseBody struct {
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A0CEDF1-06FE-44AC-8E21-21A5BE65****
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

func (s UpdateGrafanaWorkspaceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) SetCode(v int32) *UpdateGrafanaWorkspaceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) SetData(v bool) *UpdateGrafanaWorkspaceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) SetMessage(v string) *UpdateGrafanaWorkspaceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) SetRequestId(v string) *UpdateGrafanaWorkspaceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) SetSuccess(v bool) *UpdateGrafanaWorkspaceVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) SetTraceId(v string) *UpdateGrafanaWorkspaceVersionResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
