// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReparseSourceResponseBody
	GetCode() *string
	SetMessage(v string) *ReparseSourceResponseBody
	GetMessage() *string
	SetName(v string) *ReparseSourceResponseBody
	GetName() *string
	SetRequestId(v string) *ReparseSourceResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReparseSourceResponseBody
	GetSourceId() *string
	SetSourceType(v string) *ReparseSourceResponseBody
	GetSourceType() *string
	SetStatus(v string) *ReparseSourceResponseBody
	GetStatus() *string
}

type ReparseSourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 数据源 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 数据源类型
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 重新解析后的数据源状态
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReparseSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReparseSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ReparseSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReparseSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReparseSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *ReparseSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReparseSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseSourceResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *ReparseSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ReparseSourceResponseBody) SetCode(v string) *ReparseSourceResponseBody {
	s.Code = &v
	return s
}

func (s *ReparseSourceResponseBody) SetMessage(v string) *ReparseSourceResponseBody {
	s.Message = &v
	return s
}

func (s *ReparseSourceResponseBody) SetName(v string) *ReparseSourceResponseBody {
	s.Name = &v
	return s
}

func (s *ReparseSourceResponseBody) SetRequestId(v string) *ReparseSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReparseSourceResponseBody) SetSourceId(v string) *ReparseSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReparseSourceResponseBody) SetSourceType(v string) *ReparseSourceResponseBody {
	s.SourceType = &v
	return s
}

func (s *ReparseSourceResponseBody) SetStatus(v string) *ReparseSourceResponseBody {
	s.Status = &v
	return s
}

func (s *ReparseSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
