// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentResponseBody) *GetDocumentResponse
	GetBody() *GetDocumentResponseBody
}

type GetDocumentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentResponse) GetBody() *GetDocumentResponseBody {
	return s.Body
}

func (s *GetDocumentResponse) SetHeaders(v map[string]*string) *GetDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentResponse) SetStatusCode(v int32) *GetDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentResponse) SetBody(v *GetDocumentResponseBody) *GetDocumentResponse {
	s.Body = v
	return s
}

func (s *GetDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
