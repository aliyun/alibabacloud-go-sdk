// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectBodyCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectBodyCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectBodyCountResponse
	GetStatusCode() *int32
	SetBody(v *DetectBodyCountResponseBody) *DetectBodyCountResponse
	GetBody() *DetectBodyCountResponseBody
}

type DetectBodyCountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectBodyCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectBodyCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectBodyCountResponse) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectBodyCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectBodyCountResponse) GetBody() *DetectBodyCountResponseBody {
	return s.Body
}

func (s *DetectBodyCountResponse) SetHeaders(v map[string]*string) *DetectBodyCountResponse {
	s.Headers = v
	return s
}

func (s *DetectBodyCountResponse) SetStatusCode(v int32) *DetectBodyCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectBodyCountResponse) SetBody(v *DetectBodyCountResponseBody) *DetectBodyCountResponse {
	s.Body = v
	return s
}

func (s *DetectBodyCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
