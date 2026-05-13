// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairSkvDdbWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTairSkvDdbWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTairSkvDdbWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateTairSkvDdbWorkspaceResponseBody) *CreateTairSkvDdbWorkspaceResponse
	GetBody() *CreateTairSkvDdbWorkspaceResponseBody
}

type CreateTairSkvDdbWorkspaceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTairSkvDdbWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTairSkvDdbWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTairSkvDdbWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateTairSkvDdbWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTairSkvDdbWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTairSkvDdbWorkspaceResponse) GetBody() *CreateTairSkvDdbWorkspaceResponseBody {
	return s.Body
}

func (s *CreateTairSkvDdbWorkspaceResponse) SetHeaders(v map[string]*string) *CreateTairSkvDdbWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponse) SetStatusCode(v int32) *CreateTairSkvDdbWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponse) SetBody(v *CreateTairSkvDdbWorkspaceResponseBody) *CreateTairSkvDdbWorkspaceResponse {
	s.Body = v
	return s
}

func (s *CreateTairSkvDdbWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
