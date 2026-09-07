// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceAccountRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGrafanaWorkspaceAccountRoleResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateGrafanaWorkspaceAccountRoleResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateGrafanaWorkspaceAccountRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGrafanaWorkspaceAccountRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGrafanaWorkspaceAccountRoleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateGrafanaWorkspaceAccountRoleResponseBody
	GetTraceId() *string
}

type UpdateGrafanaWorkspaceAccountRoleResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//         "RequestId": "01A079AC-E254-5148-8B38-948C909DD31B",
	//
	//         "TraceId": "0a06dd2d17887477684538583ec77f",
	//
	//         "Data": true,
	//
	//         "Code": 200,
	//
	//         "Success": true
	//
	//     }
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5540BA0C-84FF-5D38-B7A9-D78B84C98C18
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

func (s UpdateGrafanaWorkspaceAccountRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceAccountRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) SetCode(v int32) *UpdateGrafanaWorkspaceAccountRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) SetData(v bool) *UpdateGrafanaWorkspaceAccountRoleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) SetMessage(v string) *UpdateGrafanaWorkspaceAccountRoleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) SetRequestId(v string) *UpdateGrafanaWorkspaceAccountRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) SetSuccess(v bool) *UpdateGrafanaWorkspaceAccountRoleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) SetTraceId(v string) *UpdateGrafanaWorkspaceAccountRoleResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
