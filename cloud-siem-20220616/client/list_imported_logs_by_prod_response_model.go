// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportedLogsByProdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImportedLogsByProdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImportedLogsByProdResponse
	GetStatusCode() *int32
	SetBody(v *ListImportedLogsByProdResponseBody) *ListImportedLogsByProdResponse
	GetBody() *ListImportedLogsByProdResponseBody
}

type ListImportedLogsByProdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImportedLogsByProdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImportedLogsByProdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImportedLogsByProdResponse) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImportedLogsByProdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImportedLogsByProdResponse) GetBody() *ListImportedLogsByProdResponseBody {
	return s.Body
}

func (s *ListImportedLogsByProdResponse) SetHeaders(v map[string]*string) *ListImportedLogsByProdResponse {
	s.Headers = v
	return s
}

func (s *ListImportedLogsByProdResponse) SetStatusCode(v int32) *ListImportedLogsByProdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImportedLogsByProdResponse) SetBody(v *ListImportedLogsByProdResponseBody) *ListImportedLogsByProdResponse {
	s.Body = v
	return s
}

func (s *ListImportedLogsByProdResponse) Validate() error {
	return dara.Validate(s)
}
