// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceSpecsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceSpecsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceSpecsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceSpecsResponseBody) *ListInstanceSpecsResponse
	GetBody() *ListInstanceSpecsResponseBody
}

type ListInstanceSpecsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceSpecsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceSpecsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceSpecsResponse) GetBody() *ListInstanceSpecsResponseBody {
	return s.Body
}

func (s *ListInstanceSpecsResponse) SetHeaders(v map[string]*string) *ListInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceSpecsResponse) SetStatusCode(v int32) *ListInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceSpecsResponse) SetBody(v *ListInstanceSpecsResponseBody) *ListInstanceSpecsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceSpecsResponse) Validate() error {
	return dara.Validate(s)
}
