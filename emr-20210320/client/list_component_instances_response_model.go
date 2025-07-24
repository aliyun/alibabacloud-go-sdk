// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComponentInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComponentInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListComponentInstancesResponseBody) *ListComponentInstancesResponse
	GetBody() *ListComponentInstancesResponseBody
}

type ListComponentInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComponentInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComponentInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComponentInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListComponentInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComponentInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComponentInstancesResponse) GetBody() *ListComponentInstancesResponseBody {
	return s.Body
}

func (s *ListComponentInstancesResponse) SetHeaders(v map[string]*string) *ListComponentInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListComponentInstancesResponse) SetStatusCode(v int32) *ListComponentInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComponentInstancesResponse) SetBody(v *ListComponentInstancesResponseBody) *ListComponentInstancesResponse {
	s.Body = v
	return s
}

func (s *ListComponentInstancesResponse) Validate() error {
	return dara.Validate(s)
}
