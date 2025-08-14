// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentCompareResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDocumentCompareResultRequest
	GetId() *string
}

type GetDocumentCompareResultRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocumentCompareResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentCompareResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentCompareResultRequest) GetId() *string {
	return s.Id
}

func (s *GetDocumentCompareResultRequest) SetId(v string) *GetDocumentCompareResultRequest {
	s.Id = &v
	return s
}

func (s *GetDocumentCompareResultRequest) Validate() error {
	return dara.Validate(s)
}
