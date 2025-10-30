// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretNoRemainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySecretNoRemainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySecretNoRemainResponse
	GetStatusCode() *int32
	SetBody(v *QuerySecretNoRemainResponseBody) *QuerySecretNoRemainResponse
	GetBody() *QuerySecretNoRemainResponseBody
}

type QuerySecretNoRemainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecretNoRemainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecretNoRemainResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoRemainResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySecretNoRemainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySecretNoRemainResponse) GetBody() *QuerySecretNoRemainResponseBody {
	return s.Body
}

func (s *QuerySecretNoRemainResponse) SetHeaders(v map[string]*string) *QuerySecretNoRemainResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretNoRemainResponse) SetStatusCode(v int32) *QuerySecretNoRemainResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecretNoRemainResponse) SetBody(v *QuerySecretNoRemainResponseBody) *QuerySecretNoRemainResponse {
	s.Body = v
	return s
}

func (s *QuerySecretNoRemainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
