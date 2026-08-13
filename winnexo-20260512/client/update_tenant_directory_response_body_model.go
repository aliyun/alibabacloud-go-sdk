// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTenantDirectoryResponseBody
	GetCode() *string
	SetDirectoryId(v string) *UpdateTenantDirectoryResponseBody
	GetDirectoryId() *string
	SetMessage(v string) *UpdateTenantDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTenantDirectoryResponseBody
	GetRequestId() *string
}

type UpdateTenantDirectoryResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 已修改的目录唯一标识
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateTenantDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTenantDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTenantDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateTenantDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTenantDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTenantDirectoryResponseBody) SetCode(v string) *UpdateTenantDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTenantDirectoryResponseBody) SetDirectoryId(v string) *UpdateTenantDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *UpdateTenantDirectoryResponseBody) SetMessage(v string) *UpdateTenantDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTenantDirectoryResponseBody) SetRequestId(v string) *UpdateTenantDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTenantDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
