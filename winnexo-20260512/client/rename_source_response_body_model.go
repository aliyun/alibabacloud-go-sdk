// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenameSourceResponseBody
	GetCode() *string
	SetGmtModified(v string) *RenameSourceResponseBody
	GetGmtModified() *string
	SetMessage(v string) *RenameSourceResponseBody
	GetMessage() *string
	SetName(v string) *RenameSourceResponseBody
	GetName() *string
	SetRequestId(v string) *RenameSourceResponseBody
	GetRequestId() *string
	SetSourceId(v string) *RenameSourceResponseBody
	GetSourceId() *string
	SetStatus(v string) *RenameSourceResponseBody
	GetStatus() *string
}

type RenameSourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 最近修改时间，ISO8601 格式
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
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
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 数据源状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RenameSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameSourceResponseBody) GoString() string {
	return s.String()
}

func (s *RenameSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenameSourceResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *RenameSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenameSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *RenameSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *RenameSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RenameSourceResponseBody) SetCode(v string) *RenameSourceResponseBody {
	s.Code = &v
	return s
}

func (s *RenameSourceResponseBody) SetGmtModified(v string) *RenameSourceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *RenameSourceResponseBody) SetMessage(v string) *RenameSourceResponseBody {
	s.Message = &v
	return s
}

func (s *RenameSourceResponseBody) SetName(v string) *RenameSourceResponseBody {
	s.Name = &v
	return s
}

func (s *RenameSourceResponseBody) SetRequestId(v string) *RenameSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameSourceResponseBody) SetSourceId(v string) *RenameSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *RenameSourceResponseBody) SetStatus(v string) *RenameSourceResponseBody {
	s.Status = &v
	return s
}

func (s *RenameSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
