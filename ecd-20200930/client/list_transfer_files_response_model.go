// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransferFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransferFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransferFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListTransferFilesResponseBody) *ListTransferFilesResponse
	GetBody() *ListTransferFilesResponseBody
}

type ListTransferFilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransferFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransferFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFilesResponse) GoString() string {
	return s.String()
}

func (s *ListTransferFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransferFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransferFilesResponse) GetBody() *ListTransferFilesResponseBody {
	return s.Body
}

func (s *ListTransferFilesResponse) SetHeaders(v map[string]*string) *ListTransferFilesResponse {
	s.Headers = v
	return s
}

func (s *ListTransferFilesResponse) SetStatusCode(v int32) *ListTransferFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransferFilesResponse) SetBody(v *ListTransferFilesResponseBody) *ListTransferFilesResponse {
	s.Body = v
	return s
}

func (s *ListTransferFilesResponse) Validate() error {
	return dara.Validate(s)
}
