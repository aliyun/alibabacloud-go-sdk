// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiModelKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMultiModelKnowledgeBaseResponseBody
	GetRequestId() *string
}

type CreateMultiModelKnowledgeBaseResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiModelKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiModelKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) SetRequestId(v string) *CreateMultiModelKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
