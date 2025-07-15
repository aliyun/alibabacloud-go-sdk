// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCdsFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCdsFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCdsFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListCdsFilesResponseBody) *ListCdsFilesResponse
	GetBody() *ListCdsFilesResponseBody
}

type ListCdsFilesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCdsFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCdsFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCdsFilesResponse) GoString() string {
	return s.String()
}

func (s *ListCdsFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCdsFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCdsFilesResponse) GetBody() *ListCdsFilesResponseBody {
	return s.Body
}

func (s *ListCdsFilesResponse) SetHeaders(v map[string]*string) *ListCdsFilesResponse {
	s.Headers = v
	return s
}

func (s *ListCdsFilesResponse) SetStatusCode(v int32) *ListCdsFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCdsFilesResponse) SetBody(v *ListCdsFilesResponseBody) *ListCdsFilesResponse {
	s.Body = v
	return s
}

func (s *ListCdsFilesResponse) Validate() error {
	return dara.Validate(s)
}
