// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedDataKeyVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManagedDataKeyVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManagedDataKeyVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListManagedDataKeyVersionsResponseBody) *ListManagedDataKeyVersionsResponse
	GetBody() *ListManagedDataKeyVersionsResponseBody
}

type ListManagedDataKeyVersionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManagedDataKeyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManagedDataKeyVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManagedDataKeyVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListManagedDataKeyVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManagedDataKeyVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManagedDataKeyVersionsResponse) GetBody() *ListManagedDataKeyVersionsResponseBody {
	return s.Body
}

func (s *ListManagedDataKeyVersionsResponse) SetHeaders(v map[string]*string) *ListManagedDataKeyVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListManagedDataKeyVersionsResponse) SetStatusCode(v int32) *ListManagedDataKeyVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManagedDataKeyVersionsResponse) SetBody(v *ListManagedDataKeyVersionsResponseBody) *ListManagedDataKeyVersionsResponse {
	s.Body = v
	return s
}

func (s *ListManagedDataKeyVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
