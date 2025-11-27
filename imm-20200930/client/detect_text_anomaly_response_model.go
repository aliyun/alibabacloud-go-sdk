// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectTextAnomalyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectTextAnomalyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectTextAnomalyResponse
	GetStatusCode() *int32
	SetBody(v *DetectTextAnomalyResponseBody) *DetectTextAnomalyResponse
	GetBody() *DetectTextAnomalyResponseBody
}

type DetectTextAnomalyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectTextAnomalyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectTextAnomalyResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectTextAnomalyResponse) GoString() string {
	return s.String()
}

func (s *DetectTextAnomalyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectTextAnomalyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectTextAnomalyResponse) GetBody() *DetectTextAnomalyResponseBody {
	return s.Body
}

func (s *DetectTextAnomalyResponse) SetHeaders(v map[string]*string) *DetectTextAnomalyResponse {
	s.Headers = v
	return s
}

func (s *DetectTextAnomalyResponse) SetStatusCode(v int32) *DetectTextAnomalyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectTextAnomalyResponse) SetBody(v *DetectTextAnomalyResponseBody) *DetectTextAnomalyResponse {
	s.Body = v
	return s
}

func (s *DetectTextAnomalyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
