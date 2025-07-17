// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageAsyncModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageAsyncModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageAsyncModerationResponse
	GetStatusCode() *int32
	SetBody(v *ImageAsyncModerationResponseBody) *ImageAsyncModerationResponse
	GetBody() *ImageAsyncModerationResponseBody
}

type ImageAsyncModerationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageAsyncModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageAsyncModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageAsyncModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageAsyncModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageAsyncModerationResponse) GetBody() *ImageAsyncModerationResponseBody {
	return s.Body
}

func (s *ImageAsyncModerationResponse) SetHeaders(v map[string]*string) *ImageAsyncModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageAsyncModerationResponse) SetStatusCode(v int32) *ImageAsyncModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageAsyncModerationResponse) SetBody(v *ImageAsyncModerationResponseBody) *ImageAsyncModerationResponse {
	s.Body = v
	return s
}

func (s *ImageAsyncModerationResponse) Validate() error {
	return dara.Validate(s)
}
