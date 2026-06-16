// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocumentTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocumentTranslateResponse
	GetStatusCode() *int32
	SetBody(v *DocumentTranslateResponseBody) *DocumentTranslateResponse
	GetBody() *DocumentTranslateResponseBody
}

type DocumentTranslateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocumentTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocumentTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s DocumentTranslateResponse) GoString() string {
	return s.String()
}

func (s *DocumentTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocumentTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocumentTranslateResponse) GetBody() *DocumentTranslateResponseBody {
	return s.Body
}

func (s *DocumentTranslateResponse) SetHeaders(v map[string]*string) *DocumentTranslateResponse {
	s.Headers = v
	return s
}

func (s *DocumentTranslateResponse) SetStatusCode(v int32) *DocumentTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *DocumentTranslateResponse) SetBody(v *DocumentTranslateResponseBody) *DocumentTranslateResponse {
	s.Body = v
	return s
}

func (s *DocumentTranslateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
