// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageMattingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageMattingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageMattingResponse
	GetStatusCode() *int32
	SetBody(v *ImageMattingResponseBody) *ImageMattingResponse
	GetBody() *ImageMattingResponseBody
}

type ImageMattingResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageMattingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageMattingResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingResponse) GoString() string {
	return s.String()
}

func (s *ImageMattingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageMattingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageMattingResponse) GetBody() *ImageMattingResponseBody {
	return s.Body
}

func (s *ImageMattingResponse) SetHeaders(v map[string]*string) *ImageMattingResponse {
	s.Headers = v
	return s
}

func (s *ImageMattingResponse) SetStatusCode(v int32) *ImageMattingResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageMattingResponse) SetBody(v *ImageMattingResponseBody) *ImageMattingResponse {
	s.Body = v
	return s
}

func (s *ImageMattingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
