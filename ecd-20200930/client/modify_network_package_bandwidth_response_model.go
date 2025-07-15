// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkPackageBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNetworkPackageBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNetworkPackageBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNetworkPackageBandwidthResponseBody) *ModifyNetworkPackageBandwidthResponse
	GetBody() *ModifyNetworkPackageBandwidthResponseBody
}

type ModifyNetworkPackageBandwidthResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNetworkPackageBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNetworkPackageBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkPackageBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNetworkPackageBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNetworkPackageBandwidthResponse) GetBody() *ModifyNetworkPackageBandwidthResponseBody {
	return s.Body
}

func (s *ModifyNetworkPackageBandwidthResponse) SetHeaders(v map[string]*string) *ModifyNetworkPackageBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkPackageBandwidthResponse) SetStatusCode(v int32) *ModifyNetworkPackageBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthResponse) SetBody(v *ModifyNetworkPackageBandwidthResponseBody) *ModifyNetworkPackageBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifyNetworkPackageBandwidthResponse) Validate() error {
	return dara.Validate(s)
}
