// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoModerationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoModerationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoModerationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoModerationTaskResponseBody) *CreateVideoModerationTaskResponse
	GetBody() *CreateVideoModerationTaskResponseBody
}

type CreateVideoModerationTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoModerationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoModerationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoModerationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoModerationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoModerationTaskResponse) GetBody() *CreateVideoModerationTaskResponseBody {
	return s.Body
}

func (s *CreateVideoModerationTaskResponse) SetHeaders(v map[string]*string) *CreateVideoModerationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoModerationTaskResponse) SetStatusCode(v int32) *CreateVideoModerationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoModerationTaskResponse) SetBody(v *CreateVideoModerationTaskResponseBody) *CreateVideoModerationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoModerationTaskResponse) Validate() error {
	return dara.Validate(s)
}
