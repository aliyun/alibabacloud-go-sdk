// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocs(v []*string) *GetDocumentRankRequest
	GetDocs() []*string
	SetQuery(v string) *GetDocumentRankRequest
	GetQuery() *string
}

type GetDocumentRankRequest struct {
	// This parameter is required.
	Docs []*string `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetDocumentRankRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRankRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentRankRequest) GetDocs() []*string {
	return s.Docs
}

func (s *GetDocumentRankRequest) GetQuery() *string {
	return s.Query
}

func (s *GetDocumentRankRequest) SetDocs(v []*string) *GetDocumentRankRequest {
	s.Docs = v
	return s
}

func (s *GetDocumentRankRequest) SetQuery(v string) *GetDocumentRankRequest {
	s.Query = &v
	return s
}

func (s *GetDocumentRankRequest) Validate() error {
	return dara.Validate(s)
}
