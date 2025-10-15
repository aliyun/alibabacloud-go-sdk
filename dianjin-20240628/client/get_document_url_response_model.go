// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentUrlResponseBody) *GetDocumentUrlResponse
	GetBody() *GetDocumentUrlResponseBody
}

type GetDocumentUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentUrlResponse) GetBody() *GetDocumentUrlResponseBody {
	return s.Body
}

func (s *GetDocumentUrlResponse) SetHeaders(v map[string]*string) *GetDocumentUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentUrlResponse) SetStatusCode(v int32) *GetDocumentUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentUrlResponse) SetBody(v *GetDocumentUrlResponseBody) *GetDocumentUrlResponse {
	s.Body = v
	return s
}

func (s *GetDocumentUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
