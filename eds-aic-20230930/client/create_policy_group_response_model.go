// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolicyGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolicyGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolicyGroupResponseBody) *CreatePolicyGroupResponse
	GetBody() *CreatePolicyGroupResponseBody
}

type CreatePolicyGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolicyGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolicyGroupResponse) GetBody() *CreatePolicyGroupResponseBody {
	return s.Body
}

func (s *CreatePolicyGroupResponse) SetHeaders(v map[string]*string) *CreatePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyGroupResponse) SetStatusCode(v int32) *CreatePolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyGroupResponse) SetBody(v *CreatePolicyGroupResponseBody) *CreatePolicyGroupResponse {
	s.Body = v
	return s
}

func (s *CreatePolicyGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
