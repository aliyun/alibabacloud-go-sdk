// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDatasetDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDatasetDocumentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDatasetDocumentResponseBody) *UpdateDatasetDocumentResponse
	GetBody() *UpdateDatasetDocumentResponseBody
}

type UpdateDatasetDocumentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetDocumentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDatasetDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDatasetDocumentResponse) GetBody() *UpdateDatasetDocumentResponseBody {
	return s.Body
}

func (s *UpdateDatasetDocumentResponse) SetHeaders(v map[string]*string) *UpdateDatasetDocumentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetDocumentResponse) SetStatusCode(v int32) *UpdateDatasetDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetDocumentResponse) SetBody(v *UpdateDatasetDocumentResponseBody) *UpdateDatasetDocumentResponse {
	s.Body = v
	return s
}

func (s *UpdateDatasetDocumentResponse) Validate() error {
	return dara.Validate(s)
}
