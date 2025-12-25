// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPredictionWallLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PredictionWallLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PredictionWallLineResponse
	GetStatusCode() *int32
	SetBody(v *PredictionWallLineResponseBody) *PredictionWallLineResponse
	GetBody() *PredictionWallLineResponseBody
}

type PredictionWallLineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PredictionWallLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PredictionWallLineResponse) String() string {
	return dara.Prettify(s)
}

func (s PredictionWallLineResponse) GoString() string {
	return s.String()
}

func (s *PredictionWallLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PredictionWallLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PredictionWallLineResponse) GetBody() *PredictionWallLineResponseBody {
	return s.Body
}

func (s *PredictionWallLineResponse) SetHeaders(v map[string]*string) *PredictionWallLineResponse {
	s.Headers = v
	return s
}

func (s *PredictionWallLineResponse) SetStatusCode(v int32) *PredictionWallLineResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictionWallLineResponse) SetBody(v *PredictionWallLineResponseBody) *PredictionWallLineResponse {
	s.Body = v
	return s
}

func (s *PredictionWallLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
