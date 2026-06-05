// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppWorkspaceDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppWorkspaceDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppWorkspaceDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *GetAppWorkspaceDirectoryResponseBody) *GetAppWorkspaceDirectoryResponse
	GetBody() *GetAppWorkspaceDirectoryResponseBody
}

type GetAppWorkspaceDirectoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppWorkspaceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppWorkspaceDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppWorkspaceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetAppWorkspaceDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppWorkspaceDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppWorkspaceDirectoryResponse) GetBody() *GetAppWorkspaceDirectoryResponseBody {
	return s.Body
}

func (s *GetAppWorkspaceDirectoryResponse) SetHeaders(v map[string]*string) *GetAppWorkspaceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *GetAppWorkspaceDirectoryResponse) SetStatusCode(v int32) *GetAppWorkspaceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponse) SetBody(v *GetAppWorkspaceDirectoryResponseBody) *GetAppWorkspaceDirectoryResponse {
	s.Body = v
	return s
}

func (s *GetAppWorkspaceDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
