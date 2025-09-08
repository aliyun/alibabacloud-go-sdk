// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectLogStoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectLogStoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectLogStoresResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectLogStoresResponseBody) *ListProjectLogStoresResponse
	GetBody() *ListProjectLogStoresResponseBody
}

type ListProjectLogStoresResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectLogStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectLogStoresResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLogStoresResponse) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectLogStoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectLogStoresResponse) GetBody() *ListProjectLogStoresResponseBody {
	return s.Body
}

func (s *ListProjectLogStoresResponse) SetHeaders(v map[string]*string) *ListProjectLogStoresResponse {
	s.Headers = v
	return s
}

func (s *ListProjectLogStoresResponse) SetStatusCode(v int32) *ListProjectLogStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectLogStoresResponse) SetBody(v *ListProjectLogStoresResponseBody) *ListProjectLogStoresResponse {
	s.Body = v
	return s
}

func (s *ListProjectLogStoresResponse) Validate() error {
	return dara.Validate(s)
}
