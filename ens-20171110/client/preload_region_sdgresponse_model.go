// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadRegionSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreloadRegionSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreloadRegionSDGResponse
	GetStatusCode() *int32
	SetBody(v *PreloadRegionSDGResponseBody) *PreloadRegionSDGResponse
	GetBody() *PreloadRegionSDGResponseBody
}

type PreloadRegionSDGResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadRegionSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadRegionSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s PreloadRegionSDGResponse) GoString() string {
	return s.String()
}

func (s *PreloadRegionSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreloadRegionSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreloadRegionSDGResponse) GetBody() *PreloadRegionSDGResponseBody {
	return s.Body
}

func (s *PreloadRegionSDGResponse) SetHeaders(v map[string]*string) *PreloadRegionSDGResponse {
	s.Headers = v
	return s
}

func (s *PreloadRegionSDGResponse) SetStatusCode(v int32) *PreloadRegionSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadRegionSDGResponse) SetBody(v *PreloadRegionSDGResponseBody) *PreloadRegionSDGResponse {
	s.Body = v
	return s
}

func (s *PreloadRegionSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
