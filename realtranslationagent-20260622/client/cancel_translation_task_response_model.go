// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTranslationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelTranslationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelTranslationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelTranslationTaskResponseBody) *CancelTranslationTaskResponse
	GetBody() *CancelTranslationTaskResponseBody
}

type CancelTranslationTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelTranslationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelTranslationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelTranslationTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTranslationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelTranslationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelTranslationTaskResponse) GetBody() *CancelTranslationTaskResponseBody {
	return s.Body
}

func (s *CancelTranslationTaskResponse) SetHeaders(v map[string]*string) *CancelTranslationTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTranslationTaskResponse) SetStatusCode(v int32) *CancelTranslationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelTranslationTaskResponse) SetBody(v *CancelTranslationTaskResponseBody) *CancelTranslationTaskResponse {
	s.Body = v
	return s
}

func (s *CancelTranslationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
