// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKnowledgeBaseResponseBody
	GetRequestId() *string
}

type DeleteKnowledgeBaseResponseBody struct {
	// example:
	//
	// C25324E3-18E6-50D8-9026-16D74AAEEB26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeBaseResponseBody) SetRequestId(v string) *DeleteKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
