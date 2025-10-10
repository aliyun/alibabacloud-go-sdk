// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAclsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAclsResponse
	GetStatusCode() *int32
	SetBody(v *ListAclsResponseBody) *ListAclsResponse
	GetBody() *ListAclsResponseBody
}

type ListAclsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAclsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAclsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAclsResponse) GoString() string {
	return s.String()
}

func (s *ListAclsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAclsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAclsResponse) GetBody() *ListAclsResponseBody {
	return s.Body
}

func (s *ListAclsResponse) SetHeaders(v map[string]*string) *ListAclsResponse {
	s.Headers = v
	return s
}

func (s *ListAclsResponse) SetStatusCode(v int32) *ListAclsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclsResponse) SetBody(v *ListAclsResponseBody) *ListAclsResponse {
	s.Body = v
	return s
}

func (s *ListAclsResponse) Validate() error {
	return dara.Validate(s)
}
