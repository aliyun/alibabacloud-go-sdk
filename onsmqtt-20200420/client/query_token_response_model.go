// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTokenResponse
	GetStatusCode() *int32
	SetBody(v *QueryTokenResponseBody) *QueryTokenResponse
	GetBody() *QueryTokenResponseBody
}

type QueryTokenResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTokenResponse) GoString() string {
	return s.String()
}

func (s *QueryTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTokenResponse) GetBody() *QueryTokenResponseBody {
	return s.Body
}

func (s *QueryTokenResponse) SetHeaders(v map[string]*string) *QueryTokenResponse {
	s.Headers = v
	return s
}

func (s *QueryTokenResponse) SetStatusCode(v int32) *QueryTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTokenResponse) SetBody(v *QueryTokenResponseBody) *QueryTokenResponse {
	s.Body = v
	return s
}

func (s *QueryTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
