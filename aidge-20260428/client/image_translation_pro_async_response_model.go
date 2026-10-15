// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationProAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageTranslationProAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageTranslationProAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ImageTranslationProAsyncResponseBody) *ImageTranslationProAsyncResponse
	GetBody() *ImageTranslationProAsyncResponseBody
}

type ImageTranslationProAsyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageTranslationProAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageTranslationProAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImageTranslationProAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageTranslationProAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageTranslationProAsyncResponse) GetBody() *ImageTranslationProAsyncResponseBody {
	return s.Body
}

func (s *ImageTranslationProAsyncResponse) SetHeaders(v map[string]*string) *ImageTranslationProAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImageTranslationProAsyncResponse) SetStatusCode(v int32) *ImageTranslationProAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageTranslationProAsyncResponse) SetBody(v *ImageTranslationProAsyncResponseBody) *ImageTranslationProAsyncResponse {
	s.Body = v
	return s
}

func (s *ImageTranslationProAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
