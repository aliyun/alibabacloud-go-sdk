// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDigitalWatermarkExtractJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDigitalWatermarkExtractJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDigitalWatermarkExtractJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDigitalWatermarkExtractJobResponseBody) *SubmitDigitalWatermarkExtractJobResponse
	GetBody() *SubmitDigitalWatermarkExtractJobResponseBody
}

type SubmitDigitalWatermarkExtractJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDigitalWatermarkExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDigitalWatermarkExtractJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDigitalWatermarkExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDigitalWatermarkExtractJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDigitalWatermarkExtractJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDigitalWatermarkExtractJobResponse) GetBody() *SubmitDigitalWatermarkExtractJobResponseBody {
	return s.Body
}

func (s *SubmitDigitalWatermarkExtractJobResponse) SetHeaders(v map[string]*string) *SubmitDigitalWatermarkExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobResponse) SetStatusCode(v int32) *SubmitDigitalWatermarkExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobResponse) SetBody(v *SubmitDigitalWatermarkExtractJobResponseBody) *SubmitDigitalWatermarkExtractJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
