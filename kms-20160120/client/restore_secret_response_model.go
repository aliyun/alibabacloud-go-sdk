// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreSecretResponse
	GetStatusCode() *int32
	SetBody(v *RestoreSecretResponseBody) *RestoreSecretResponse
	GetBody() *RestoreSecretResponseBody
}

type RestoreSecretResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreSecretResponse) GoString() string {
	return s.String()
}

func (s *RestoreSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreSecretResponse) GetBody() *RestoreSecretResponseBody {
	return s.Body
}

func (s *RestoreSecretResponse) SetHeaders(v map[string]*string) *RestoreSecretResponse {
	s.Headers = v
	return s
}

func (s *RestoreSecretResponse) SetStatusCode(v int32) *RestoreSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreSecretResponse) SetBody(v *RestoreSecretResponseBody) *RestoreSecretResponse {
	s.Body = v
	return s
}

func (s *RestoreSecretResponse) Validate() error {
	return dara.Validate(s)
}
