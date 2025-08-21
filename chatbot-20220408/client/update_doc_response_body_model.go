// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeId(v int64) *UpdateDocResponseBody
	GetKnowledgeId() *int64
	SetRequestId(v string) *UpdateDocResponseBody
	GetRequestId() *string
}

type UpdateDocResponseBody struct {
	// example:
	//
	// 30002406051
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0F9F136A-1BF6-5CC1-9D57-9717761F03B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocResponseBody) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *UpdateDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDocResponseBody) SetKnowledgeId(v int64) *UpdateDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateDocResponseBody) SetRequestId(v string) *UpdateDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocResponseBody) Validate() error {
	return dara.Validate(s)
}
