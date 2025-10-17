// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectIntlStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeepfakeDetectIntlStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeepfakeDetectIntlStreamResponse
	GetStatusCode() *int32
	SetBody(v *DeepfakeDetectIntlStreamResponseBody) *DeepfakeDetectIntlStreamResponse
	GetBody() *DeepfakeDetectIntlStreamResponseBody
}

type DeepfakeDetectIntlStreamResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeepfakeDetectIntlStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeepfakeDetectIntlStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlStreamResponse) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeepfakeDetectIntlStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeepfakeDetectIntlStreamResponse) GetBody() *DeepfakeDetectIntlStreamResponseBody {
	return s.Body
}

func (s *DeepfakeDetectIntlStreamResponse) SetHeaders(v map[string]*string) *DeepfakeDetectIntlStreamResponse {
	s.Headers = v
	return s
}

func (s *DeepfakeDetectIntlStreamResponse) SetStatusCode(v int32) *DeepfakeDetectIntlStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeepfakeDetectIntlStreamResponse) SetBody(v *DeepfakeDetectIntlStreamResponseBody) *DeepfakeDetectIntlStreamResponse {
	s.Body = v
	return s
}

func (s *DeepfakeDetectIntlStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
