// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectWhiteBaseImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectWhiteBaseImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectWhiteBaseImageResponse
	GetStatusCode() *int32
	SetBody(v *DetectWhiteBaseImageResponseBody) *DetectWhiteBaseImageResponse
	GetBody() *DetectWhiteBaseImageResponseBody
}

type DetectWhiteBaseImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectWhiteBaseImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectWhiteBaseImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectWhiteBaseImageResponse) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectWhiteBaseImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectWhiteBaseImageResponse) GetBody() *DetectWhiteBaseImageResponseBody {
	return s.Body
}

func (s *DetectWhiteBaseImageResponse) SetHeaders(v map[string]*string) *DetectWhiteBaseImageResponse {
	s.Headers = v
	return s
}

func (s *DetectWhiteBaseImageResponse) SetStatusCode(v int32) *DetectWhiteBaseImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectWhiteBaseImageResponse) SetBody(v *DetectWhiteBaseImageResponseBody) *DetectWhiteBaseImageResponse {
	s.Body = v
	return s
}

func (s *DetectWhiteBaseImageResponse) Validate() error {
	return dara.Validate(s)
}
