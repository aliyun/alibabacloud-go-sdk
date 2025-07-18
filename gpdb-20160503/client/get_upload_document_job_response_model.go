// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadDocumentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadDocumentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadDocumentJobResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadDocumentJobResponseBody) *GetUploadDocumentJobResponse
	GetBody() *GetUploadDocumentJobResponseBody
}

type GetUploadDocumentJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadDocumentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadDocumentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDocumentJobResponse) GoString() string {
	return s.String()
}

func (s *GetUploadDocumentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadDocumentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadDocumentJobResponse) GetBody() *GetUploadDocumentJobResponseBody {
	return s.Body
}

func (s *GetUploadDocumentJobResponse) SetHeaders(v map[string]*string) *GetUploadDocumentJobResponse {
	s.Headers = v
	return s
}

func (s *GetUploadDocumentJobResponse) SetStatusCode(v int32) *GetUploadDocumentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadDocumentJobResponse) SetBody(v *GetUploadDocumentJobResponseBody) *GetUploadDocumentJobResponse {
	s.Body = v
	return s
}

func (s *GetUploadDocumentJobResponse) Validate() error {
	return dara.Validate(s)
}
