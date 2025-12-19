// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteResourceTokenAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteResourceTokenAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteResourceTokenAuthResponse
	GetStatusCode() *int32
	SetBody(v *CompleteResourceTokenAuthResponseBody) *CompleteResourceTokenAuthResponse
	GetBody() *CompleteResourceTokenAuthResponseBody
}

type CompleteResourceTokenAuthResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteResourceTokenAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteResourceTokenAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteResourceTokenAuthResponse) GoString() string {
	return s.String()
}

func (s *CompleteResourceTokenAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteResourceTokenAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteResourceTokenAuthResponse) GetBody() *CompleteResourceTokenAuthResponseBody {
	return s.Body
}

func (s *CompleteResourceTokenAuthResponse) SetHeaders(v map[string]*string) *CompleteResourceTokenAuthResponse {
	s.Headers = v
	return s
}

func (s *CompleteResourceTokenAuthResponse) SetStatusCode(v int32) *CompleteResourceTokenAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteResourceTokenAuthResponse) SetBody(v *CompleteResourceTokenAuthResponseBody) *CompleteResourceTokenAuthResponse {
	s.Body = v
	return s
}

func (s *CompleteResourceTokenAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
