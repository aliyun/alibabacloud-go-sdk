// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveEnvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveEnvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveEnvResponse
	GetStatusCode() *int32
	SetBody(v *SaveEnvResponseBody) *SaveEnvResponse
	GetBody() *SaveEnvResponseBody
}

type SaveEnvResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveEnvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveEnvResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveEnvResponse) GoString() string {
	return s.String()
}

func (s *SaveEnvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveEnvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveEnvResponse) GetBody() *SaveEnvResponseBody {
	return s.Body
}

func (s *SaveEnvResponse) SetHeaders(v map[string]*string) *SaveEnvResponse {
	s.Headers = v
	return s
}

func (s *SaveEnvResponse) SetStatusCode(v int32) *SaveEnvResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveEnvResponse) SetBody(v *SaveEnvResponseBody) *SaveEnvResponse {
	s.Body = v
	return s
}

func (s *SaveEnvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
