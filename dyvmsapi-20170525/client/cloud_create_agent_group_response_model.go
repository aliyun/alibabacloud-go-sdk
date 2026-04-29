// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateAgentGroupResponseBody) *CloudCreateAgentGroupResponse
	GetBody() *CloudCreateAgentGroupResponseBody
}

type CloudCreateAgentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateAgentGroupResponse) GetBody() *CloudCreateAgentGroupResponseBody {
	return s.Body
}

func (s *CloudCreateAgentGroupResponse) SetHeaders(v map[string]*string) *CloudCreateAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateAgentGroupResponse) SetStatusCode(v int32) *CloudCreateAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateAgentGroupResponse) SetBody(v *CloudCreateAgentGroupResponseBody) *CloudCreateAgentGroupResponse {
	s.Body = v
	return s
}

func (s *CloudCreateAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
