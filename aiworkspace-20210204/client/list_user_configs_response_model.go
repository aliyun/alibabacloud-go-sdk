// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserConfigsResponseBody) *ListUserConfigsResponse
	GetBody() *ListUserConfigsResponseBody
}

type ListUserConfigsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListUserConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserConfigsResponse) GetBody() *ListUserConfigsResponseBody {
	return s.Body
}

func (s *ListUserConfigsResponse) SetHeaders(v map[string]*string) *ListUserConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListUserConfigsResponse) SetStatusCode(v int32) *ListUserConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserConfigsResponse) SetBody(v *ListUserConfigsResponseBody) *ListUserConfigsResponse {
	s.Body = v
	return s
}

func (s *ListUserConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
