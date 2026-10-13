// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleOidcCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleOidcCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleOidcCallbackResponse
	GetStatusCode() *int32
	SetBody(v *HandleOidcCallbackResponseBody) *HandleOidcCallbackResponse
	GetBody() *HandleOidcCallbackResponseBody
}

type HandleOidcCallbackResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleOidcCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleOidcCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleOidcCallbackResponse) GoString() string {
	return s.String()
}

func (s *HandleOidcCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleOidcCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleOidcCallbackResponse) GetBody() *HandleOidcCallbackResponseBody {
	return s.Body
}

func (s *HandleOidcCallbackResponse) SetHeaders(v map[string]*string) *HandleOidcCallbackResponse {
	s.Headers = v
	return s
}

func (s *HandleOidcCallbackResponse) SetStatusCode(v int32) *HandleOidcCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleOidcCallbackResponse) SetBody(v *HandleOidcCallbackResponseBody) *HandleOidcCallbackResponse {
	s.Body = v
	return s
}

func (s *HandleOidcCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
