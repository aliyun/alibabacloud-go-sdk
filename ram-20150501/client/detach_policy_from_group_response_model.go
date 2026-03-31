// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachPolicyFromGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachPolicyFromGroupResponse
	GetStatusCode() *int32
	SetBody(v *DetachPolicyFromGroupResponseBody) *DetachPolicyFromGroupResponse
	GetBody() *DetachPolicyFromGroupResponseBody
}

type DetachPolicyFromGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyFromGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachPolicyFromGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachPolicyFromGroupResponse) GetBody() *DetachPolicyFromGroupResponseBody {
	return s.Body
}

func (s *DetachPolicyFromGroupResponse) SetHeaders(v map[string]*string) *DetachPolicyFromGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyFromGroupResponse) SetStatusCode(v int32) *DetachPolicyFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyFromGroupResponse) SetBody(v *DetachPolicyFromGroupResponseBody) *DetachPolicyFromGroupResponse {
	s.Body = v
	return s
}

func (s *DetachPolicyFromGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
