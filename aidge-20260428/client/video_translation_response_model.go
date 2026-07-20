// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoTranslationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VideoTranslationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VideoTranslationResponse
	GetStatusCode() *int32
	SetBody(v *VideoTranslationResponseBody) *VideoTranslationResponse
	GetBody() *VideoTranslationResponseBody
}

type VideoTranslationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoTranslationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoTranslationResponse) String() string {
	return dara.Prettify(s)
}

func (s VideoTranslationResponse) GoString() string {
	return s.String()
}

func (s *VideoTranslationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VideoTranslationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VideoTranslationResponse) GetBody() *VideoTranslationResponseBody {
	return s.Body
}

func (s *VideoTranslationResponse) SetHeaders(v map[string]*string) *VideoTranslationResponse {
	s.Headers = v
	return s
}

func (s *VideoTranslationResponse) SetStatusCode(v int32) *VideoTranslationResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoTranslationResponse) SetBody(v *VideoTranslationResponseBody) *VideoTranslationResponse {
	s.Body = v
	return s
}

func (s *VideoTranslationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
