// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasFileIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendByAliasFileIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendByAliasFileIdResponse
	GetStatusCode() *int32
	SetBody(v *SendByAliasFileIdResponseBody) *SendByAliasFileIdResponse
	GetBody() *SendByAliasFileIdResponseBody
}

type SendByAliasFileIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByAliasFileIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByAliasFileIdResponse) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasFileIdResponse) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendByAliasFileIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendByAliasFileIdResponse) GetBody() *SendByAliasFileIdResponseBody {
	return s.Body
}

func (s *SendByAliasFileIdResponse) SetHeaders(v map[string]*string) *SendByAliasFileIdResponse {
	s.Headers = v
	return s
}

func (s *SendByAliasFileIdResponse) SetStatusCode(v int32) *SendByAliasFileIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByAliasFileIdResponse) SetBody(v *SendByAliasFileIdResponseBody) *SendByAliasFileIdResponse {
	s.Body = v
	return s
}

func (s *SendByAliasFileIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
