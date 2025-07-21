// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutSecretValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutSecretValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutSecretValueResponse
	GetStatusCode() *int32
	SetBody(v *PutSecretValueResponseBody) *PutSecretValueResponse
	GetBody() *PutSecretValueResponseBody
}

type PutSecretValueResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutSecretValueResponse) String() string {
	return dara.Prettify(s)
}

func (s PutSecretValueResponse) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutSecretValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutSecretValueResponse) GetBody() *PutSecretValueResponseBody {
	return s.Body
}

func (s *PutSecretValueResponse) SetHeaders(v map[string]*string) *PutSecretValueResponse {
	s.Headers = v
	return s
}

func (s *PutSecretValueResponse) SetStatusCode(v int32) *PutSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *PutSecretValueResponse) SetBody(v *PutSecretValueResponseBody) *PutSecretValueResponse {
	s.Body = v
	return s
}

func (s *PutSecretValueResponse) Validate() error {
	return dara.Validate(s)
}
