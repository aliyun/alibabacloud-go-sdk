// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationPlusAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageTranslationPlusAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageTranslationPlusAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ImageTranslationPlusAsyncResponseBody) *ImageTranslationPlusAsyncResponse
	GetBody() *ImageTranslationPlusAsyncResponseBody
}

type ImageTranslationPlusAsyncResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageTranslationPlusAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageTranslationPlusAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageTranslationPlusAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageTranslationPlusAsyncResponse) GetBody() *ImageTranslationPlusAsyncResponseBody {
	return s.Body
}

func (s *ImageTranslationPlusAsyncResponse) SetHeaders(v map[string]*string) *ImageTranslationPlusAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImageTranslationPlusAsyncResponse) SetStatusCode(v int32) *ImageTranslationPlusAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageTranslationPlusAsyncResponse) SetBody(v *ImageTranslationPlusAsyncResponseBody) *ImageTranslationPlusAsyncResponse {
	s.Body = v
	return s
}

func (s *ImageTranslationPlusAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
