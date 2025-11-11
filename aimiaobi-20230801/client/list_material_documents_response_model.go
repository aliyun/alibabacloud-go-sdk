// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaterialDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMaterialDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMaterialDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *ListMaterialDocumentsResponseBody) *ListMaterialDocumentsResponse
	GetBody() *ListMaterialDocumentsResponseBody
}

type ListMaterialDocumentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMaterialDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMaterialDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMaterialDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMaterialDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMaterialDocumentsResponse) GetBody() *ListMaterialDocumentsResponseBody {
	return s.Body
}

func (s *ListMaterialDocumentsResponse) SetHeaders(v map[string]*string) *ListMaterialDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListMaterialDocumentsResponse) SetStatusCode(v int32) *ListMaterialDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMaterialDocumentsResponse) SetBody(v *ListMaterialDocumentsResponseBody) *ListMaterialDocumentsResponse {
	s.Body = v
	return s
}

func (s *ListMaterialDocumentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
