// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkspaceUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWorkspaceUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWorkspaceUserListResponse
	GetStatusCode() *int32
	SetBody(v *QueryWorkspaceUserListResponseBody) *QueryWorkspaceUserListResponse
	GetBody() *QueryWorkspaceUserListResponseBody
}

type QueryWorkspaceUserListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorkspaceUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorkspaceUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceUserListResponse) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWorkspaceUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWorkspaceUserListResponse) GetBody() *QueryWorkspaceUserListResponseBody {
	return s.Body
}

func (s *QueryWorkspaceUserListResponse) SetHeaders(v map[string]*string) *QueryWorkspaceUserListResponse {
	s.Headers = v
	return s
}

func (s *QueryWorkspaceUserListResponse) SetStatusCode(v int32) *QueryWorkspaceUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorkspaceUserListResponse) SetBody(v *QueryWorkspaceUserListResponseBody) *QueryWorkspaceUserListResponse {
	s.Body = v
	return s
}

func (s *QueryWorkspaceUserListResponse) Validate() error {
	return dara.Validate(s)
}
