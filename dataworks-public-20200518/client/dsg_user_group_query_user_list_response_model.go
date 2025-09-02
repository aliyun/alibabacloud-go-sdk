// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupQueryUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgUserGroupQueryUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgUserGroupQueryUserListResponse
	GetStatusCode() *int32
	SetBody(v *DsgUserGroupQueryUserListResponseBody) *DsgUserGroupQueryUserListResponse
	GetBody() *DsgUserGroupQueryUserListResponseBody
}

type DsgUserGroupQueryUserListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgUserGroupQueryUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgUserGroupQueryUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryUserListResponse) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgUserGroupQueryUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgUserGroupQueryUserListResponse) GetBody() *DsgUserGroupQueryUserListResponseBody {
	return s.Body
}

func (s *DsgUserGroupQueryUserListResponse) SetHeaders(v map[string]*string) *DsgUserGroupQueryUserListResponse {
	s.Headers = v
	return s
}

func (s *DsgUserGroupQueryUserListResponse) SetStatusCode(v int32) *DsgUserGroupQueryUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponse) SetBody(v *DsgUserGroupQueryUserListResponseBody) *DsgUserGroupQueryUserListResponse {
	s.Body = v
	return s
}

func (s *DsgUserGroupQueryUserListResponse) Validate() error {
	return dara.Validate(s)
}
