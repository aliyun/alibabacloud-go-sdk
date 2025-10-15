// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainApplicationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainApplicationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainApplicationTokenResponse
	GetStatusCode() *int32
	SetBody(v *ObtainApplicationTokenResponseBody) *ObtainApplicationTokenResponse
	GetBody() *ObtainApplicationTokenResponseBody
}

type ObtainApplicationTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainApplicationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainApplicationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationTokenResponse) GoString() string {
	return s.String()
}

func (s *ObtainApplicationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainApplicationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainApplicationTokenResponse) GetBody() *ObtainApplicationTokenResponseBody {
	return s.Body
}

func (s *ObtainApplicationTokenResponse) SetHeaders(v map[string]*string) *ObtainApplicationTokenResponse {
	s.Headers = v
	return s
}

func (s *ObtainApplicationTokenResponse) SetStatusCode(v int32) *ObtainApplicationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainApplicationTokenResponse) SetBody(v *ObtainApplicationTokenResponseBody) *ObtainApplicationTokenResponse {
	s.Body = v
	return s
}

func (s *ObtainApplicationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
