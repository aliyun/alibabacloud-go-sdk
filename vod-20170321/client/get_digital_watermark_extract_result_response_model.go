// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalWatermarkExtractResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDigitalWatermarkExtractResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDigitalWatermarkExtractResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDigitalWatermarkExtractResultResponseBody) *GetDigitalWatermarkExtractResultResponse
	GetBody() *GetDigitalWatermarkExtractResultResponseBody
}

type GetDigitalWatermarkExtractResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDigitalWatermarkExtractResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDigitalWatermarkExtractResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalWatermarkExtractResultResponse) GoString() string {
	return s.String()
}

func (s *GetDigitalWatermarkExtractResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDigitalWatermarkExtractResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDigitalWatermarkExtractResultResponse) GetBody() *GetDigitalWatermarkExtractResultResponseBody {
	return s.Body
}

func (s *GetDigitalWatermarkExtractResultResponse) SetHeaders(v map[string]*string) *GetDigitalWatermarkExtractResultResponse {
	s.Headers = v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponse) SetStatusCode(v int32) *GetDigitalWatermarkExtractResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponse) SetBody(v *GetDigitalWatermarkExtractResultResponseBody) *GetDigitalWatermarkExtractResultResponse {
	s.Body = v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponse) Validate() error {
	return dara.Validate(s)
}
