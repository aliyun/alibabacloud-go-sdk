// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageModerationResponse
	GetStatusCode() *int32
	SetBody(v *ImageModerationResponseBody) *ImageModerationResponse
	GetBody() *ImageModerationResponseBody
}

type ImageModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageModerationResponse) GetBody() *ImageModerationResponseBody {
	return s.Body
}

func (s *ImageModerationResponse) SetHeaders(v map[string]*string) *ImageModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageModerationResponse) SetStatusCode(v int32) *ImageModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageModerationResponse) SetBody(v *ImageModerationResponseBody) *ImageModerationResponse {
	s.Body = v
	return s
}

func (s *ImageModerationResponse) Validate() error {
	return dara.Validate(s)
}
