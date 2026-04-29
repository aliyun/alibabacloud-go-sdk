// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteAgentGroupResponseBody) *CloudDeleteAgentGroupResponse
	GetBody() *CloudDeleteAgentGroupResponseBody
}

type CloudDeleteAgentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteAgentGroupResponse) GetBody() *CloudDeleteAgentGroupResponseBody {
	return s.Body
}

func (s *CloudDeleteAgentGroupResponse) SetHeaders(v map[string]*string) *CloudDeleteAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteAgentGroupResponse) SetStatusCode(v int32) *CloudDeleteAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteAgentGroupResponse) SetBody(v *CloudDeleteAgentGroupResponseBody) *CloudDeleteAgentGroupResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
