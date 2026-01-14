// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddAgentGroupResponseBody) *AddAgentGroupResponse
	GetBody() *AddAgentGroupResponseBody
}

type AddAgentGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *AddAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAgentGroupResponse) GetBody() *AddAgentGroupResponseBody {
	return s.Body
}

func (s *AddAgentGroupResponse) SetHeaders(v map[string]*string) *AddAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *AddAgentGroupResponse) SetStatusCode(v int32) *AddAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAgentGroupResponse) SetBody(v *AddAgentGroupResponseBody) *AddAgentGroupResponse {
	s.Body = v
	return s
}

func (s *AddAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
