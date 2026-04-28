// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoDRMLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VideoDRMLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VideoDRMLicenseResponse
	GetStatusCode() *int32
	SetBody(v *VideoDRMLicenseResponseBody) *VideoDRMLicenseResponse
	GetBody() *VideoDRMLicenseResponseBody
}

type VideoDRMLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoDRMLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoDRMLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s VideoDRMLicenseResponse) GoString() string {
	return s.String()
}

func (s *VideoDRMLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VideoDRMLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VideoDRMLicenseResponse) GetBody() *VideoDRMLicenseResponseBody {
	return s.Body
}

func (s *VideoDRMLicenseResponse) SetHeaders(v map[string]*string) *VideoDRMLicenseResponse {
	s.Headers = v
	return s
}

func (s *VideoDRMLicenseResponse) SetStatusCode(v int32) *VideoDRMLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoDRMLicenseResponse) SetBody(v *VideoDRMLicenseResponseBody) *VideoDRMLicenseResponse {
	s.Body = v
	return s
}

func (s *VideoDRMLicenseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
