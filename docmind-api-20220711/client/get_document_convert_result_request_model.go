// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentConvertResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDocumentConvertResultRequest
	GetId() *string
}

type GetDocumentConvertResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocumentConvertResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentConvertResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultRequest) GetId() *string {
	return s.Id
}

func (s *GetDocumentConvertResultRequest) SetId(v string) *GetDocumentConvertResultRequest {
	s.Id = &v
	return s
}

func (s *GetDocumentConvertResultRequest) Validate() error {
	return dara.Validate(s)
}
