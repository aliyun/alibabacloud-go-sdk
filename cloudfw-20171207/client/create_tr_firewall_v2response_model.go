// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrFirewallV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrFirewallV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrFirewallV2Response
	GetStatusCode() *int32
	SetBody(v *CreateTrFirewallV2ResponseBody) *CreateTrFirewallV2Response
	GetBody() *CreateTrFirewallV2ResponseBody
}

type CreateTrFirewallV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrFirewallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrFirewallV2Response) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2Response) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrFirewallV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrFirewallV2Response) GetBody() *CreateTrFirewallV2ResponseBody {
	return s.Body
}

func (s *CreateTrFirewallV2Response) SetHeaders(v map[string]*string) *CreateTrFirewallV2Response {
	s.Headers = v
	return s
}

func (s *CreateTrFirewallV2Response) SetStatusCode(v int32) *CreateTrFirewallV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateTrFirewallV2Response) SetBody(v *CreateTrFirewallV2ResponseBody) *CreateTrFirewallV2Response {
	s.Body = v
	return s
}

func (s *CreateTrFirewallV2Response) Validate() error {
	return dara.Validate(s)
}
