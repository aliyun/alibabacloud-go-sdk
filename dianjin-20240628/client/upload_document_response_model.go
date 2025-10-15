// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDocumentResponse
	GetStatusCode() *int32
	SetBody(v *UploadDocumentResponseBody) *UploadDocumentResponse
	GetBody() *UploadDocumentResponseBody
}

type UploadDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentResponse) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDocumentResponse) GetBody() *UploadDocumentResponseBody {
	return s.Body
}

func (s *UploadDocumentResponse) SetHeaders(v map[string]*string) *UploadDocumentResponse {
	s.Headers = v
	return s
}

func (s *UploadDocumentResponse) SetStatusCode(v int32) *UploadDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDocumentResponse) SetBody(v *UploadDocumentResponseBody) *UploadDocumentResponse {
	s.Body = v
	return s
}

func (s *UploadDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
