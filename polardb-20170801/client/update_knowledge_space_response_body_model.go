// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeSpaceId(v string) *UpdateKnowledgeSpaceResponseBody
	GetKnowledgeSpaceId() *string
	SetRequestId(v string) *UpdateKnowledgeSpaceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *UpdateKnowledgeSpaceResponseBody
	GetTaskId() *int32
}

type UpdateKnowledgeSpaceResponseBody struct {
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 52*****03
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateKnowledgeSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceResponseBody) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *UpdateKnowledgeSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeSpaceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *UpdateKnowledgeSpaceResponseBody) SetKnowledgeSpaceId(v string) *UpdateKnowledgeSpaceResponseBody {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *UpdateKnowledgeSpaceResponseBody) SetRequestId(v string) *UpdateKnowledgeSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeSpaceResponseBody) SetTaskId(v int32) *UpdateKnowledgeSpaceResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateKnowledgeSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
