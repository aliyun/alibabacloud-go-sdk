// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrganizationWorkspaceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrganizationWorkspaceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrganizationWorkspaceListResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrganizationWorkspaceListResponseBody) *QueryOrganizationWorkspaceListResponse
	GetBody() *QueryOrganizationWorkspaceListResponseBody
}

type QueryOrganizationWorkspaceListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrganizationWorkspaceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrganizationWorkspaceListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrganizationWorkspaceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrganizationWorkspaceListResponse) GetBody() *QueryOrganizationWorkspaceListResponseBody {
	return s.Body
}

func (s *QueryOrganizationWorkspaceListResponse) SetHeaders(v map[string]*string) *QueryOrganizationWorkspaceListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrganizationWorkspaceListResponse) SetStatusCode(v int32) *QueryOrganizationWorkspaceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponse) SetBody(v *QueryOrganizationWorkspaceListResponseBody) *QueryOrganizationWorkspaceListResponse {
	s.Body = v
	return s
}

func (s *QueryOrganizationWorkspaceListResponse) Validate() error {
	return dara.Validate(s)
}
