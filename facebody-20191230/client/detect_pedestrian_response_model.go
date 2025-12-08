// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectPedestrianResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectPedestrianResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectPedestrianResponse
	GetStatusCode() *int32
	SetBody(v *DetectPedestrianResponseBody) *DetectPedestrianResponse
	GetBody() *DetectPedestrianResponseBody
}

type DetectPedestrianResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectPedestrianResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectPedestrianResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectPedestrianResponse) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectPedestrianResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectPedestrianResponse) GetBody() *DetectPedestrianResponseBody {
	return s.Body
}

func (s *DetectPedestrianResponse) SetHeaders(v map[string]*string) *DetectPedestrianResponse {
	s.Headers = v
	return s
}

func (s *DetectPedestrianResponse) SetStatusCode(v int32) *DetectPedestrianResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectPedestrianResponse) SetBody(v *DetectPedestrianResponseBody) *DetectPedestrianResponse {
	s.Body = v
	return s
}

func (s *DetectPedestrianResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
