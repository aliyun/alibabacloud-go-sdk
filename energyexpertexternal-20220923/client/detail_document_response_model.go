// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetailDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetailDocumentResponse
	GetStatusCode() *int32
	SetBody(v *DetailDocumentResponseBody) *DetailDocumentResponse
	GetBody() *DetailDocumentResponseBody
}

type DetailDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s DetailDocumentResponse) GoString() string {
	return s.String()
}

func (s *DetailDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetailDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetailDocumentResponse) GetBody() *DetailDocumentResponseBody {
	return s.Body
}

func (s *DetailDocumentResponse) SetHeaders(v map[string]*string) *DetailDocumentResponse {
	s.Headers = v
	return s
}

func (s *DetailDocumentResponse) SetStatusCode(v int32) *DetailDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailDocumentResponse) SetBody(v *DetailDocumentResponseBody) *DetailDocumentResponse {
	s.Body = v
	return s
}

func (s *DetailDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
