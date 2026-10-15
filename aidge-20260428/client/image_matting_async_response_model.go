// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageMattingAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageMattingAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageMattingAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ImageMattingAsyncResponseBody) *ImageMattingAsyncResponse
	GetBody() *ImageMattingAsyncResponseBody
}

type ImageMattingAsyncResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageMattingAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageMattingAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImageMattingAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageMattingAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageMattingAsyncResponse) GetBody() *ImageMattingAsyncResponseBody {
	return s.Body
}

func (s *ImageMattingAsyncResponse) SetHeaders(v map[string]*string) *ImageMattingAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImageMattingAsyncResponse) SetStatusCode(v int32) *ImageMattingAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageMattingAsyncResponse) SetBody(v *ImageMattingAsyncResponseBody) *ImageMattingAsyncResponse {
	s.Body = v
	return s
}

func (s *ImageMattingAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
