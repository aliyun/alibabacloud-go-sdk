// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkRuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNetworkRuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNetworkRuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNetworkRuleAttributeResponseBody) *ModifyNetworkRuleAttributeResponse
	GetBody() *ModifyNetworkRuleAttributeResponseBody
}

type ModifyNetworkRuleAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNetworkRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNetworkRuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkRuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNetworkRuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNetworkRuleAttributeResponse) GetBody() *ModifyNetworkRuleAttributeResponseBody {
	return s.Body
}

func (s *ModifyNetworkRuleAttributeResponse) SetHeaders(v map[string]*string) *ModifyNetworkRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkRuleAttributeResponse) SetStatusCode(v int32) *ModifyNetworkRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNetworkRuleAttributeResponse) SetBody(v *ModifyNetworkRuleAttributeResponseBody) *ModifyNetworkRuleAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyNetworkRuleAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
