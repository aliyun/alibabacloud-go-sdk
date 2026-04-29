// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloudListAgentGroupResponseBody) *CloudListAgentGroupResponse
	GetBody() *CloudListAgentGroupResponseBody
}

type CloudListAgentGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloudListAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListAgentGroupResponse) GetBody() *CloudListAgentGroupResponseBody {
	return s.Body
}

func (s *CloudListAgentGroupResponse) SetHeaders(v map[string]*string) *CloudListAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloudListAgentGroupResponse) SetStatusCode(v int32) *CloudListAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListAgentGroupResponse) SetBody(v *CloudListAgentGroupResponseBody) *CloudListAgentGroupResponse {
	s.Body = v
	return s
}

func (s *CloudListAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
