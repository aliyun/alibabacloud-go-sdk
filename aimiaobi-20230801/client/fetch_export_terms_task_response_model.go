// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchExportTermsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchExportTermsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchExportTermsTaskResponse
	GetStatusCode() *int32
	SetBody(v *FetchExportTermsTaskResponseBody) *FetchExportTermsTaskResponse
	GetBody() *FetchExportTermsTaskResponseBody
}

type FetchExportTermsTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchExportTermsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchExportTermsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchExportTermsTaskResponse) GoString() string {
	return s.String()
}

func (s *FetchExportTermsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchExportTermsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchExportTermsTaskResponse) GetBody() *FetchExportTermsTaskResponseBody {
	return s.Body
}

func (s *FetchExportTermsTaskResponse) SetHeaders(v map[string]*string) *FetchExportTermsTaskResponse {
	s.Headers = v
	return s
}

func (s *FetchExportTermsTaskResponse) SetStatusCode(v int32) *FetchExportTermsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchExportTermsTaskResponse) SetBody(v *FetchExportTermsTaskResponseBody) *FetchExportTermsTaskResponse {
	s.Body = v
	return s
}

func (s *FetchExportTermsTaskResponse) Validate() error {
	return dara.Validate(s)
}
