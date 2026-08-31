// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgAuthorizedWorkspacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKgAuthorizedWorkspacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKgAuthorizedWorkspacesResponse
	GetStatusCode() *int32
	SetBody(v *GetKgAuthorizedWorkspacesResponseBody) *GetKgAuthorizedWorkspacesResponse
	GetBody() *GetKgAuthorizedWorkspacesResponseBody
}

type GetKgAuthorizedWorkspacesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKgAuthorizedWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKgAuthorizedWorkspacesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKgAuthorizedWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *GetKgAuthorizedWorkspacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKgAuthorizedWorkspacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKgAuthorizedWorkspacesResponse) GetBody() *GetKgAuthorizedWorkspacesResponseBody {
	return s.Body
}

func (s *GetKgAuthorizedWorkspacesResponse) SetHeaders(v map[string]*string) *GetKgAuthorizedWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponse) SetStatusCode(v int32) *GetKgAuthorizedWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponse) SetBody(v *GetKgAuthorizedWorkspacesResponseBody) *GetKgAuthorizedWorkspacesResponse {
	s.Body = v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
