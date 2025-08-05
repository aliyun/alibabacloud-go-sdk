// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListPackagesResponseBody) *ListPackagesResponse
	GetBody() *ListPackagesResponseBody
}

type ListPackagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPackagesResponse) GetBody() *ListPackagesResponseBody {
	return s.Body
}

func (s *ListPackagesResponse) SetHeaders(v map[string]*string) *ListPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListPackagesResponse) SetStatusCode(v int32) *ListPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPackagesResponse) SetBody(v *ListPackagesResponseBody) *ListPackagesResponse {
	s.Body = v
	return s
}

func (s *ListPackagesResponse) Validate() error {
	return dara.Validate(s)
}
