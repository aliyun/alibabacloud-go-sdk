// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDefaultWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDefaultWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *GetDefaultWorkspaceResponseBody) *GetDefaultWorkspaceResponse
	GetBody() *GetDefaultWorkspaceResponseBody
}

type GetDefaultWorkspaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDefaultWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDefaultWorkspaceResponse) GetBody() *GetDefaultWorkspaceResponseBody {
	return s.Body
}

func (s *GetDefaultWorkspaceResponse) SetHeaders(v map[string]*string) *GetDefaultWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultWorkspaceResponse) SetStatusCode(v int32) *GetDefaultWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultWorkspaceResponse) SetBody(v *GetDefaultWorkspaceResponseBody) *GetDefaultWorkspaceResponse {
	s.Body = v
	return s
}

func (s *GetDefaultWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
