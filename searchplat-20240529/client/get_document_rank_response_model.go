// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentRankResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentRankResponseBody) *GetDocumentRankResponse
	GetBody() *GetDocumentRankResponseBody
}

type GetDocumentRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentRankResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRankResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentRankResponse) GetBody() *GetDocumentRankResponseBody {
	return s.Body
}

func (s *GetDocumentRankResponse) SetHeaders(v map[string]*string) *GetDocumentRankResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentRankResponse) SetStatusCode(v int32) *GetDocumentRankResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentRankResponse) SetBody(v *GetDocumentRankResponseBody) *GetDocumentRankResponse {
	s.Body = v
	return s
}

func (s *GetDocumentRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
