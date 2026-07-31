// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddKnowledgeFileResponseBody
	GetRequestId() *string
}

type AddKnowledgeFileResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddKnowledgeFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddKnowledgeFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKnowledgeFileResponseBody) SetRequestId(v string) *AddKnowledgeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKnowledgeFileResponseBody) Validate() error {
	return dara.Validate(s)
}
