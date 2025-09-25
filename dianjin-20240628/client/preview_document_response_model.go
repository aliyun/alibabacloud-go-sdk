// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewDocumentResponse
	GetStatusCode() *int32
	SetBody(v *PreviewDocumentResponseBody) *PreviewDocumentResponse
	GetBody() *PreviewDocumentResponseBody
}

type PreviewDocumentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewDocumentResponse) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewDocumentResponse) GetBody() *PreviewDocumentResponseBody {
	return s.Body
}

func (s *PreviewDocumentResponse) SetHeaders(v map[string]*string) *PreviewDocumentResponse {
	s.Headers = v
	return s
}

func (s *PreviewDocumentResponse) SetStatusCode(v int32) *PreviewDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewDocumentResponse) SetBody(v *PreviewDocumentResponseBody) *PreviewDocumentResponse {
	s.Body = v
	return s
}

func (s *PreviewDocumentResponse) Validate() error {
	return dara.Validate(s)
}
