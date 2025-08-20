// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLakeStoragesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLakeStoragesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLakeStoragesResponse
	GetStatusCode() *int32
	SetBody(v *ListLakeStoragesResponseBody) *ListLakeStoragesResponse
	GetBody() *ListLakeStoragesResponseBody
}

type ListLakeStoragesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLakeStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLakeStoragesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLakeStoragesResponse) GoString() string {
	return s.String()
}

func (s *ListLakeStoragesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLakeStoragesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLakeStoragesResponse) GetBody() *ListLakeStoragesResponseBody {
	return s.Body
}

func (s *ListLakeStoragesResponse) SetHeaders(v map[string]*string) *ListLakeStoragesResponse {
	s.Headers = v
	return s
}

func (s *ListLakeStoragesResponse) SetStatusCode(v int32) *ListLakeStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLakeStoragesResponse) SetBody(v *ListLakeStoragesResponseBody) *ListLakeStoragesResponse {
	s.Body = v
	return s
}

func (s *ListLakeStoragesResponse) Validate() error {
	return dara.Validate(s)
}
