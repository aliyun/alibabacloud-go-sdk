// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryAgentGroupResponseBody) *CloudQueryAgentGroupResponse
	GetBody() *CloudQueryAgentGroupResponseBody
}

type CloudQueryAgentGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryAgentGroupResponse) GetBody() *CloudQueryAgentGroupResponseBody {
	return s.Body
}

func (s *CloudQueryAgentGroupResponse) SetHeaders(v map[string]*string) *CloudQueryAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryAgentGroupResponse) SetStatusCode(v int32) *CloudQueryAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryAgentGroupResponse) SetBody(v *CloudQueryAgentGroupResponseBody) *CloudQueryAgentGroupResponse {
	s.Body = v
	return s
}

func (s *CloudQueryAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
