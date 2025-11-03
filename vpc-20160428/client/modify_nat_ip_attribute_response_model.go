// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatIpAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatIpAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatIpAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatIpAttributeResponseBody) *ModifyNatIpAttributeResponse
	GetBody() *ModifyNatIpAttributeResponseBody
}

type ModifyNatIpAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatIpAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatIpAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatIpAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatIpAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatIpAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatIpAttributeResponse) GetBody() *ModifyNatIpAttributeResponseBody {
	return s.Body
}

func (s *ModifyNatIpAttributeResponse) SetHeaders(v map[string]*string) *ModifyNatIpAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatIpAttributeResponse) SetStatusCode(v int32) *ModifyNatIpAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatIpAttributeResponse) SetBody(v *ModifyNatIpAttributeResponseBody) *ModifyNatIpAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyNatIpAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
