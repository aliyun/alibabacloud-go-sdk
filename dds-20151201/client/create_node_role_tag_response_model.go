// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeRoleTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodeRoleTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodeRoleTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodeRoleTagResponseBody) *CreateNodeRoleTagResponse
	GetBody() *CreateNodeRoleTagResponseBody
}

type CreateNodeRoleTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeRoleTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeRoleTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeRoleTagResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeRoleTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodeRoleTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodeRoleTagResponse) GetBody() *CreateNodeRoleTagResponseBody {
	return s.Body
}

func (s *CreateNodeRoleTagResponse) SetHeaders(v map[string]*string) *CreateNodeRoleTagResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeRoleTagResponse) SetStatusCode(v int32) *CreateNodeRoleTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeRoleTagResponse) SetBody(v *CreateNodeRoleTagResponseBody) *CreateNodeRoleTagResponse {
	s.Body = v
	return s
}

func (s *CreateNodeRoleTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
