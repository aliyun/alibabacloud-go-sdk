// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAliasResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAliasResponseBody) *UnbindAliasResponse
	GetBody() *UnbindAliasResponseBody
}

type UnbindAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAliasResponse) GoString() string {
	return s.String()
}

func (s *UnbindAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAliasResponse) GetBody() *UnbindAliasResponseBody {
	return s.Body
}

func (s *UnbindAliasResponse) SetHeaders(v map[string]*string) *UnbindAliasResponse {
	s.Headers = v
	return s
}

func (s *UnbindAliasResponse) SetStatusCode(v int32) *UnbindAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAliasResponse) SetBody(v *UnbindAliasResponseBody) *UnbindAliasResponse {
	s.Body = v
	return s
}

func (s *UnbindAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
