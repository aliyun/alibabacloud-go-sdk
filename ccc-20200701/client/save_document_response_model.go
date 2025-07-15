// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveDocumentResponse
	GetStatusCode() *int32
	SetBody(v *SaveDocumentResponseBody) *SaveDocumentResponse
	GetBody() *SaveDocumentResponseBody
}

type SaveDocumentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveDocumentResponse) GoString() string {
	return s.String()
}

func (s *SaveDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveDocumentResponse) GetBody() *SaveDocumentResponseBody {
	return s.Body
}

func (s *SaveDocumentResponse) SetHeaders(v map[string]*string) *SaveDocumentResponse {
	s.Headers = v
	return s
}

func (s *SaveDocumentResponse) SetStatusCode(v int32) *SaveDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveDocumentResponse) SetBody(v *SaveDocumentResponseBody) *SaveDocumentResponse {
	s.Body = v
	return s
}

func (s *SaveDocumentResponse) Validate() error {
	return dara.Validate(s)
}
