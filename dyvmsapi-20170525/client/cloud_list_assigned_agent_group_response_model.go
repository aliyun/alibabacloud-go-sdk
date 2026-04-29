// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAssignedAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListAssignedAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListAssignedAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloudListAssignedAgentGroupResponseBody) *CloudListAssignedAgentGroupResponse
	GetBody() *CloudListAssignedAgentGroupResponseBody
}

type CloudListAssignedAgentGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListAssignedAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListAssignedAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListAssignedAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloudListAssignedAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListAssignedAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListAssignedAgentGroupResponse) GetBody() *CloudListAssignedAgentGroupResponseBody {
	return s.Body
}

func (s *CloudListAssignedAgentGroupResponse) SetHeaders(v map[string]*string) *CloudListAssignedAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloudListAssignedAgentGroupResponse) SetStatusCode(v int32) *CloudListAssignedAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponse) SetBody(v *CloudListAssignedAgentGroupResponseBody) *CloudListAssignedAgentGroupResponse {
	s.Body = v
	return s
}

func (s *CloudListAssignedAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
