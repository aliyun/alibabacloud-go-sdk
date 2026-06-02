// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspaceCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspaceCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspaceCodeResponseBody) *GetWorkspaceCodeResponse
	GetBody() *GetWorkspaceCodeResponseBody
}

type GetWorkspaceCodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodeResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspaceCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspaceCodeResponse) GetBody() *GetWorkspaceCodeResponseBody {
	return s.Body
}

func (s *GetWorkspaceCodeResponse) SetHeaders(v map[string]*string) *GetWorkspaceCodeResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceCodeResponse) SetStatusCode(v int32) *GetWorkspaceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceCodeResponse) SetBody(v *GetWorkspaceCodeResponseBody) *GetWorkspaceCodeResponse {
	s.Body = v
	return s
}

func (s *GetWorkspaceCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
