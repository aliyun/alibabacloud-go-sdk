// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiModelKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbClusterId(v string) *CreateMultiModelKnowledgeBaseResponseBody
	GetDbClusterId() *string
	SetMmkbName(v string) *CreateMultiModelKnowledgeBaseResponseBody
	GetMmkbName() *string
	SetRequestId(v string) *CreateMultiModelKnowledgeBaseResponseBody
	GetRequestId() *string
}

type CreateMultiModelKnowledgeBaseResponseBody struct {
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	MmkbName    *string `json:"MmkbName,omitempty" xml:"MmkbName,omitempty"`
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

func (s *CreateMultiModelKnowledgeBaseResponseBody) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) GetMmkbName() *string {
	return s.MmkbName
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) SetDbClusterId(v string) *CreateMultiModelKnowledgeBaseResponseBody {
	s.DbClusterId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) SetMmkbName(v string) *CreateMultiModelKnowledgeBaseResponseBody {
	s.MmkbName = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) SetRequestId(v string) *CreateMultiModelKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
