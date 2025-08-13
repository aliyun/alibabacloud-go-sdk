// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportResourceTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupportResourceTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupportResourceTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListSupportResourceTypesResponseBody) *ListSupportResourceTypesResponse
	GetBody() *ListSupportResourceTypesResponseBody
}

type ListSupportResourceTypesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportResourceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupportResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupportResourceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupportResourceTypesResponse) GetBody() *ListSupportResourceTypesResponseBody {
	return s.Body
}

func (s *ListSupportResourceTypesResponse) SetHeaders(v map[string]*string) *ListSupportResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListSupportResourceTypesResponse) SetStatusCode(v int32) *ListSupportResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportResourceTypesResponse) SetBody(v *ListSupportResourceTypesResponseBody) *ListSupportResourceTypesResponse {
	s.Body = v
	return s
}

func (s *ListSupportResourceTypesResponse) Validate() error {
	return dara.Validate(s)
}
