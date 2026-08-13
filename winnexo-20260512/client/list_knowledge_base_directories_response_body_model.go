// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBaseDirectoriesResponseBody
	GetCode() *string
	SetDirectories(v []interface{}) *ListKnowledgeBaseDirectoriesResponseBody
	GetDirectories() []interface{}
	SetMessage(v string) *ListKnowledgeBaseDirectoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeBaseDirectoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListKnowledgeBaseDirectoriesResponseBody
	GetTotalCount() *int64
}

type ListKnowledgeBaseDirectoriesResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// directories
	//
	// example:
	//
	// string_value
	Directories []interface{} `json:"directories,omitempty" xml:"directories,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 分类总数
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListKnowledgeBaseDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) GetDirectories() []interface{} {
	return s.Directories
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) SetCode(v string) *ListKnowledgeBaseDirectoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) SetDirectories(v []interface{}) *ListKnowledgeBaseDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) SetMessage(v string) *ListKnowledgeBaseDirectoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) SetRequestId(v string) *ListKnowledgeBaseDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) SetTotalCount(v int64) *ListKnowledgeBaseDirectoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
