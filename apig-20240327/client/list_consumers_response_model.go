// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumersResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumersResponseBody) *ListConsumersResponse
	GetBody() *ListConsumersResponseBody
}

type ListConsumersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumersResponse) GoString() string {
	return s.String()
}

func (s *ListConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumersResponse) GetBody() *ListConsumersResponseBody {
	return s.Body
}

func (s *ListConsumersResponse) SetHeaders(v map[string]*string) *ListConsumersResponse {
	s.Headers = v
	return s
}

func (s *ListConsumersResponse) SetStatusCode(v int32) *ListConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumersResponse) SetBody(v *ListConsumersResponseBody) *ListConsumersResponse {
	s.Body = v
	return s
}

func (s *ListConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
