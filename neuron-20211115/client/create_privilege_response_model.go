// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrivilegeResponseBody) *CreatePrivilegeResponse
	GetBody() *CreatePrivilegeResponseBody
}

type CreatePrivilegeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrivilegeResponse) GetBody() *CreatePrivilegeResponseBody {
	return s.Body
}

func (s *CreatePrivilegeResponse) SetHeaders(v map[string]*string) *CreatePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivilegeResponse) SetStatusCode(v int32) *CreatePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivilegeResponse) SetBody(v *CreatePrivilegeResponseBody) *CreatePrivilegeResponse {
	s.Body = v
	return s
}

func (s *CreatePrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
