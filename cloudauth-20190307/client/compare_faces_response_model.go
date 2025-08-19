// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareFacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareFacesResponse
	GetStatusCode() *int32
	SetBody(v *CompareFacesResponseBody) *CompareFacesResponse
	GetBody() *CompareFacesResponseBody
}

type CompareFacesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareFacesResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareFacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareFacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareFacesResponse) GetBody() *CompareFacesResponseBody {
	return s.Body
}

func (s *CompareFacesResponse) SetHeaders(v map[string]*string) *CompareFacesResponse {
	s.Headers = v
	return s
}

func (s *CompareFacesResponse) SetStatusCode(v int32) *CompareFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFacesResponse) SetBody(v *CompareFacesResponseBody) *CompareFacesResponse {
	s.Body = v
	return s
}

func (s *CompareFacesResponse) Validate() error {
	return dara.Validate(s)
}
