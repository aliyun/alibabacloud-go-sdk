// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *ImportDocumentsResponseBody) *ImportDocumentsResponse
	GetBody() *ImportDocumentsResponseBody
}

type ImportDocumentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ImportDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportDocumentsResponse) GetBody() *ImportDocumentsResponseBody {
	return s.Body
}

func (s *ImportDocumentsResponse) SetHeaders(v map[string]*string) *ImportDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ImportDocumentsResponse) SetStatusCode(v int32) *ImportDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportDocumentsResponse) SetBody(v *ImportDocumentsResponseBody) *ImportDocumentsResponse {
	s.Body = v
	return s
}

func (s *ImportDocumentsResponse) Validate() error {
	return dara.Validate(s)
}
