// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGrafanaWorkspaceAccountResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteGrafanaWorkspaceAccountResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteGrafanaWorkspaceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGrafanaWorkspaceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGrafanaWorkspaceAccountResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteGrafanaWorkspaceAccountResponseBody
	GetTraceId() *string
}

type DeleteGrafanaWorkspaceAccountResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// []
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 92FEE957-B3CF-56E7-8138-45945A9B471F
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

func (s DeleteGrafanaWorkspaceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) SetCode(v int32) *DeleteGrafanaWorkspaceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) SetData(v bool) *DeleteGrafanaWorkspaceAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) SetMessage(v string) *DeleteGrafanaWorkspaceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) SetRequestId(v string) *DeleteGrafanaWorkspaceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) SetSuccess(v bool) *DeleteGrafanaWorkspaceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) SetTraceId(v string) *DeleteGrafanaWorkspaceAccountResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
