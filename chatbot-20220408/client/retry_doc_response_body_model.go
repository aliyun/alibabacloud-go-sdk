// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeId(v int64) *RetryDocResponseBody
	GetKnowledgeId() *int64
	SetRequestId(v string) *RetryDocResponseBody
	GetRequestId() *string
}

type RetryDocResponseBody struct {
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6419BA93-D111-5225-8998-13E63E6D3940
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryDocResponseBody) GoString() string {
	return s.String()
}

func (s *RetryDocResponseBody) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *RetryDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryDocResponseBody) SetKnowledgeId(v int64) *RetryDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *RetryDocResponseBody) SetRequestId(v string) *RetryDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryDocResponseBody) Validate() error {
	return dara.Validate(s)
}
