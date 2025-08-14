// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentExtractResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDocumentExtractResultRequest
	GetId() *string
}

type GetDocumentExtractResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocumentExtractResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentExtractResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentExtractResultRequest) GetId() *string {
	return s.Id
}

func (s *GetDocumentExtractResultRequest) SetId(v string) *GetDocumentExtractResultRequest {
	s.Id = &v
	return s
}

func (s *GetDocumentExtractResultRequest) Validate() error {
	return dara.Validate(s)
}
