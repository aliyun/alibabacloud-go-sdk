// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentExtractionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocumentExtractionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocumentExtractionResponse
	GetStatusCode() *int32
	SetBody(v *DocumentExtractionResponseBody) *DocumentExtractionResponse
	GetBody() *DocumentExtractionResponseBody
}

type DocumentExtractionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocumentExtractionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocumentExtractionResponse) String() string {
	return dara.Prettify(s)
}

func (s DocumentExtractionResponse) GoString() string {
	return s.String()
}

func (s *DocumentExtractionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocumentExtractionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocumentExtractionResponse) GetBody() *DocumentExtractionResponseBody {
	return s.Body
}

func (s *DocumentExtractionResponse) SetHeaders(v map[string]*string) *DocumentExtractionResponse {
	s.Headers = v
	return s
}

func (s *DocumentExtractionResponse) SetStatusCode(v int32) *DocumentExtractionResponse {
	s.StatusCode = &v
	return s
}

func (s *DocumentExtractionResponse) SetBody(v *DocumentExtractionResponseBody) *DocumentExtractionResponse {
	s.Body = v
	return s
}

func (s *DocumentExtractionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
