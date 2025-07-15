// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImportTermsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitImportTermsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitImportTermsTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitImportTermsTaskResponseBody) *SubmitImportTermsTaskResponse
	GetBody() *SubmitImportTermsTaskResponseBody
}

type SubmitImportTermsTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImportTermsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImportTermsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportTermsTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitImportTermsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitImportTermsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitImportTermsTaskResponse) GetBody() *SubmitImportTermsTaskResponseBody {
	return s.Body
}

func (s *SubmitImportTermsTaskResponse) SetHeaders(v map[string]*string) *SubmitImportTermsTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitImportTermsTaskResponse) SetStatusCode(v int32) *SubmitImportTermsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImportTermsTaskResponse) SetBody(v *SubmitImportTermsTaskResponseBody) *SubmitImportTermsTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitImportTermsTaskResponse) Validate() error {
	return dara.Validate(s)
}
