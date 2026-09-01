// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKnowledgeBaseFileResponseBody
	GetRequestId() *string
}

type DeleteKnowledgeBaseFileResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKnowledgeBaseFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeBaseFileResponseBody) SetRequestId(v string) *DeleteKnowledgeBaseFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeBaseFileResponseBody) Validate() error {
	return dara.Validate(s)
}
