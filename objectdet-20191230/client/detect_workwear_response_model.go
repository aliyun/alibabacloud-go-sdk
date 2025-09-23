// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectWorkwearResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectWorkwearResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectWorkwearResponse
	GetStatusCode() *int32
	SetBody(v *DetectWorkwearResponseBody) *DetectWorkwearResponse
	GetBody() *DetectWorkwearResponseBody
}

type DetectWorkwearResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectWorkwearResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectWorkwearResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearResponse) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectWorkwearResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectWorkwearResponse) GetBody() *DetectWorkwearResponseBody {
	return s.Body
}

func (s *DetectWorkwearResponse) SetHeaders(v map[string]*string) *DetectWorkwearResponse {
	s.Headers = v
	return s
}

func (s *DetectWorkwearResponse) SetStatusCode(v int32) *DetectWorkwearResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectWorkwearResponse) SetBody(v *DetectWorkwearResponseBody) *DetectWorkwearResponse {
	s.Body = v
	return s
}

func (s *DetectWorkwearResponse) Validate() error {
	return dara.Validate(s)
}
