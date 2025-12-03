// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendByAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendByAliasResponse
	GetStatusCode() *int32
	SetBody(v *SendByAliasResponseBody) *SendByAliasResponse
	GetBody() *SendByAliasResponseBody
}

type SendByAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasResponse) GoString() string {
	return s.String()
}

func (s *SendByAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendByAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendByAliasResponse) GetBody() *SendByAliasResponseBody {
	return s.Body
}

func (s *SendByAliasResponse) SetHeaders(v map[string]*string) *SendByAliasResponse {
	s.Headers = v
	return s
}

func (s *SendByAliasResponse) SetStatusCode(v int32) *SendByAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByAliasResponse) SetBody(v *SendByAliasResponseBody) *SendByAliasResponse {
	s.Body = v
	return s
}

func (s *SendByAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
