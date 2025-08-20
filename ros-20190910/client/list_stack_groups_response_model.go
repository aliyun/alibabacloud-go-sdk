// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListStackGroupsResponseBody) *ListStackGroupsResponse
	GetBody() *ListStackGroupsResponseBody
}

type ListStackGroupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackGroupsResponse) GetBody() *ListStackGroupsResponseBody {
	return s.Body
}

func (s *ListStackGroupsResponse) SetHeaders(v map[string]*string) *ListStackGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupsResponse) SetStatusCode(v int32) *ListStackGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackGroupsResponse) SetBody(v *ListStackGroupsResponseBody) *ListStackGroupsResponse {
	s.Body = v
	return s
}

func (s *ListStackGroupsResponse) Validate() error {
	return dara.Validate(s)
}
