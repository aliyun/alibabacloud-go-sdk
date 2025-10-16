// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcFirewallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcFirewallTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcFirewallTaskResponseBody) *CreateVpcFirewallTaskResponse
	GetBody() *CreateVpcFirewallTaskResponseBody
}

type CreateVpcFirewallTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcFirewallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcFirewallTaskResponse) GetBody() *CreateVpcFirewallTaskResponseBody {
	return s.Body
}

func (s *CreateVpcFirewallTaskResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallTaskResponse) SetStatusCode(v int32) *CreateVpcFirewallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallTaskResponse) SetBody(v *CreateVpcFirewallTaskResponseBody) *CreateVpcFirewallTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVpcFirewallTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
