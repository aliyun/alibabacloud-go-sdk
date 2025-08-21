// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetworkRegionBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigNetworkRegionBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigNetworkRegionBlockResponse
	GetStatusCode() *int32
	SetBody(v *ConfigNetworkRegionBlockResponseBody) *ConfigNetworkRegionBlockResponse
	GetBody() *ConfigNetworkRegionBlockResponseBody
}

type ConfigNetworkRegionBlockResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigNetworkRegionBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigNetworkRegionBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetworkRegionBlockResponse) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRegionBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigNetworkRegionBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigNetworkRegionBlockResponse) GetBody() *ConfigNetworkRegionBlockResponseBody {
	return s.Body
}

func (s *ConfigNetworkRegionBlockResponse) SetHeaders(v map[string]*string) *ConfigNetworkRegionBlockResponse {
	s.Headers = v
	return s
}

func (s *ConfigNetworkRegionBlockResponse) SetStatusCode(v int32) *ConfigNetworkRegionBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigNetworkRegionBlockResponse) SetBody(v *ConfigNetworkRegionBlockResponseBody) *ConfigNetworkRegionBlockResponse {
	s.Body = v
	return s
}

func (s *ConfigNetworkRegionBlockResponse) Validate() error {
	return dara.Validate(s)
}
