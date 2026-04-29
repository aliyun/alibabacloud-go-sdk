// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAssignAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudAssignAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudAssignAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloudAssignAgentGroupResponseBody) *CloudAssignAgentGroupResponse
	GetBody() *CloudAssignAgentGroupResponseBody
}

type CloudAssignAgentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudAssignAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudAssignAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudAssignAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloudAssignAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudAssignAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudAssignAgentGroupResponse) GetBody() *CloudAssignAgentGroupResponseBody {
	return s.Body
}

func (s *CloudAssignAgentGroupResponse) SetHeaders(v map[string]*string) *CloudAssignAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloudAssignAgentGroupResponse) SetStatusCode(v int32) *CloudAssignAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudAssignAgentGroupResponse) SetBody(v *CloudAssignAgentGroupResponseBody) *CloudAssignAgentGroupResponse {
	s.Body = v
	return s
}

func (s *CloudAssignAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
