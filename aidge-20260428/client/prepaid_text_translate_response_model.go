// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepaidTextTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrepaidTextTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrepaidTextTranslateResponse
	GetStatusCode() *int32
	SetBody(v *PrepaidTextTranslateResponseBody) *PrepaidTextTranslateResponse
	GetBody() *PrepaidTextTranslateResponseBody
}

type PrepaidTextTranslateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrepaidTextTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrepaidTextTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s PrepaidTextTranslateResponse) GoString() string {
	return s.String()
}

func (s *PrepaidTextTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrepaidTextTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrepaidTextTranslateResponse) GetBody() *PrepaidTextTranslateResponseBody {
	return s.Body
}

func (s *PrepaidTextTranslateResponse) SetHeaders(v map[string]*string) *PrepaidTextTranslateResponse {
	s.Headers = v
	return s
}

func (s *PrepaidTextTranslateResponse) SetStatusCode(v int32) *PrepaidTextTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *PrepaidTextTranslateResponse) SetBody(v *PrepaidTextTranslateResponseBody) *PrepaidTextTranslateResponse {
	s.Body = v
	return s
}

func (s *PrepaidTextTranslateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
