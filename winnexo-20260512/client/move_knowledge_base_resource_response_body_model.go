// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveKnowledgeBaseResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MoveKnowledgeBaseResourceResponseBody
	GetCode() *string
	SetKnowledgeId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetKnowledgeId() *string
	SetMessage(v string) *MoveKnowledgeBaseResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetRequestId() *string
	SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetTargetDirectoryId() *string
}

type MoveKnowledgeBaseResourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 目标知识库 ID（echo 回入参）
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
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
	// 目标目录 ID（echo 回入参）
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
}

func (s MoveKnowledgeBaseResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveKnowledgeBaseResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetCode(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.Code = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetKnowledgeId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetMessage(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.Message = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetRequestId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetSourceId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
