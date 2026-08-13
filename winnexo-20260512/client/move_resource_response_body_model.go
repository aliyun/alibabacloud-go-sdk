// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MoveResourceResponseBody
	GetCode() *string
	SetMessage(v string) *MoveResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MoveResourceResponseBody
	GetRequestId() *string
	SetSourceDirectoryId(v string) *MoveResourceResponseBody
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveResourceResponseBody
	GetSourceId() *string
	SetSuccess(v bool) *MoveResourceResponseBody
	GetSuccess() *bool
	SetTargetDirectoryId(v string) *MoveResourceResponseBody
	GetTargetDirectoryId() *string
}

type MoveResourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 源目录 ID（echo 回入参）
	//
	// example:
	//
	// exampleSourceDirectoryId
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// 资源 ID（echo 回入参）
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 操作是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 目标目录 ID（echo 回入参）
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
}

func (s MoveResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourceResponseBody) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveResourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveResourceResponseBody) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveResourceResponseBody) SetCode(v string) *MoveResourceResponseBody {
	s.Code = &v
	return s
}

func (s *MoveResourceResponseBody) SetMessage(v string) *MoveResourceResponseBody {
	s.Message = &v
	return s
}

func (s *MoveResourceResponseBody) SetRequestId(v string) *MoveResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSourceDirectoryId(v string) *MoveResourceResponseBody {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSourceId(v string) *MoveResourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSuccess(v bool) *MoveResourceResponseBody {
	s.Success = &v
	return s
}

func (s *MoveResourceResponseBody) SetTargetDirectoryId(v string) *MoveResourceResponseBody {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
