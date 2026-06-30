// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranslationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTranslationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTranslationTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTranslationTaskResponseBody) *SubmitTranslationTaskResponse
	GetBody() *SubmitTranslationTaskResponseBody
}

type SubmitTranslationTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTranslationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTranslationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranslationTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTranslationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTranslationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTranslationTaskResponse) GetBody() *SubmitTranslationTaskResponseBody {
	return s.Body
}

func (s *SubmitTranslationTaskResponse) SetHeaders(v map[string]*string) *SubmitTranslationTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTranslationTaskResponse) SetStatusCode(v int32) *SubmitTranslationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTranslationTaskResponse) SetBody(v *SubmitTranslationTaskResponseBody) *SubmitTranslationTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitTranslationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
