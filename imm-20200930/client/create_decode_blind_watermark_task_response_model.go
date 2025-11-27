// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDecodeBlindWatermarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDecodeBlindWatermarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDecodeBlindWatermarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDecodeBlindWatermarkTaskResponseBody) *CreateDecodeBlindWatermarkTaskResponse
	GetBody() *CreateDecodeBlindWatermarkTaskResponseBody
}

type CreateDecodeBlindWatermarkTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDecodeBlindWatermarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDecodeBlindWatermarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDecodeBlindWatermarkTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDecodeBlindWatermarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDecodeBlindWatermarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDecodeBlindWatermarkTaskResponse) GetBody() *CreateDecodeBlindWatermarkTaskResponseBody {
	return s.Body
}

func (s *CreateDecodeBlindWatermarkTaskResponse) SetHeaders(v map[string]*string) *CreateDecodeBlindWatermarkTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskResponse) SetStatusCode(v int32) *CreateDecodeBlindWatermarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskResponse) SetBody(v *CreateDecodeBlindWatermarkTaskResponseBody) *CreateDecodeBlindWatermarkTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
