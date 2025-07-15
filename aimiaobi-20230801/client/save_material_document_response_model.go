// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMaterialDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveMaterialDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveMaterialDocumentResponse
	GetStatusCode() *int32
	SetBody(v *SaveMaterialDocumentResponseBody) *SaveMaterialDocumentResponse
	GetBody() *SaveMaterialDocumentResponseBody
}

type SaveMaterialDocumentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveMaterialDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveMaterialDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveMaterialDocumentResponse) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveMaterialDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveMaterialDocumentResponse) GetBody() *SaveMaterialDocumentResponseBody {
	return s.Body
}

func (s *SaveMaterialDocumentResponse) SetHeaders(v map[string]*string) *SaveMaterialDocumentResponse {
	s.Headers = v
	return s
}

func (s *SaveMaterialDocumentResponse) SetStatusCode(v int32) *SaveMaterialDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveMaterialDocumentResponse) SetBody(v *SaveMaterialDocumentResponseBody) *SaveMaterialDocumentResponse {
	s.Body = v
	return s
}

func (s *SaveMaterialDocumentResponse) Validate() error {
	return dara.Validate(s)
}
