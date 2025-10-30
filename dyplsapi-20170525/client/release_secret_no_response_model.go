// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSecretNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseSecretNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseSecretNoResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseSecretNoResponseBody) *ReleaseSecretNoResponse
	GetBody() *ReleaseSecretNoResponseBody
}

type ReleaseSecretNoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseSecretNoResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSecretNoResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseSecretNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseSecretNoResponse) GetBody() *ReleaseSecretNoResponseBody {
	return s.Body
}

func (s *ReleaseSecretNoResponse) SetHeaders(v map[string]*string) *ReleaseSecretNoResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSecretNoResponse) SetStatusCode(v int32) *ReleaseSecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseSecretNoResponse) SetBody(v *ReleaseSecretNoResponseBody) *ReleaseSecretNoResponse {
	s.Body = v
	return s
}

func (s *ReleaseSecretNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
