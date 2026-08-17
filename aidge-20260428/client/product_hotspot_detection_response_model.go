// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductHotspotDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProductHotspotDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProductHotspotDetectionResponse
	GetStatusCode() *int32
	SetBody(v *ProductHotspotDetectionResponseBody) *ProductHotspotDetectionResponse
	GetBody() *ProductHotspotDetectionResponseBody
}

type ProductHotspotDetectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductHotspotDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProductHotspotDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ProductHotspotDetectionResponse) GoString() string {
	return s.String()
}

func (s *ProductHotspotDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProductHotspotDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProductHotspotDetectionResponse) GetBody() *ProductHotspotDetectionResponseBody {
	return s.Body
}

func (s *ProductHotspotDetectionResponse) SetHeaders(v map[string]*string) *ProductHotspotDetectionResponse {
	s.Headers = v
	return s
}

func (s *ProductHotspotDetectionResponse) SetStatusCode(v int32) *ProductHotspotDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ProductHotspotDetectionResponse) SetBody(v *ProductHotspotDetectionResponseBody) *ProductHotspotDetectionResponse {
	s.Body = v
	return s
}

func (s *ProductHotspotDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
