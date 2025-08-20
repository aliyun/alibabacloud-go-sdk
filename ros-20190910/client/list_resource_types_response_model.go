// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse
	GetBody() *ListResourceTypesResponseBody
}

type ListResourceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceTypesResponse) GetBody() *ListResourceTypesResponseBody {
	return s.Body
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetStatusCode(v int32) *ListResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

func (s *ListResourceTypesResponse) Validate() error {
	return dara.Validate(s)
}
