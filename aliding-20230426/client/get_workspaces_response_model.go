// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspacesResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspacesResponseBody) *GetWorkspacesResponse
	GetBody() *GetWorkspacesResponseBody
}

type GetWorkspacesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspacesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspacesResponse) GetBody() *GetWorkspacesResponseBody {
	return s.Body
}

func (s *GetWorkspacesResponse) SetHeaders(v map[string]*string) *GetWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspacesResponse) SetStatusCode(v int32) *GetWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspacesResponse) SetBody(v *GetWorkspacesResponseBody) *GetWorkspacesResponse {
	s.Body = v
	return s
}

func (s *GetWorkspacesResponse) Validate() error {
	return dara.Validate(s)
}
