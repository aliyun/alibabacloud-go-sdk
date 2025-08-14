// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodPackagingGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodPackagingGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListVodPackagingGroupsResponseBody) *ListVodPackagingGroupsResponse
	GetBody() *ListVodPackagingGroupsResponseBody
}

type ListVodPackagingGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodPackagingGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodPackagingGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListVodPackagingGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodPackagingGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodPackagingGroupsResponse) GetBody() *ListVodPackagingGroupsResponseBody {
	return s.Body
}

func (s *ListVodPackagingGroupsResponse) SetHeaders(v map[string]*string) *ListVodPackagingGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListVodPackagingGroupsResponse) SetStatusCode(v int32) *ListVodPackagingGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodPackagingGroupsResponse) SetBody(v *ListVodPackagingGroupsResponseBody) *ListVodPackagingGroupsResponse {
	s.Body = v
	return s
}

func (s *ListVodPackagingGroupsResponse) Validate() error {
	return dara.Validate(s)
}
