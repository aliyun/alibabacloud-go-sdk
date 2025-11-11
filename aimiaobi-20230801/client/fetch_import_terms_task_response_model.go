// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchImportTermsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchImportTermsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchImportTermsTaskResponse
	GetStatusCode() *int32
	SetBody(v *FetchImportTermsTaskResponseBody) *FetchImportTermsTaskResponse
	GetBody() *FetchImportTermsTaskResponseBody
}

type FetchImportTermsTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchImportTermsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchImportTermsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchImportTermsTaskResponse) GoString() string {
	return s.String()
}

func (s *FetchImportTermsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchImportTermsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchImportTermsTaskResponse) GetBody() *FetchImportTermsTaskResponseBody {
	return s.Body
}

func (s *FetchImportTermsTaskResponse) SetHeaders(v map[string]*string) *FetchImportTermsTaskResponse {
	s.Headers = v
	return s
}

func (s *FetchImportTermsTaskResponse) SetStatusCode(v int32) *FetchImportTermsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchImportTermsTaskResponse) SetBody(v *FetchImportTermsTaskResponseBody) *FetchImportTermsTaskResponse {
	s.Body = v
	return s
}

func (s *FetchImportTermsTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
