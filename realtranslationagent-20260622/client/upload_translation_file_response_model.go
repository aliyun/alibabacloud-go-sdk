// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadTranslationFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadTranslationFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadTranslationFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadTranslationFileResponseBody) *UploadTranslationFileResponse
	GetBody() *UploadTranslationFileResponseBody
}

type UploadTranslationFileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadTranslationFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadTranslationFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadTranslationFileResponse) GoString() string {
	return s.String()
}

func (s *UploadTranslationFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadTranslationFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadTranslationFileResponse) GetBody() *UploadTranslationFileResponseBody {
	return s.Body
}

func (s *UploadTranslationFileResponse) SetHeaders(v map[string]*string) *UploadTranslationFileResponse {
	s.Headers = v
	return s
}

func (s *UploadTranslationFileResponse) SetStatusCode(v int32) *UploadTranslationFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadTranslationFileResponse) SetBody(v *UploadTranslationFileResponseBody) *UploadTranslationFileResponse {
	s.Body = v
	return s
}

func (s *UploadTranslationFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
