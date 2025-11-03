// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpsecServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpsecServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpsecServersResponse
	GetStatusCode() *int32
	SetBody(v *ListIpsecServersResponseBody) *ListIpsecServersResponse
	GetBody() *ListIpsecServersResponseBody
}

type ListIpsecServersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpsecServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpsecServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServersResponse) GoString() string {
	return s.String()
}

func (s *ListIpsecServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpsecServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpsecServersResponse) GetBody() *ListIpsecServersResponseBody {
	return s.Body
}

func (s *ListIpsecServersResponse) SetHeaders(v map[string]*string) *ListIpsecServersResponse {
	s.Headers = v
	return s
}

func (s *ListIpsecServersResponse) SetStatusCode(v int32) *ListIpsecServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpsecServersResponse) SetBody(v *ListIpsecServersResponseBody) *ListIpsecServersResponse {
	s.Body = v
	return s
}

func (s *ListIpsecServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
