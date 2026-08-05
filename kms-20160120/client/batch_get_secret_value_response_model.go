// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetSecretValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetSecretValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetSecretValueResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetSecretValueResponseBody) *BatchGetSecretValueResponse
	GetBody() *BatchGetSecretValueResponseBody
}

type BatchGetSecretValueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetSecretValueResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueResponse) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetSecretValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetSecretValueResponse) GetBody() *BatchGetSecretValueResponseBody {
	return s.Body
}

func (s *BatchGetSecretValueResponse) SetHeaders(v map[string]*string) *BatchGetSecretValueResponse {
	s.Headers = v
	return s
}

func (s *BatchGetSecretValueResponse) SetStatusCode(v int32) *BatchGetSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetSecretValueResponse) SetBody(v *BatchGetSecretValueResponseBody) *BatchGetSecretValueResponse {
	s.Body = v
	return s
}

func (s *BatchGetSecretValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
