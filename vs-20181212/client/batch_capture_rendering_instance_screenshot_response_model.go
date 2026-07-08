// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCaptureRenderingInstanceScreenshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCaptureRenderingInstanceScreenshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCaptureRenderingInstanceScreenshotResponse
	GetStatusCode() *int32
	SetBody(v *BatchCaptureRenderingInstanceScreenshotResponseBody) *BatchCaptureRenderingInstanceScreenshotResponse
	GetBody() *BatchCaptureRenderingInstanceScreenshotResponseBody
}

type BatchCaptureRenderingInstanceScreenshotResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCaptureRenderingInstanceScreenshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCaptureRenderingInstanceScreenshotResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotResponse) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCaptureRenderingInstanceScreenshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCaptureRenderingInstanceScreenshotResponse) GetBody() *BatchCaptureRenderingInstanceScreenshotResponseBody {
	return s.Body
}

func (s *BatchCaptureRenderingInstanceScreenshotResponse) SetHeaders(v map[string]*string) *BatchCaptureRenderingInstanceScreenshotResponse {
	s.Headers = v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponse) SetStatusCode(v int32) *BatchCaptureRenderingInstanceScreenshotResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponse) SetBody(v *BatchCaptureRenderingInstanceScreenshotResponseBody) *BatchCaptureRenderingInstanceScreenshotResponse {
	s.Body = v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
