// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectAigcImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectAigcImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectAigcImageResponse
	GetStatusCode() *int32
	SetBody(v *DetectAigcImageResponseBody) *DetectAigcImageResponse
	GetBody() *DetectAigcImageResponseBody
}

type DetectAigcImageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectAigcImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectAigcImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectAigcImageResponse) GoString() string {
	return s.String()
}

func (s *DetectAigcImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectAigcImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectAigcImageResponse) GetBody() *DetectAigcImageResponseBody {
	return s.Body
}

func (s *DetectAigcImageResponse) SetHeaders(v map[string]*string) *DetectAigcImageResponse {
	s.Headers = v
	return s
}

func (s *DetectAigcImageResponse) SetStatusCode(v int32) *DetectAigcImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectAigcImageResponse) SetBody(v *DetectAigcImageResponseBody) *DetectAigcImageResponse {
	s.Body = v
	return s
}

func (s *DetectAigcImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
