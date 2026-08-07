// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemConfigsResponseBody) *ListSystemConfigsResponse
	GetBody() *ListSystemConfigsResponseBody
}

type ListSystemConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListSystemConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemConfigsResponse) GetBody() *ListSystemConfigsResponseBody {
	return s.Body
}

func (s *ListSystemConfigsResponse) SetHeaders(v map[string]*string) *ListSystemConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListSystemConfigsResponse) SetStatusCode(v int32) *ListSystemConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemConfigsResponse) SetBody(v *ListSystemConfigsResponseBody) *ListSystemConfigsResponse {
	s.Body = v
	return s
}

func (s *ListSystemConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
