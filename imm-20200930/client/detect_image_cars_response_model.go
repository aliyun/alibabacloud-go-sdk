// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCarsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageCarsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageCarsResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageCarsResponseBody) *DetectImageCarsResponse
	GetBody() *DetectImageCarsResponseBody
}

type DetectImageCarsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageCarsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageCarsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCarsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageCarsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageCarsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageCarsResponse) GetBody() *DetectImageCarsResponseBody {
	return s.Body
}

func (s *DetectImageCarsResponse) SetHeaders(v map[string]*string) *DetectImageCarsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageCarsResponse) SetStatusCode(v int32) *DetectImageCarsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageCarsResponse) SetBody(v *DetectImageCarsResponseBody) *DetectImageCarsResponse {
	s.Body = v
	return s
}

func (s *DetectImageCarsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
