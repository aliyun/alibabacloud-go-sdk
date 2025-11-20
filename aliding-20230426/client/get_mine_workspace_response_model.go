// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMineWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMineWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMineWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *GetMineWorkspaceResponseBody) *GetMineWorkspaceResponse
	GetBody() *GetMineWorkspaceResponseBody
}

type GetMineWorkspaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMineWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMineWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMineWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMineWorkspaceResponse) GetBody() *GetMineWorkspaceResponseBody {
	return s.Body
}

func (s *GetMineWorkspaceResponse) SetHeaders(v map[string]*string) *GetMineWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetMineWorkspaceResponse) SetStatusCode(v int32) *GetMineWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMineWorkspaceResponse) SetBody(v *GetMineWorkspaceResponseBody) *GetMineWorkspaceResponse {
	s.Body = v
	return s
}

func (s *GetMineWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
