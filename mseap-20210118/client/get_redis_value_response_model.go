// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedisValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRedisValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRedisValueResponse
	GetStatusCode() *int32
	SetBody(v *GetRedisValueResponseBody) *GetRedisValueResponse
	GetBody() *GetRedisValueResponseBody
}

type GetRedisValueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRedisValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRedisValueResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRedisValueResponse) GoString() string {
	return s.String()
}

func (s *GetRedisValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRedisValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRedisValueResponse) GetBody() *GetRedisValueResponseBody {
	return s.Body
}

func (s *GetRedisValueResponse) SetHeaders(v map[string]*string) *GetRedisValueResponse {
	s.Headers = v
	return s
}

func (s *GetRedisValueResponse) SetStatusCode(v int32) *GetRedisValueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRedisValueResponse) SetBody(v *GetRedisValueResponseBody) *GetRedisValueResponse {
	s.Body = v
	return s
}

func (s *GetRedisValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
