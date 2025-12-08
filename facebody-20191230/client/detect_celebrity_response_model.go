// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectCelebrityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectCelebrityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectCelebrityResponse
	GetStatusCode() *int32
	SetBody(v *DetectCelebrityResponseBody) *DetectCelebrityResponse
	GetBody() *DetectCelebrityResponseBody
}

type DetectCelebrityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectCelebrityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectCelebrityResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectCelebrityResponse) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectCelebrityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectCelebrityResponse) GetBody() *DetectCelebrityResponseBody {
	return s.Body
}

func (s *DetectCelebrityResponse) SetHeaders(v map[string]*string) *DetectCelebrityResponse {
	s.Headers = v
	return s
}

func (s *DetectCelebrityResponse) SetStatusCode(v int32) *DetectCelebrityResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectCelebrityResponse) SetBody(v *DetectCelebrityResponseBody) *DetectCelebrityResponse {
	s.Body = v
	return s
}

func (s *DetectCelebrityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
