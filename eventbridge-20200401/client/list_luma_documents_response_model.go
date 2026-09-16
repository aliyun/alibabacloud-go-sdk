// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLumaDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLumaDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *ListLumaDocumentsResponseBody) *ListLumaDocumentsResponse
	GetBody() *ListLumaDocumentsResponseBody
}

type ListLumaDocumentsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLumaDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLumaDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLumaDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListLumaDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLumaDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLumaDocumentsResponse) GetBody() *ListLumaDocumentsResponseBody {
	return s.Body
}

func (s *ListLumaDocumentsResponse) SetHeaders(v map[string]*string) *ListLumaDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListLumaDocumentsResponse) SetStatusCode(v int32) *ListLumaDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLumaDocumentsResponse) SetBody(v *ListLumaDocumentsResponseBody) *ListLumaDocumentsResponse {
	s.Body = v
	return s
}

func (s *ListLumaDocumentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
