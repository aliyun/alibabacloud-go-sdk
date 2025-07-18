// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentCollectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDocumentCollectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDocumentCollectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDocumentCollectionsResponseBody) *ListDocumentCollectionsResponse
	GetBody() *ListDocumentCollectionsResponseBody
}

type ListDocumentCollectionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDocumentCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDocumentCollectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListDocumentCollectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDocumentCollectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDocumentCollectionsResponse) GetBody() *ListDocumentCollectionsResponseBody {
	return s.Body
}

func (s *ListDocumentCollectionsResponse) SetHeaders(v map[string]*string) *ListDocumentCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListDocumentCollectionsResponse) SetStatusCode(v int32) *ListDocumentCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocumentCollectionsResponse) SetBody(v *ListDocumentCollectionsResponseBody) *ListDocumentCollectionsResponse {
	s.Body = v
	return s
}

func (s *ListDocumentCollectionsResponse) Validate() error {
	return dara.Validate(s)
}
