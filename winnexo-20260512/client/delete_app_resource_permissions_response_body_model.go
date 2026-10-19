// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResourcePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAppResourcePermissionsResponseBody
	GetCode() *string
	SetDeletedCount(v int64) *DeleteAppResourcePermissionsResponseBody
	GetDeletedCount() *int64
	SetMessage(v string) *DeleteAppResourcePermissionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAppResourcePermissionsResponseBody
	GetRequestId() *string
}

type DeleteAppResourcePermissionsResponseBody struct {
	// 业务状态码；成功为 200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 删除的权限记录数
	//
	// example:
	//
	// 2
	DeletedCount *int64 `json:"deletedCount,omitempty" xml:"deletedCount,omitempty"`
	// 错误描述；成功时为空
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAppResourcePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResourcePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResourcePermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAppResourcePermissionsResponseBody) GetDeletedCount() *int64 {
	return s.DeletedCount
}

func (s *DeleteAppResourcePermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppResourcePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppResourcePermissionsResponseBody) SetCode(v string) *DeleteAppResourcePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) SetDeletedCount(v int64) *DeleteAppResourcePermissionsResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) SetMessage(v string) *DeleteAppResourcePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) SetRequestId(v string) *DeleteAppResourcePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
