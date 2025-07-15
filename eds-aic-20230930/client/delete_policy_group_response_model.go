// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicyGroupResponseBody) *DeletePolicyGroupResponse
	GetBody() *DeletePolicyGroupResponseBody
}

type DeletePolicyGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyGroupResponse) GetBody() *DeletePolicyGroupResponseBody {
	return s.Body
}

func (s *DeletePolicyGroupResponse) SetHeaders(v map[string]*string) *DeletePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyGroupResponse) SetStatusCode(v int32) *DeletePolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyGroupResponse) SetBody(v *DeletePolicyGroupResponseBody) *DeletePolicyGroupResponse {
	s.Body = v
	return s
}

func (s *DeletePolicyGroupResponse) Validate() error {
	return dara.Validate(s)
}
