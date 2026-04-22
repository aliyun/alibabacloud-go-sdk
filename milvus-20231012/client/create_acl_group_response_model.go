// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAclGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAclGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAclGroupResponseBody) *CreateAclGroupResponse
	GetBody() *CreateAclGroupResponseBody
}

type CreateAclGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAclGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAclGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAclGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAclGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAclGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAclGroupResponse) GetBody() *CreateAclGroupResponseBody {
	return s.Body
}

func (s *CreateAclGroupResponse) SetHeaders(v map[string]*string) *CreateAclGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAclGroupResponse) SetStatusCode(v int32) *CreateAclGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAclGroupResponse) SetBody(v *CreateAclGroupResponseBody) *CreateAclGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAclGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
