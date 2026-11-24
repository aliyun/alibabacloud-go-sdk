// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *AddKnowledgeFileResponseBody
	GetId() *int64
	SetLocation(v string) *AddKnowledgeFileResponseBody
	GetLocation() *string
	SetRequestId(v string) *AddKnowledgeFileResponseBody
	GetRequestId() *string
}

type AddKnowledgeFileResponseBody struct {
	// The ID of the imported file.
	//
	// example:
	//
	// 2059412857631539200
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The location of the imported file.
	//
	// example:
	//
	// oss://bucket_name/path/to/file
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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

func (s *AddKnowledgeFileResponseBody) GetId() *int64 {
	return s.Id
}

func (s *AddKnowledgeFileResponseBody) GetLocation() *string {
	return s.Location
}

func (s *AddKnowledgeFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKnowledgeFileResponseBody) SetId(v int64) *AddKnowledgeFileResponseBody {
	s.Id = &v
	return s
}

func (s *AddKnowledgeFileResponseBody) SetLocation(v string) *AddKnowledgeFileResponseBody {
	s.Location = &v
	return s
}

func (s *AddKnowledgeFileResponseBody) SetRequestId(v string) *AddKnowledgeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKnowledgeFileResponseBody) Validate() error {
	return dara.Validate(s)
}
