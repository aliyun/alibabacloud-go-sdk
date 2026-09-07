// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGrafanaWorkspaceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateGrafanaWorkspaceAccountResponseBody
	GetCode() *int32
	SetData(v map[string]*string) *CreateGrafanaWorkspaceAccountResponseBody
	GetData() map[string]*string
	SetMessage(v string) *CreateGrafanaWorkspaceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGrafanaWorkspaceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGrafanaWorkspaceAccountResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateGrafanaWorkspaceAccountResponseBody
	GetTraceId() *string
}

type CreateGrafanaWorkspaceAccountResponseBody struct {
	// example:
	//
	// 200
	Code *int32             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 528079E6-B69C-5165-B6AF-DA9FC1E96688
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

func (s CreateGrafanaWorkspaceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) GetData() map[string]*string {
	return s.Data
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) SetCode(v int32) *CreateGrafanaWorkspaceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) SetData(v map[string]*string) *CreateGrafanaWorkspaceAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) SetMessage(v string) *CreateGrafanaWorkspaceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) SetRequestId(v string) *CreateGrafanaWorkspaceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) SetSuccess(v bool) *CreateGrafanaWorkspaceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) SetTraceId(v string) *CreateGrafanaWorkspaceAccountResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
