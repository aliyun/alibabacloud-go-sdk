// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceConfigsResponseBody) *ListServiceConfigsResponse
	GetBody() *ListServiceConfigsResponseBody
}

type ListServiceConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceConfigsResponse) GetBody() *ListServiceConfigsResponseBody {
	return s.Body
}

func (s *ListServiceConfigsResponse) SetHeaders(v map[string]*string) *ListServiceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceConfigsResponse) SetStatusCode(v int32) *ListServiceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceConfigsResponse) SetBody(v *ListServiceConfigsResponseBody) *ListServiceConfigsResponse {
	s.Body = v
	return s
}

func (s *ListServiceConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
