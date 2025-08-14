// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoTranslationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoTranslationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoTranslationJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoTranslationJobResponseBody) *SubmitVideoTranslationJobResponse
	GetBody() *SubmitVideoTranslationJobResponseBody
}

type SubmitVideoTranslationJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoTranslationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoTranslationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoTranslationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoTranslationJobResponse) GetBody() *SubmitVideoTranslationJobResponseBody {
	return s.Body
}

func (s *SubmitVideoTranslationJobResponse) SetHeaders(v map[string]*string) *SubmitVideoTranslationJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoTranslationJobResponse) SetStatusCode(v int32) *SubmitVideoTranslationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoTranslationJobResponse) SetBody(v *SubmitVideoTranslationJobResponseBody) *SubmitVideoTranslationJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoTranslationJobResponse) Validate() error {
	return dara.Validate(s)
}
