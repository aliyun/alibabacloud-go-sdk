// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrFirewallV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrFirewallV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrFirewallV2Response
	GetStatusCode() *int32
	SetBody(v *DeleteTrFirewallV2ResponseBody) *DeleteTrFirewallV2Response
	GetBody() *DeleteTrFirewallV2ResponseBody
}

type DeleteTrFirewallV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrFirewallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrFirewallV2Response) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrFirewallV2Response) GoString() string {
	return s.String()
}

func (s *DeleteTrFirewallV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrFirewallV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrFirewallV2Response) GetBody() *DeleteTrFirewallV2ResponseBody {
	return s.Body
}

func (s *DeleteTrFirewallV2Response) SetHeaders(v map[string]*string) *DeleteTrFirewallV2Response {
	s.Headers = v
	return s
}

func (s *DeleteTrFirewallV2Response) SetStatusCode(v int32) *DeleteTrFirewallV2Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrFirewallV2Response) SetBody(v *DeleteTrFirewallV2ResponseBody) *DeleteTrFirewallV2Response {
	s.Body = v
	return s
}

func (s *DeleteTrFirewallV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
