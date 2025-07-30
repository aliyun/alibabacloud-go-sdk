// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceBandwidthResponseBody) *ModifyInstanceBandwidthResponse
	GetBody() *ModifyInstanceBandwidthResponseBody
}

type ModifyInstanceBandwidthResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceBandwidthResponse) GetBody() *ModifyInstanceBandwidthResponseBody {
	return s.Body
}

func (s *ModifyInstanceBandwidthResponse) SetHeaders(v map[string]*string) *ModifyInstanceBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceBandwidthResponse) SetStatusCode(v int32) *ModifyInstanceBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceBandwidthResponse) SetBody(v *ModifyInstanceBandwidthResponseBody) *ModifyInstanceBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceBandwidthResponse) Validate() error {
	return dara.Validate(s)
}
