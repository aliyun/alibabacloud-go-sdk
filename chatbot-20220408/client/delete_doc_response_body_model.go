// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeId(v int64) *DeleteDocResponseBody
	GetKnowledgeId() *int64
	SetRequestId(v string) *DeleteDocResponseBody
	GetRequestId() *string
}

type DeleteDocResponseBody struct {
	// example:
	//
	// 30002406051
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// DFB71B34-4188-4EA2-9988-EF3014E75910
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocResponseBody) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *DeleteDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocResponseBody) SetKnowledgeId(v int64) *DeleteDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *DeleteDocResponseBody) SetRequestId(v string) *DeleteDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocResponseBody) Validate() error {
	return dara.Validate(s)
}
