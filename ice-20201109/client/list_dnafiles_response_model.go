// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNAFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDNAFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDNAFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListDNAFilesResponseBody) *ListDNAFilesResponse
	GetBody() *ListDNAFilesResponseBody
}

type ListDNAFilesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDNAFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDNAFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDNAFilesResponse) GoString() string {
	return s.String()
}

func (s *ListDNAFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDNAFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDNAFilesResponse) GetBody() *ListDNAFilesResponseBody {
	return s.Body
}

func (s *ListDNAFilesResponse) SetHeaders(v map[string]*string) *ListDNAFilesResponse {
	s.Headers = v
	return s
}

func (s *ListDNAFilesResponse) SetStatusCode(v int32) *ListDNAFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDNAFilesResponse) SetBody(v *ListDNAFilesResponseBody) *ListDNAFilesResponse {
	s.Body = v
	return s
}

func (s *ListDNAFilesResponse) Validate() error {
	return dara.Validate(s)
}
