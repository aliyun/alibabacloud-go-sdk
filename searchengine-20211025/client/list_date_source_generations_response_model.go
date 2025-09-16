// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDateSourceGenerationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDateSourceGenerationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDateSourceGenerationsResponse
	GetStatusCode() *int32
	SetBody(v *ListDateSourceGenerationsResponseBody) *ListDateSourceGenerationsResponse
	GetBody() *ListDateSourceGenerationsResponseBody
}

type ListDateSourceGenerationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDateSourceGenerationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDateSourceGenerationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDateSourceGenerationsResponse) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDateSourceGenerationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDateSourceGenerationsResponse) GetBody() *ListDateSourceGenerationsResponseBody {
	return s.Body
}

func (s *ListDateSourceGenerationsResponse) SetHeaders(v map[string]*string) *ListDateSourceGenerationsResponse {
	s.Headers = v
	return s
}

func (s *ListDateSourceGenerationsResponse) SetStatusCode(v int32) *ListDateSourceGenerationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDateSourceGenerationsResponse) SetBody(v *ListDateSourceGenerationsResponseBody) *ListDateSourceGenerationsResponse {
	s.Body = v
	return s
}

func (s *ListDateSourceGenerationsResponse) Validate() error {
	return dara.Validate(s)
}
