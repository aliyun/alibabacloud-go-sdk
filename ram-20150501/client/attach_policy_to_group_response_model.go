// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachPolicyToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachPolicyToGroupResponse
	GetStatusCode() *int32
	SetBody(v *AttachPolicyToGroupResponseBody) *AttachPolicyToGroupResponse
	GetBody() *AttachPolicyToGroupResponseBody
}

type AttachPolicyToGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyToGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachPolicyToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachPolicyToGroupResponse) GetBody() *AttachPolicyToGroupResponseBody {
	return s.Body
}

func (s *AttachPolicyToGroupResponse) SetHeaders(v map[string]*string) *AttachPolicyToGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyToGroupResponse) SetStatusCode(v int32) *AttachPolicyToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyToGroupResponse) SetBody(v *AttachPolicyToGroupResponseBody) *AttachPolicyToGroupResponse {
	s.Body = v
	return s
}

func (s *AttachPolicyToGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
