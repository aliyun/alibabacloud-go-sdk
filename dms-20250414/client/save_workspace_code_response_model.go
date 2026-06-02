// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWorkspaceCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveWorkspaceCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveWorkspaceCodeResponse
	GetStatusCode() *int32
	SetBody(v *SaveWorkspaceCodeResponseBody) *SaveWorkspaceCodeResponse
	GetBody() *SaveWorkspaceCodeResponseBody
}

type SaveWorkspaceCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveWorkspaceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveWorkspaceCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveWorkspaceCodeResponse) GoString() string {
	return s.String()
}

func (s *SaveWorkspaceCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveWorkspaceCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveWorkspaceCodeResponse) GetBody() *SaveWorkspaceCodeResponseBody {
	return s.Body
}

func (s *SaveWorkspaceCodeResponse) SetHeaders(v map[string]*string) *SaveWorkspaceCodeResponse {
	s.Headers = v
	return s
}

func (s *SaveWorkspaceCodeResponse) SetStatusCode(v int32) *SaveWorkspaceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWorkspaceCodeResponse) SetBody(v *SaveWorkspaceCodeResponseBody) *SaveWorkspaceCodeResponse {
	s.Body = v
	return s
}

func (s *SaveWorkspaceCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
