// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRedisValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRedisValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRedisValueResponse
	GetStatusCode() *int32
	SetBody(v *SetRedisValueResponseBody) *SetRedisValueResponse
	GetBody() *SetRedisValueResponseBody
}

type SetRedisValueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRedisValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRedisValueResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRedisValueResponse) GoString() string {
	return s.String()
}

func (s *SetRedisValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRedisValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRedisValueResponse) GetBody() *SetRedisValueResponseBody {
	return s.Body
}

func (s *SetRedisValueResponse) SetHeaders(v map[string]*string) *SetRedisValueResponse {
	s.Headers = v
	return s
}

func (s *SetRedisValueResponse) SetStatusCode(v int32) *SetRedisValueResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRedisValueResponse) SetBody(v *SetRedisValueResponseBody) *SetRedisValueResponse {
	s.Body = v
	return s
}

func (s *SetRedisValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
