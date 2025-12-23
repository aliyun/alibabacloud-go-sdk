// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SignResponse
	GetStatusCode() *int32
	SetBody(v *SignResponseBody) *SignResponse
	GetBody() *SignResponseBody
}

type SignResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignResponse) String() string {
	return dara.Prettify(s)
}

func (s SignResponse) GoString() string {
	return s.String()
}

func (s *SignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SignResponse) GetBody() *SignResponseBody {
	return s.Body
}

func (s *SignResponse) SetHeaders(v map[string]*string) *SignResponse {
	s.Headers = v
	return s
}

func (s *SignResponse) SetStatusCode(v int32) *SignResponse {
	s.StatusCode = &v
	return s
}

func (s *SignResponse) SetBody(v *SignResponseBody) *SignResponse {
	s.Body = v
	return s
}

func (s *SignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
