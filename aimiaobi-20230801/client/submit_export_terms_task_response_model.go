// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitExportTermsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitExportTermsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitExportTermsTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitExportTermsTaskResponseBody) *SubmitExportTermsTaskResponse
	GetBody() *SubmitExportTermsTaskResponseBody
}

type SubmitExportTermsTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitExportTermsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitExportTermsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitExportTermsTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitExportTermsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitExportTermsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitExportTermsTaskResponse) GetBody() *SubmitExportTermsTaskResponseBody {
	return s.Body
}

func (s *SubmitExportTermsTaskResponse) SetHeaders(v map[string]*string) *SubmitExportTermsTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitExportTermsTaskResponse) SetStatusCode(v int32) *SubmitExportTermsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitExportTermsTaskResponse) SetBody(v *SubmitExportTermsTaskResponseBody) *SubmitExportTermsTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitExportTermsTaskResponse) Validate() error {
	return dara.Validate(s)
}
