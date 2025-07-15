// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnassignedNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnassignedNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnassignedNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ListUnassignedNumbersResponseBody) *ListUnassignedNumbersResponse
	GetBody() *ListUnassignedNumbersResponseBody
}

type ListUnassignedNumbersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnassignedNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnassignedNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnassignedNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnassignedNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnassignedNumbersResponse) GetBody() *ListUnassignedNumbersResponseBody {
	return s.Body
}

func (s *ListUnassignedNumbersResponse) SetHeaders(v map[string]*string) *ListUnassignedNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListUnassignedNumbersResponse) SetStatusCode(v int32) *ListUnassignedNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnassignedNumbersResponse) SetBody(v *ListUnassignedNumbersResponseBody) *ListUnassignedNumbersResponse {
	s.Body = v
	return s
}

func (s *ListUnassignedNumbersResponse) Validate() error {
	return dara.Validate(s)
}
