// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAutoMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClearedCount(v int64) *ClearAutoMemoryResponseBody
	GetClearedCount() *int64
	SetCode(v string) *ClearAutoMemoryResponseBody
	GetCode() *string
	SetMessage(v string) *ClearAutoMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClearAutoMemoryResponseBody
	GetRequestId() *string
	SetType(v string) *ClearAutoMemoryResponseBody
	GetType() *string
}

type ClearAutoMemoryResponseBody struct {
	// 本次成功删除的自动记忆文件数量；无可清空记忆时为 0
	//
	// example:
	//
	// 3
	ClearedCount *int64 `json:"clearedCount,omitempty" xml:"clearedCount,omitempty"`
	// 业务状态码；成功为 200，失败时为对应错误码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述；成功时为空字符串，失败时返回具体原因
	//
	// example:
	//
	// 当前用户不是群组 cg_12345 的活跃成员
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID，用于定位本次调用
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 本次清空的自动记忆类型；取值为 user、feedback、reference、project 或 all
	//
	// example:
	//
	// project
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ClearAutoMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearAutoMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ClearAutoMemoryResponseBody) GetClearedCount() *int64 {
	return s.ClearedCount
}

func (s *ClearAutoMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClearAutoMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClearAutoMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearAutoMemoryResponseBody) GetType() *string {
	return s.Type
}

func (s *ClearAutoMemoryResponseBody) SetClearedCount(v int64) *ClearAutoMemoryResponseBody {
	s.ClearedCount = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetCode(v string) *ClearAutoMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetMessage(v string) *ClearAutoMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetRequestId(v string) *ClearAutoMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) SetType(v string) *ClearAutoMemoryResponseBody {
	s.Type = &v
	return s
}

func (s *ClearAutoMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
