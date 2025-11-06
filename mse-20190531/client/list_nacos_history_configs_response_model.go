// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosHistoryConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNacosHistoryConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNacosHistoryConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListNacosHistoryConfigsResponseBody) *ListNacosHistoryConfigsResponse
	GetBody() *ListNacosHistoryConfigsResponseBody
}

type ListNacosHistoryConfigsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNacosHistoryConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNacosHistoryConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNacosHistoryConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNacosHistoryConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNacosHistoryConfigsResponse) GetBody() *ListNacosHistoryConfigsResponseBody {
	return s.Body
}

func (s *ListNacosHistoryConfigsResponse) SetHeaders(v map[string]*string) *ListNacosHistoryConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListNacosHistoryConfigsResponse) SetStatusCode(v int32) *ListNacosHistoryConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNacosHistoryConfigsResponse) SetBody(v *ListNacosHistoryConfigsResponseBody) *ListNacosHistoryConfigsResponse {
	s.Body = v
	return s
}

func (s *ListNacosHistoryConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
