// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseChunkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKnowledgeBaseChunkResponseBody
	GetRequestId() *string
}

type UpdateKnowledgeBaseChunkResponseBody struct {
	// example:
	//
	// 963BD7F9-0C02-5594-9550-BCC6DD43E3C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKnowledgeBaseChunkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseChunkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseChunkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseChunkResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseChunkResponseBody) Validate() error {
	return dara.Validate(s)
}
