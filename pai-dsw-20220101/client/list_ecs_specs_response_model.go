// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsSpecsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEcsSpecsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEcsSpecsResponse
	GetStatusCode() *int32
	SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse
	GetBody() *ListEcsSpecsResponseBody
}

type ListEcsSpecsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEcsSpecsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEcsSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEcsSpecsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEcsSpecsResponse) GetBody() *ListEcsSpecsResponseBody {
	return s.Body
}

func (s *ListEcsSpecsResponse) SetHeaders(v map[string]*string) *ListEcsSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListEcsSpecsResponse) SetStatusCode(v int32) *ListEcsSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEcsSpecsResponse) SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse {
	s.Body = v
	return s
}

func (s *ListEcsSpecsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
