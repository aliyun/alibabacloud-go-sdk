// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserAuthConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserAuthConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserAuthConfigsResponseBody) *ListUserAuthConfigsResponse
	GetBody() *ListUserAuthConfigsResponseBody
}

type ListUserAuthConfigsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAuthConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserAuthConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListUserAuthConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserAuthConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserAuthConfigsResponse) GetBody() *ListUserAuthConfigsResponseBody {
	return s.Body
}

func (s *ListUserAuthConfigsResponse) SetHeaders(v map[string]*string) *ListUserAuthConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListUserAuthConfigsResponse) SetStatusCode(v int32) *ListUserAuthConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAuthConfigsResponse) SetBody(v *ListUserAuthConfigsResponseBody) *ListUserAuthConfigsResponse {
	s.Body = v
	return s
}

func (s *ListUserAuthConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
