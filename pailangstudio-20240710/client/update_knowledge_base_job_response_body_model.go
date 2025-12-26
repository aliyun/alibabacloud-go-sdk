// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKnowledgeBaseJobResponseBody
	GetRequestId() *string
}

type UpdateKnowledgeBaseJobResponseBody struct {
	// example:
	//
	// 35686626-8D83-5ADE-B400-08A6613A6057
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKnowledgeBaseJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseJobResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseJobResponseBody) Validate() error {
	return dara.Validate(s)
}
