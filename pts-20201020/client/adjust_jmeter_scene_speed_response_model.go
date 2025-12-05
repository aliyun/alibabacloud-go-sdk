// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustJMeterSceneSpeedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AdjustJMeterSceneSpeedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AdjustJMeterSceneSpeedResponse
	GetStatusCode() *int32
	SetBody(v *AdjustJMeterSceneSpeedResponseBody) *AdjustJMeterSceneSpeedResponse
	GetBody() *AdjustJMeterSceneSpeedResponseBody
}

type AdjustJMeterSceneSpeedResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdjustJMeterSceneSpeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AdjustJMeterSceneSpeedResponse) String() string {
	return dara.Prettify(s)
}

func (s AdjustJMeterSceneSpeedResponse) GoString() string {
	return s.String()
}

func (s *AdjustJMeterSceneSpeedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AdjustJMeterSceneSpeedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AdjustJMeterSceneSpeedResponse) GetBody() *AdjustJMeterSceneSpeedResponseBody {
	return s.Body
}

func (s *AdjustJMeterSceneSpeedResponse) SetHeaders(v map[string]*string) *AdjustJMeterSceneSpeedResponse {
	s.Headers = v
	return s
}

func (s *AdjustJMeterSceneSpeedResponse) SetStatusCode(v int32) *AdjustJMeterSceneSpeedResponse {
	s.StatusCode = &v
	return s
}

func (s *AdjustJMeterSceneSpeedResponse) SetBody(v *AdjustJMeterSceneSpeedResponseBody) *AdjustJMeterSceneSpeedResponse {
	s.Body = v
	return s
}

func (s *AdjustJMeterSceneSpeedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
