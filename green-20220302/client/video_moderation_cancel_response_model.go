// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VideoModerationCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VideoModerationCancelResponse
	GetStatusCode() *int32
	SetBody(v *VideoModerationCancelResponseBody) *VideoModerationCancelResponse
	GetBody() *VideoModerationCancelResponseBody
}

type VideoModerationCancelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoModerationCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoModerationCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationCancelResponse) GoString() string {
	return s.String()
}

func (s *VideoModerationCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VideoModerationCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VideoModerationCancelResponse) GetBody() *VideoModerationCancelResponseBody {
	return s.Body
}

func (s *VideoModerationCancelResponse) SetHeaders(v map[string]*string) *VideoModerationCancelResponse {
	s.Headers = v
	return s
}

func (s *VideoModerationCancelResponse) SetStatusCode(v int32) *VideoModerationCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoModerationCancelResponse) SetBody(v *VideoModerationCancelResponseBody) *VideoModerationCancelResponse {
	s.Body = v
	return s
}

func (s *VideoModerationCancelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
