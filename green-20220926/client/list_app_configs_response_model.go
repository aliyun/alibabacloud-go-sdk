// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppConfigsResponseBody) *ListAppConfigsResponse
	GetBody() *ListAppConfigsResponseBody
}

type ListAppConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAppConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppConfigsResponse) GetBody() *ListAppConfigsResponseBody {
	return s.Body
}

func (s *ListAppConfigsResponse) SetHeaders(v map[string]*string) *ListAppConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAppConfigsResponse) SetStatusCode(v int32) *ListAppConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppConfigsResponse) SetBody(v *ListAppConfigsResponseBody) *ListAppConfigsResponse {
	s.Body = v
	return s
}

func (s *ListAppConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
