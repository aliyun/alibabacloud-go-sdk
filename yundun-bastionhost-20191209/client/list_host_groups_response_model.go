// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListHostGroupsResponseBody) *ListHostGroupsResponse
	GetBody() *ListHostGroupsResponseBody
}

type ListHostGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostGroupsResponse) GetBody() *ListHostGroupsResponseBody {
	return s.Body
}

func (s *ListHostGroupsResponse) SetHeaders(v map[string]*string) *ListHostGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupsResponse) SetStatusCode(v int32) *ListHostGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupsResponse) SetBody(v *ListHostGroupsResponseBody) *ListHostGroupsResponse {
	s.Body = v
	return s
}

func (s *ListHostGroupsResponse) Validate() error {
	return dara.Validate(s)
}
