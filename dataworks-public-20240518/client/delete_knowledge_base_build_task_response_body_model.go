// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseBuildTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteKnowledgeBaseBuildTaskResponseBody
	GetCode() *string
	SetData(v *DeleteKnowledgeBaseBuildTaskResponseBodyData) *DeleteKnowledgeBaseBuildTaskResponseBody
	GetData() *DeleteKnowledgeBaseBuildTaskResponseBodyData
	SetMessage(v string) *DeleteKnowledgeBaseBuildTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteKnowledgeBaseBuildTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKnowledgeBaseBuildTaskResponseBody
	GetSuccess() *bool
}

type DeleteKnowledgeBaseBuildTaskResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *DeleteKnowledgeBaseBuildTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The security error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKnowledgeBaseBuildTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) GetData() *DeleteKnowledgeBaseBuildTaskResponseBodyData {
	return s.Data
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) SetCode(v string) *DeleteKnowledgeBaseBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) SetData(v *DeleteKnowledgeBaseBuildTaskResponseBodyData) *DeleteKnowledgeBaseBuildTaskResponseBody {
	s.Data = v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) SetMessage(v string) *DeleteKnowledgeBaseBuildTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) SetRequestId(v string) *DeleteKnowledgeBaseBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) SetSuccess(v bool) *DeleteKnowledgeBaseBuildTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteKnowledgeBaseBuildTaskResponseBodyData struct {
	// The ID of the knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the source task.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteKnowledgeBaseBuildTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseBuildTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBodyData) SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseBuildTaskResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBodyData) SetTaskId(v string) *DeleteKnowledgeBaseBuildTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
