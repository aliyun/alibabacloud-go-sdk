// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VideoModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VideoModerationResponse
	GetStatusCode() *int32
	SetBody(v *VideoModerationResponseBody) *VideoModerationResponse
	GetBody() *VideoModerationResponseBody
}

type VideoModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResponse) GoString() string {
	return s.String()
}

func (s *VideoModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VideoModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VideoModerationResponse) GetBody() *VideoModerationResponseBody {
	return s.Body
}

func (s *VideoModerationResponse) SetHeaders(v map[string]*string) *VideoModerationResponse {
	s.Headers = v
	return s
}

func (s *VideoModerationResponse) SetStatusCode(v int32) *VideoModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoModerationResponse) SetBody(v *VideoModerationResponseBody) *VideoModerationResponse {
	s.Body = v
	return s
}

func (s *VideoModerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
