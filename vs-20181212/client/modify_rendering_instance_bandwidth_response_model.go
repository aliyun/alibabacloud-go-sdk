// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRenderingInstanceBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRenderingInstanceBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRenderingInstanceBandwidthResponseBody) *ModifyRenderingInstanceBandwidthResponse
	GetBody() *ModifyRenderingInstanceBandwidthResponseBody
}

type ModifyRenderingInstanceBandwidthResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRenderingInstanceBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRenderingInstanceBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRenderingInstanceBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRenderingInstanceBandwidthResponse) GetBody() *ModifyRenderingInstanceBandwidthResponseBody {
	return s.Body
}

func (s *ModifyRenderingInstanceBandwidthResponse) SetHeaders(v map[string]*string) *ModifyRenderingInstanceBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyRenderingInstanceBandwidthResponse) SetStatusCode(v int32) *ModifyRenderingInstanceBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRenderingInstanceBandwidthResponse) SetBody(v *ModifyRenderingInstanceBandwidthResponseBody) *ModifyRenderingInstanceBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifyRenderingInstanceBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
