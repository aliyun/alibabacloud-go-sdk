// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiModalRerankerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocs(v []*GetMultiModalRerankerRequestDocs) *GetMultiModalRerankerRequest
	GetDocs() []*GetMultiModalRerankerRequestDocs
	SetOptions(v map[string]interface{}) *GetMultiModalRerankerRequest
	GetOptions() map[string]interface{}
	SetQuery(v *GetMultiModalRerankerRequestQuery) *GetMultiModalRerankerRequest
	GetQuery() *GetMultiModalRerankerRequestQuery
}

type GetMultiModalRerankerRequest struct {
	Docs    []*GetMultiModalRerankerRequestDocs `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	Options map[string]interface{}              `json:"options,omitempty" xml:"options,omitempty"`
	Query   *GetMultiModalRerankerRequestQuery  `json:"query,omitempty" xml:"query,omitempty" type:"Struct"`
}

func (s GetMultiModalRerankerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerRequest) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerRequest) GetDocs() []*GetMultiModalRerankerRequestDocs {
	return s.Docs
}

func (s *GetMultiModalRerankerRequest) GetOptions() map[string]interface{} {
	return s.Options
}

func (s *GetMultiModalRerankerRequest) GetQuery() *GetMultiModalRerankerRequestQuery {
	return s.Query
}

func (s *GetMultiModalRerankerRequest) SetDocs(v []*GetMultiModalRerankerRequestDocs) *GetMultiModalRerankerRequest {
	s.Docs = v
	return s
}

func (s *GetMultiModalRerankerRequest) SetOptions(v map[string]interface{}) *GetMultiModalRerankerRequest {
	s.Options = v
	return s
}

func (s *GetMultiModalRerankerRequest) SetQuery(v *GetMultiModalRerankerRequestQuery) *GetMultiModalRerankerRequest {
	s.Query = v
	return s
}

func (s *GetMultiModalRerankerRequest) Validate() error {
	if s.Docs != nil {
		for _, item := range s.Docs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Query != nil {
		if err := s.Query.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiModalRerankerRequestDocs struct {
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetMultiModalRerankerRequestDocs) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerRequestDocs) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerRequestDocs) GetImage() *string {
	return s.Image
}

func (s *GetMultiModalRerankerRequestDocs) GetText() *string {
	return s.Text
}

func (s *GetMultiModalRerankerRequestDocs) SetImage(v string) *GetMultiModalRerankerRequestDocs {
	s.Image = &v
	return s
}

func (s *GetMultiModalRerankerRequestDocs) SetText(v string) *GetMultiModalRerankerRequestDocs {
	s.Text = &v
	return s
}

func (s *GetMultiModalRerankerRequestDocs) Validate() error {
	return dara.Validate(s)
}

type GetMultiModalRerankerRequestQuery struct {
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetMultiModalRerankerRequestQuery) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerRequestQuery) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerRequestQuery) GetImage() *string {
	return s.Image
}

func (s *GetMultiModalRerankerRequestQuery) GetText() *string {
	return s.Text
}

func (s *GetMultiModalRerankerRequestQuery) SetImage(v string) *GetMultiModalRerankerRequestQuery {
	s.Image = &v
	return s
}

func (s *GetMultiModalRerankerRequestQuery) SetText(v string) *GetMultiModalRerankerRequestQuery {
	s.Text = &v
	return s
}

func (s *GetMultiModalRerankerRequestQuery) Validate() error {
	return dara.Validate(s)
}
