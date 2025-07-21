// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKmsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKmsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKmsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListKmsInstancesResponseBody) *ListKmsInstancesResponse
	GetBody() *ListKmsInstancesResponseBody
}

type ListKmsInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKmsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKmsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKmsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKmsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKmsInstancesResponse) GetBody() *ListKmsInstancesResponseBody {
	return s.Body
}

func (s *ListKmsInstancesResponse) SetHeaders(v map[string]*string) *ListKmsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListKmsInstancesResponse) SetStatusCode(v int32) *ListKmsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKmsInstancesResponse) SetBody(v *ListKmsInstancesResponseBody) *ListKmsInstancesResponse {
	s.Body = v
	return s
}

func (s *ListKmsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
