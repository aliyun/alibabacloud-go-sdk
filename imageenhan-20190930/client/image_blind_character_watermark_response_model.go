// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBlindCharacterWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageBlindCharacterWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageBlindCharacterWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *ImageBlindCharacterWatermarkResponseBody) *ImageBlindCharacterWatermarkResponse
	GetBody() *ImageBlindCharacterWatermarkResponseBody
}

type ImageBlindCharacterWatermarkResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageBlindCharacterWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageBlindCharacterWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponse) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageBlindCharacterWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageBlindCharacterWatermarkResponse) GetBody() *ImageBlindCharacterWatermarkResponseBody {
	return s.Body
}

func (s *ImageBlindCharacterWatermarkResponse) SetHeaders(v map[string]*string) *ImageBlindCharacterWatermarkResponse {
	s.Headers = v
	return s
}

func (s *ImageBlindCharacterWatermarkResponse) SetStatusCode(v int32) *ImageBlindCharacterWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponse) SetBody(v *ImageBlindCharacterWatermarkResponseBody) *ImageBlindCharacterWatermarkResponse {
	s.Body = v
	return s
}

func (s *ImageBlindCharacterWatermarkResponse) Validate() error {
	return dara.Validate(s)
}
