// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayUpBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySmartAccessGatewayUpBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySmartAccessGatewayUpBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifySmartAccessGatewayUpBandwidthResponseBody) *ModifySmartAccessGatewayUpBandwidthResponse
	GetBody() *ModifySmartAccessGatewayUpBandwidthResponseBody
}

type ModifySmartAccessGatewayUpBandwidthResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmartAccessGatewayUpBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySmartAccessGatewayUpBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayUpBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayUpBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySmartAccessGatewayUpBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySmartAccessGatewayUpBandwidthResponse) GetBody() *ModifySmartAccessGatewayUpBandwidthResponseBody {
	return s.Body
}

func (s *ModifySmartAccessGatewayUpBandwidthResponse) SetHeaders(v map[string]*string) *ModifySmartAccessGatewayUpBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthResponse) SetStatusCode(v int32) *ModifySmartAccessGatewayUpBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthResponse) SetBody(v *ModifySmartAccessGatewayUpBandwidthResponseBody) *ModifySmartAccessGatewayUpBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
