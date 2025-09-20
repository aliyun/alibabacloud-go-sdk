// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageModerationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageModerationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageModerationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageModerationTaskResponseBody) *CreateImageModerationTaskResponse
	GetBody() *CreateImageModerationTaskResponseBody
}

type CreateImageModerationTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageModerationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageModerationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageModerationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageModerationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageModerationTaskResponse) GetBody() *CreateImageModerationTaskResponseBody {
	return s.Body
}

func (s *CreateImageModerationTaskResponse) SetHeaders(v map[string]*string) *CreateImageModerationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageModerationTaskResponse) SetStatusCode(v int32) *CreateImageModerationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageModerationTaskResponse) SetBody(v *CreateImageModerationTaskResponseBody) *CreateImageModerationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageModerationTaskResponse) Validate() error {
	return dara.Validate(s)
}
