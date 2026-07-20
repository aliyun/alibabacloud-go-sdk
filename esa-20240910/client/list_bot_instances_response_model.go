// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBotInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBotInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBotInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListBotInstancesResponseBody) *ListBotInstancesResponse
	GetBody() *ListBotInstancesResponseBody
}

type ListBotInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBotInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBotInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBotInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListBotInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBotInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBotInstancesResponse) GetBody() *ListBotInstancesResponseBody {
	return s.Body
}

func (s *ListBotInstancesResponse) SetHeaders(v map[string]*string) *ListBotInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListBotInstancesResponse) SetStatusCode(v int32) *ListBotInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBotInstancesResponse) SetBody(v *ListBotInstancesResponseBody) *ListBotInstancesResponse {
	s.Body = v
	return s
}

func (s *ListBotInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
