// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatasetDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDatasetDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDatasetDocumentResponse
	GetStatusCode() *int32
	SetBody(v *AddDatasetDocumentResponseBody) *AddDatasetDocumentResponse
	GetBody() *AddDatasetDocumentResponseBody
}

type AddDatasetDocumentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDatasetDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDatasetDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentResponse) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDatasetDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDatasetDocumentResponse) GetBody() *AddDatasetDocumentResponseBody {
	return s.Body
}

func (s *AddDatasetDocumentResponse) SetHeaders(v map[string]*string) *AddDatasetDocumentResponse {
	s.Headers = v
	return s
}

func (s *AddDatasetDocumentResponse) SetStatusCode(v int32) *AddDatasetDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDatasetDocumentResponse) SetBody(v *AddDatasetDocumentResponseBody) *AddDatasetDocumentResponse {
	s.Body = v
	return s
}

func (s *AddDatasetDocumentResponse) Validate() error {
	return dara.Validate(s)
}
