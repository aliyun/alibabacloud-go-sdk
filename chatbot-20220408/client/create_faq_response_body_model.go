// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaqResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeId(v int64) *CreateFaqResponseBody
	GetKnowledgeId() *int64
	SetRequestId(v string) *CreateFaqResponseBody
	GetRequestId() *string
}

type CreateFaqResponseBody struct {
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 28805A7C-D695-548C-A31B-67E52C2C274F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFaqResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFaqResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaqResponseBody) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *CreateFaqResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFaqResponseBody) SetKnowledgeId(v int64) *CreateFaqResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *CreateFaqResponseBody) SetRequestId(v string) *CreateFaqResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFaqResponseBody) Validate() error {
	return dara.Validate(s)
}
