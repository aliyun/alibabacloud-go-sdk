// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetDocumentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetDocumentResponseBody) *DeleteDatasetDocumentResponse
	GetBody() *DeleteDatasetDocumentResponseBody
}

type DeleteDatasetDocumentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetDocumentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetDocumentResponse) GetBody() *DeleteDatasetDocumentResponseBody {
	return s.Body
}

func (s *DeleteDatasetDocumentResponse) SetHeaders(v map[string]*string) *DeleteDatasetDocumentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetDocumentResponse) SetStatusCode(v int32) *DeleteDatasetDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetDocumentResponse) SetBody(v *DeleteDatasetDocumentResponseBody) *DeleteDatasetDocumentResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
