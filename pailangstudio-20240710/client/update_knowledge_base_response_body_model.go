// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKnowledgeBaseResponseBody
	GetRequestId() *string
}

type UpdateKnowledgeBaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
