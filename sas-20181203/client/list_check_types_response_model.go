// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckTypesResponseBody) *ListCheckTypesResponse
	GetBody() *ListCheckTypesResponseBody
}

type ListCheckTypesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckTypesResponse) GoString() string {
	return s.String()
}

func (s *ListCheckTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckTypesResponse) GetBody() *ListCheckTypesResponseBody {
	return s.Body
}

func (s *ListCheckTypesResponse) SetHeaders(v map[string]*string) *ListCheckTypesResponse {
	s.Headers = v
	return s
}

func (s *ListCheckTypesResponse) SetStatusCode(v int32) *ListCheckTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckTypesResponse) SetBody(v *ListCheckTypesResponseBody) *ListCheckTypesResponse {
	s.Body = v
	return s
}

func (s *ListCheckTypesResponse) Validate() error {
	return dara.Validate(s)
}
