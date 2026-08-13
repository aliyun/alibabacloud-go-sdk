// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableConfigsResponseBody) *ListAvailableConfigsResponse
	GetBody() *ListAvailableConfigsResponseBody
}

type ListAvailableConfigsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableConfigsResponse) GetBody() *ListAvailableConfigsResponseBody {
	return s.Body
}

func (s *ListAvailableConfigsResponse) SetHeaders(v map[string]*string) *ListAvailableConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableConfigsResponse) SetStatusCode(v int32) *ListAvailableConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableConfigsResponse) SetBody(v *ListAvailableConfigsResponseBody) *ListAvailableConfigsResponse {
	s.Body = v
	return s
}

func (s *ListAvailableConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
