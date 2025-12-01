// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentPageResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentPageResponseBody) *GetDocumentPageResponse
	GetBody() *GetDocumentPageResponseBody
}

type GetDocumentPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentPageResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentPageResponse) GetBody() *GetDocumentPageResponseBody {
	return s.Body
}

func (s *GetDocumentPageResponse) SetHeaders(v map[string]*string) *GetDocumentPageResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentPageResponse) SetStatusCode(v int32) *GetDocumentPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentPageResponse) SetBody(v *GetDocumentPageResponseBody) *GetDocumentPageResponse {
	s.Body = v
	return s
}

func (s *GetDocumentPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
