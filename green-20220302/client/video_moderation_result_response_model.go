// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VideoModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VideoModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *VideoModerationResultResponseBody) *VideoModerationResultResponse
	GetBody() *VideoModerationResultResponseBody
}

type VideoModerationResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponse) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VideoModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VideoModerationResultResponse) GetBody() *VideoModerationResultResponseBody {
	return s.Body
}

func (s *VideoModerationResultResponse) SetHeaders(v map[string]*string) *VideoModerationResultResponse {
	s.Headers = v
	return s
}

func (s *VideoModerationResultResponse) SetStatusCode(v int32) *VideoModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoModerationResultResponse) SetBody(v *VideoModerationResultResponseBody) *VideoModerationResultResponse {
	s.Body = v
	return s
}

func (s *VideoModerationResultResponse) Validate() error {
	return dara.Validate(s)
}
