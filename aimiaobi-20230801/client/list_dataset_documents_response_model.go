// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasetDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasetDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasetDocumentsResponseBody) *ListDatasetDocumentsResponse
	GetBody() *ListDatasetDocumentsResponseBody
}

type ListDatasetDocumentsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasetDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasetDocumentsResponse) GetBody() *ListDatasetDocumentsResponseBody {
	return s.Body
}

func (s *ListDatasetDocumentsResponse) SetHeaders(v map[string]*string) *ListDatasetDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetDocumentsResponse) SetStatusCode(v int32) *ListDatasetDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetDocumentsResponse) SetBody(v *ListDatasetDocumentsResponseBody) *ListDatasetDocumentsResponse {
	s.Body = v
	return s
}

func (s *ListDatasetDocumentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
