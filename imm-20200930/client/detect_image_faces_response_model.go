// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageFacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageFacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageFacesResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageFacesResponseBody) *DetectImageFacesResponse
	GetBody() *DetectImageFacesResponseBody
}

type DetectImageFacesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageFacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageFacesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageFacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageFacesResponse) GetBody() *DetectImageFacesResponseBody {
	return s.Body
}

func (s *DetectImageFacesResponse) SetHeaders(v map[string]*string) *DetectImageFacesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageFacesResponse) SetStatusCode(v int32) *DetectImageFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageFacesResponse) SetBody(v *DetectImageFacesResponseBody) *DetectImageFacesResponse {
	s.Body = v
	return s
}

func (s *DetectImageFacesResponse) Validate() error {
	return dara.Validate(s)
}
