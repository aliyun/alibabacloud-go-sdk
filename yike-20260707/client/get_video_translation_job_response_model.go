// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoTranslationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoTranslationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoTranslationJobResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoTranslationJobResponseBody) *GetVideoTranslationJobResponse
	GetBody() *GetVideoTranslationJobResponseBody
}

type GetVideoTranslationJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoTranslationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoTranslationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoTranslationJobResponse) GoString() string {
	return s.String()
}

func (s *GetVideoTranslationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoTranslationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoTranslationJobResponse) GetBody() *GetVideoTranslationJobResponseBody {
	return s.Body
}

func (s *GetVideoTranslationJobResponse) SetHeaders(v map[string]*string) *GetVideoTranslationJobResponse {
	s.Headers = v
	return s
}

func (s *GetVideoTranslationJobResponse) SetStatusCode(v int32) *GetVideoTranslationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoTranslationJobResponse) SetBody(v *GetVideoTranslationJobResponseBody) *GetVideoTranslationJobResponse {
	s.Body = v
	return s
}

func (s *GetVideoTranslationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
