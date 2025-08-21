// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeId(v int64) *CreateDocResponseBody
	GetKnowledgeId() *int64
	SetRequestId(v string) *CreateDocResponseBody
	GetRequestId() *string
}

type CreateDocResponseBody struct {
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 07B270A4-61D8-57F6-A609-A3C216CFB872
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocResponseBody) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *CreateDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDocResponseBody) SetKnowledgeId(v int64) *CreateDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *CreateDocResponseBody) SetRequestId(v string) *CreateDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocResponseBody) Validate() error {
	return dara.Validate(s)
}
