// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksByWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWorksByWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWorksByWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *QueryWorksByWorkspaceResponseBody) *QueryWorksByWorkspaceResponse
	GetBody() *QueryWorksByWorkspaceResponseBody
}

type QueryWorksByWorkspaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksByWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorksByWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWorksByWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWorksByWorkspaceResponse) GetBody() *QueryWorksByWorkspaceResponseBody {
	return s.Body
}

func (s *QueryWorksByWorkspaceResponse) SetHeaders(v map[string]*string) *QueryWorksByWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksByWorkspaceResponse) SetStatusCode(v int32) *QueryWorksByWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksByWorkspaceResponse) SetBody(v *QueryWorksByWorkspaceResponseBody) *QueryWorksByWorkspaceResponse {
	s.Body = v
	return s
}

func (s *QueryWorksByWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
