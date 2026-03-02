// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegePopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrivilegePopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrivilegePopResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrivilegePopResponseBody) *CreatePrivilegePopResponse
	GetBody() *CreatePrivilegePopResponseBody
}

type CreatePrivilegePopResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivilegePopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivilegePopResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegePopResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivilegePopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrivilegePopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrivilegePopResponse) GetBody() *CreatePrivilegePopResponseBody {
	return s.Body
}

func (s *CreatePrivilegePopResponse) SetHeaders(v map[string]*string) *CreatePrivilegePopResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivilegePopResponse) SetStatusCode(v int32) *CreatePrivilegePopResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivilegePopResponse) SetBody(v *CreatePrivilegePopResponseBody) *CreatePrivilegePopResponse {
	s.Body = v
	return s
}

func (s *CreatePrivilegePopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
