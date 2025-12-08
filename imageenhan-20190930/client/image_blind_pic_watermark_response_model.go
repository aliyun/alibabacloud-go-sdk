// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBlindPicWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageBlindPicWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageBlindPicWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *ImageBlindPicWatermarkResponseBody) *ImageBlindPicWatermarkResponse
	GetBody() *ImageBlindPicWatermarkResponseBody
}

type ImageBlindPicWatermarkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageBlindPicWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageBlindPicWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindPicWatermarkResponse) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageBlindPicWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageBlindPicWatermarkResponse) GetBody() *ImageBlindPicWatermarkResponseBody {
	return s.Body
}

func (s *ImageBlindPicWatermarkResponse) SetHeaders(v map[string]*string) *ImageBlindPicWatermarkResponse {
	s.Headers = v
	return s
}

func (s *ImageBlindPicWatermarkResponse) SetStatusCode(v int32) *ImageBlindPicWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageBlindPicWatermarkResponse) SetBody(v *ImageBlindPicWatermarkResponseBody) *ImageBlindPicWatermarkResponse {
	s.Body = v
	return s
}

func (s *ImageBlindPicWatermarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
