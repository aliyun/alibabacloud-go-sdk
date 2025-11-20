// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterpolateVideoFrameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InterpolateVideoFrameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InterpolateVideoFrameResponse
	GetStatusCode() *int32
	SetBody(v *InterpolateVideoFrameResponseBody) *InterpolateVideoFrameResponse
	GetBody() *InterpolateVideoFrameResponseBody
}

type InterpolateVideoFrameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InterpolateVideoFrameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InterpolateVideoFrameResponse) String() string {
	return dara.Prettify(s)
}

func (s InterpolateVideoFrameResponse) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InterpolateVideoFrameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InterpolateVideoFrameResponse) GetBody() *InterpolateVideoFrameResponseBody {
	return s.Body
}

func (s *InterpolateVideoFrameResponse) SetHeaders(v map[string]*string) *InterpolateVideoFrameResponse {
	s.Headers = v
	return s
}

func (s *InterpolateVideoFrameResponse) SetStatusCode(v int32) *InterpolateVideoFrameResponse {
	s.StatusCode = &v
	return s
}

func (s *InterpolateVideoFrameResponse) SetBody(v *InterpolateVideoFrameResponseBody) *InterpolateVideoFrameResponse {
	s.Body = v
	return s
}

func (s *InterpolateVideoFrameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
