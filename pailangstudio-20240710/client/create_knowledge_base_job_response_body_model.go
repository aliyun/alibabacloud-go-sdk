// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseJobId(v string) *CreateKnowledgeBaseJobResponseBody
	GetKnowledgeBaseJobId() *string
	SetRequestId(v string) *CreateKnowledgeBaseJobResponseBody
	GetRequestId() *string
}

type CreateKnowledgeBaseJobResponseBody struct {
	// example:
	//
	// kbjob-9mn******1z54
	KnowledgeBaseJobId *string `json:"KnowledgeBaseJobId,omitempty" xml:"KnowledgeBaseJobId,omitempty"`
	// example:
	//
	// 963BD7F9-0C02-5594-9550-BCC6DD43E3C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKnowledgeBaseJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseJobResponseBody) GetKnowledgeBaseJobId() *string {
	return s.KnowledgeBaseJobId
}

func (s *CreateKnowledgeBaseJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseJobResponseBody) SetKnowledgeBaseJobId(v string) *CreateKnowledgeBaseJobResponseBody {
	s.KnowledgeBaseJobId = &v
	return s
}

func (s *CreateKnowledgeBaseJobResponseBody) SetRequestId(v string) *CreateKnowledgeBaseJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseJobResponseBody) Validate() error {
	return dara.Validate(s)
}
