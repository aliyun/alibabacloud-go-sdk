// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressCloudConnectionBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressCloudConnectionBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressCloudConnectionBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressCloudConnectionBandwidthResponseBody) *ModifyExpressCloudConnectionBandwidthResponse
	GetBody() *ModifyExpressCloudConnectionBandwidthResponseBody
}

type ModifyExpressCloudConnectionBandwidthResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressCloudConnectionBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressCloudConnectionBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressCloudConnectionBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressCloudConnectionBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressCloudConnectionBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressCloudConnectionBandwidthResponse) GetBody() *ModifyExpressCloudConnectionBandwidthResponseBody {
	return s.Body
}

func (s *ModifyExpressCloudConnectionBandwidthResponse) SetHeaders(v map[string]*string) *ModifyExpressCloudConnectionBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthResponse) SetStatusCode(v int32) *ModifyExpressCloudConnectionBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthResponse) SetBody(v *ModifyExpressCloudConnectionBandwidthResponseBody) *ModifyExpressCloudConnectionBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
