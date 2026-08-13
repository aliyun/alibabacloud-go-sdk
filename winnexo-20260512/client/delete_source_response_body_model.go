// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSourceResponseBody
	GetCode() *string
	SetDeleted(v bool) *DeleteSourceResponseBody
	GetDeleted() *bool
	SetMessage(v string) *DeleteSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSourceResponseBody
	GetRequestId() *string
	SetSourceId(v string) *DeleteSourceResponseBody
	GetSourceId() *string
}

type DeleteSourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否已删除
	//
	// example:
	//
	// true
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 被删除的数据源 ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s DeleteSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSourceResponseBody) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeleteSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *DeleteSourceResponseBody) SetCode(v string) *DeleteSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSourceResponseBody) SetDeleted(v bool) *DeleteSourceResponseBody {
	s.Deleted = &v
	return s
}

func (s *DeleteSourceResponseBody) SetMessage(v string) *DeleteSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSourceResponseBody) SetRequestId(v string) *DeleteSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSourceResponseBody) SetSourceId(v string) *DeleteSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *DeleteSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
