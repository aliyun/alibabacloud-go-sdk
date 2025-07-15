// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListParametersResponse
	GetStatusCode() *int32
	SetBody(v *ListParametersResponseBody) *ListParametersResponse
	GetBody() *ListParametersResponseBody
}

type ListParametersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListParametersResponse) GoString() string {
	return s.String()
}

func (s *ListParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListParametersResponse) GetBody() *ListParametersResponseBody {
	return s.Body
}

func (s *ListParametersResponse) SetHeaders(v map[string]*string) *ListParametersResponse {
	s.Headers = v
	return s
}

func (s *ListParametersResponse) SetStatusCode(v int32) *ListParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParametersResponse) SetBody(v *ListParametersResponseBody) *ListParametersResponse {
	s.Body = v
	return s
}

func (s *ListParametersResponse) Validate() error {
	return dara.Validate(s)
}
