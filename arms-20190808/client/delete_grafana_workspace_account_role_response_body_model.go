// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceAccountRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGrafanaWorkspaceAccountRoleResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteGrafanaWorkspaceAccountRoleResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteGrafanaWorkspaceAccountRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGrafanaWorkspaceAccountRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGrafanaWorkspaceAccountRoleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteGrafanaWorkspaceAccountRoleResponseBody
	GetTraceId() *string
}

type DeleteGrafanaWorkspaceAccountRoleResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//         "RequestId": "01A079B0-7AA2-50A6-9A74-D89FE01432A5",
	//
	//         "TraceId": "0a03282e17887480040532718e74b7",
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
	// 22614CC7-7EA5-576F-9536-28717A886EB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 0a06dfe917788110578173646e4852
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteGrafanaWorkspaceAccountRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceAccountRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) SetCode(v int32) *DeleteGrafanaWorkspaceAccountRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) SetData(v bool) *DeleteGrafanaWorkspaceAccountRoleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) SetMessage(v string) *DeleteGrafanaWorkspaceAccountRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) SetRequestId(v string) *DeleteGrafanaWorkspaceAccountRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) SetSuccess(v bool) *DeleteGrafanaWorkspaceAccountRoleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) SetTraceId(v string) *DeleteGrafanaWorkspaceAccountRoleResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
