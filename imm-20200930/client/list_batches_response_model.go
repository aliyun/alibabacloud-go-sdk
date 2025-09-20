// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBatchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBatchesResponse
	GetStatusCode() *int32
	SetBody(v *ListBatchesResponseBody) *ListBatchesResponse
	GetBody() *ListBatchesResponseBody
}

type ListBatchesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBatchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBatchesResponse) GoString() string {
	return s.String()
}

func (s *ListBatchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBatchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBatchesResponse) GetBody() *ListBatchesResponseBody {
	return s.Body
}

func (s *ListBatchesResponse) SetHeaders(v map[string]*string) *ListBatchesResponse {
	s.Headers = v
	return s
}

func (s *ListBatchesResponse) SetStatusCode(v int32) *ListBatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchesResponse) SetBody(v *ListBatchesResponseBody) *ListBatchesResponse {
	s.Body = v
	return s
}

func (s *ListBatchesResponse) Validate() error {
	return dara.Validate(s)
}
