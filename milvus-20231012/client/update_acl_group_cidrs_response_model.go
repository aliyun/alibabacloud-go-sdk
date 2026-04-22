// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclGroupCidrsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAclGroupCidrsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAclGroupCidrsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAclGroupCidrsResponseBody) *UpdateAclGroupCidrsResponse
	GetBody() *UpdateAclGroupCidrsResponseBody
}

type UpdateAclGroupCidrsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAclGroupCidrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAclGroupCidrsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclGroupCidrsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclGroupCidrsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAclGroupCidrsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAclGroupCidrsResponse) GetBody() *UpdateAclGroupCidrsResponseBody {
	return s.Body
}

func (s *UpdateAclGroupCidrsResponse) SetHeaders(v map[string]*string) *UpdateAclGroupCidrsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclGroupCidrsResponse) SetStatusCode(v int32) *UpdateAclGroupCidrsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAclGroupCidrsResponse) SetBody(v *UpdateAclGroupCidrsResponseBody) *UpdateAclGroupCidrsResponse {
	s.Body = v
	return s
}

func (s *UpdateAclGroupCidrsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
