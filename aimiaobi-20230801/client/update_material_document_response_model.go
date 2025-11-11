// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaterialDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMaterialDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMaterialDocumentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMaterialDocumentResponseBody) *UpdateMaterialDocumentResponse
	GetBody() *UpdateMaterialDocumentResponseBody
}

type UpdateMaterialDocumentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMaterialDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMaterialDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaterialDocumentResponse) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMaterialDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMaterialDocumentResponse) GetBody() *UpdateMaterialDocumentResponseBody {
	return s.Body
}

func (s *UpdateMaterialDocumentResponse) SetHeaders(v map[string]*string) *UpdateMaterialDocumentResponse {
	s.Headers = v
	return s
}

func (s *UpdateMaterialDocumentResponse) SetStatusCode(v int32) *UpdateMaterialDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMaterialDocumentResponse) SetBody(v *UpdateMaterialDocumentResponseBody) *UpdateMaterialDocumentResponse {
	s.Body = v
	return s
}

func (s *UpdateMaterialDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
