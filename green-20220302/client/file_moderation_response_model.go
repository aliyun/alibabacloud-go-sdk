// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FileModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FileModerationResponse
	GetStatusCode() *int32
	SetBody(v *FileModerationResponseBody) *FileModerationResponse
	GetBody() *FileModerationResponseBody
}

type FileModerationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FileModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s FileModerationResponse) GoString() string {
	return s.String()
}

func (s *FileModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FileModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FileModerationResponse) GetBody() *FileModerationResponseBody {
	return s.Body
}

func (s *FileModerationResponse) SetHeaders(v map[string]*string) *FileModerationResponse {
	s.Headers = v
	return s
}

func (s *FileModerationResponse) SetStatusCode(v int32) *FileModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *FileModerationResponse) SetBody(v *FileModerationResponseBody) *FileModerationResponse {
	s.Body = v
	return s
}

func (s *FileModerationResponse) Validate() error {
	return dara.Validate(s)
}
