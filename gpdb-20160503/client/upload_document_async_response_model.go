// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDocumentAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDocumentAsyncResponse
	GetStatusCode() *int32
	SetBody(v *UploadDocumentAsyncResponseBody) *UploadDocumentAsyncResponse
	GetBody() *UploadDocumentAsyncResponseBody
}

type UploadDocumentAsyncResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDocumentAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDocumentAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentAsyncResponse) GoString() string {
	return s.String()
}

func (s *UploadDocumentAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDocumentAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDocumentAsyncResponse) GetBody() *UploadDocumentAsyncResponseBody {
	return s.Body
}

func (s *UploadDocumentAsyncResponse) SetHeaders(v map[string]*string) *UploadDocumentAsyncResponse {
	s.Headers = v
	return s
}

func (s *UploadDocumentAsyncResponse) SetStatusCode(v int32) *UploadDocumentAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDocumentAsyncResponse) SetBody(v *UploadDocumentAsyncResponseBody) *UploadDocumentAsyncResponse {
	s.Body = v
	return s
}

func (s *UploadDocumentAsyncResponse) Validate() error {
	return dara.Validate(s)
}
