// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoModerationResultResponseBody) *GetVideoModerationResultResponse
	GetBody() *GetVideoModerationResultResponseBody
}

type GetVideoModerationResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoModerationResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoModerationResultResponse) GetBody() *GetVideoModerationResultResponseBody {
	return s.Body
}

func (s *GetVideoModerationResultResponse) SetHeaders(v map[string]*string) *GetVideoModerationResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoModerationResultResponse) SetStatusCode(v int32) *GetVideoModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoModerationResultResponse) SetBody(v *GetVideoModerationResultResponseBody) *GetVideoModerationResultResponse {
	s.Body = v
	return s
}

func (s *GetVideoModerationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
