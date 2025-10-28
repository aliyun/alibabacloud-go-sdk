// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageQueueModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageQueueModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageQueueModerationResponse
	GetStatusCode() *int32
	SetBody(v *ImageQueueModerationResponseBody) *ImageQueueModerationResponse
	GetBody() *ImageQueueModerationResponseBody
}

type ImageQueueModerationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageQueueModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageQueueModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageQueueModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageQueueModerationResponse) GetBody() *ImageQueueModerationResponseBody {
	return s.Body
}

func (s *ImageQueueModerationResponse) SetHeaders(v map[string]*string) *ImageQueueModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageQueueModerationResponse) SetStatusCode(v int32) *ImageQueueModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageQueueModerationResponse) SetBody(v *ImageQueueModerationResponseBody) *ImageQueueModerationResponse {
	s.Body = v
	return s
}

func (s *ImageQueueModerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
