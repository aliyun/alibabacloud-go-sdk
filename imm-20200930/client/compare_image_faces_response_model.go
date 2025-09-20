// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareImageFacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareImageFacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareImageFacesResponse
	GetStatusCode() *int32
	SetBody(v *CompareImageFacesResponseBody) *CompareImageFacesResponse
	GetBody() *CompareImageFacesResponseBody
}

type CompareImageFacesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareImageFacesResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareImageFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareImageFacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareImageFacesResponse) GetBody() *CompareImageFacesResponseBody {
	return s.Body
}

func (s *CompareImageFacesResponse) SetHeaders(v map[string]*string) *CompareImageFacesResponse {
	s.Headers = v
	return s
}

func (s *CompareImageFacesResponse) SetStatusCode(v int32) *CompareImageFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareImageFacesResponse) SetBody(v *CompareImageFacesResponseBody) *CompareImageFacesResponse {
	s.Body = v
	return s
}

func (s *CompareImageFacesResponse) Validate() error {
	return dara.Validate(s)
}
