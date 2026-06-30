// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranslationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranslationTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTranslationTaskResponseBody) *GetTranslationTaskResponse
	GetBody() *GetTranslationTaskResponseBody
}

type GetTranslationTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranslationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranslationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranslationTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTranslationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranslationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranslationTaskResponse) GetBody() *GetTranslationTaskResponseBody {
	return s.Body
}

func (s *GetTranslationTaskResponse) SetHeaders(v map[string]*string) *GetTranslationTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTranslationTaskResponse) SetStatusCode(v int32) *GetTranslationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranslationTaskResponse) SetBody(v *GetTranslationTaskResponseBody) *GetTranslationTaskResponse {
	s.Body = v
	return s
}

func (s *GetTranslationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
