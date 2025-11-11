// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetDocumentResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetDocumentResponseBody) *GetDatasetDocumentResponse
	GetBody() *GetDatasetDocumentResponseBody
}

type GetDatasetDocumentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetDocumentResponse) GetBody() *GetDatasetDocumentResponseBody {
	return s.Body
}

func (s *GetDatasetDocumentResponse) SetHeaders(v map[string]*string) *GetDatasetDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetDocumentResponse) SetStatusCode(v int32) *GetDatasetDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetDocumentResponse) SetBody(v *GetDatasetDocumentResponseBody) *GetDatasetDocumentResponse {
	s.Body = v
	return s
}

func (s *GetDatasetDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
