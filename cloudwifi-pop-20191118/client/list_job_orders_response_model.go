// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ListJobOrdersResponseBody) *ListJobOrdersResponse
	GetBody() *ListJobOrdersResponseBody
}

type ListJobOrdersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListJobOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobOrdersResponse) GetBody() *ListJobOrdersResponseBody {
	return s.Body
}

func (s *ListJobOrdersResponse) SetHeaders(v map[string]*string) *ListJobOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListJobOrdersResponse) SetStatusCode(v int32) *ListJobOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobOrdersResponse) SetBody(v *ListJobOrdersResponseBody) *ListJobOrdersResponse {
	s.Body = v
	return s
}

func (s *ListJobOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
