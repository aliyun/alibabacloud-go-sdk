// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKnowledgeBaseJobResponseBody
	GetRequestId() *string
}

type DeleteKnowledgeBaseJobResponseBody struct {
	// example:
	//
	// 48E6392E-C3C9-5212-9FAD-13256ABD9AF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKnowledgeBaseJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeBaseJobResponseBody) SetRequestId(v string) *DeleteKnowledgeBaseJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeBaseJobResponseBody) Validate() error {
	return dara.Validate(s)
}
