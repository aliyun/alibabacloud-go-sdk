// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDecodeBlindWatermarkResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDecodeBlindWatermarkResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDecodeBlindWatermarkResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDecodeBlindWatermarkResultResponseBody) *GetDecodeBlindWatermarkResultResponse
	GetBody() *GetDecodeBlindWatermarkResultResponseBody
}

type GetDecodeBlindWatermarkResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDecodeBlindWatermarkResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDecodeBlindWatermarkResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDecodeBlindWatermarkResultResponse) GoString() string {
	return s.String()
}

func (s *GetDecodeBlindWatermarkResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDecodeBlindWatermarkResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDecodeBlindWatermarkResultResponse) GetBody() *GetDecodeBlindWatermarkResultResponseBody {
	return s.Body
}

func (s *GetDecodeBlindWatermarkResultResponse) SetHeaders(v map[string]*string) *GetDecodeBlindWatermarkResultResponse {
	s.Headers = v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponse) SetStatusCode(v int32) *GetDecodeBlindWatermarkResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponse) SetBody(v *GetDecodeBlindWatermarkResultResponseBody) *GetDecodeBlindWatermarkResultResponse {
	s.Body = v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponse) Validate() error {
	return dara.Validate(s)
}
