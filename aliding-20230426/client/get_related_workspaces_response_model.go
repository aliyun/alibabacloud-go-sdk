// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelatedWorkspacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRelatedWorkspacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRelatedWorkspacesResponse
	GetStatusCode() *int32
	SetBody(v *GetRelatedWorkspacesResponseBody) *GetRelatedWorkspacesResponse
	GetBody() *GetRelatedWorkspacesResponseBody
}

type GetRelatedWorkspacesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRelatedWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRelatedWorkspacesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRelatedWorkspacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRelatedWorkspacesResponse) GetBody() *GetRelatedWorkspacesResponseBody {
	return s.Body
}

func (s *GetRelatedWorkspacesResponse) SetHeaders(v map[string]*string) *GetRelatedWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *GetRelatedWorkspacesResponse) SetStatusCode(v int32) *GetRelatedWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelatedWorkspacesResponse) SetBody(v *GetRelatedWorkspacesResponseBody) *GetRelatedWorkspacesResponse {
	s.Body = v
	return s
}

func (s *GetRelatedWorkspacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
