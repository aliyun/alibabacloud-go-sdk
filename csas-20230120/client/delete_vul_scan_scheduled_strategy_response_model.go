// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulScanScheduledStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVulScanScheduledStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVulScanScheduledStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVulScanScheduledStrategyResponseBody) *DeleteVulScanScheduledStrategyResponse
	GetBody() *DeleteVulScanScheduledStrategyResponseBody
}

type DeleteVulScanScheduledStrategyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVulScanScheduledStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVulScanScheduledStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulScanScheduledStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteVulScanScheduledStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVulScanScheduledStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVulScanScheduledStrategyResponse) GetBody() *DeleteVulScanScheduledStrategyResponseBody {
	return s.Body
}

func (s *DeleteVulScanScheduledStrategyResponse) SetHeaders(v map[string]*string) *DeleteVulScanScheduledStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteVulScanScheduledStrategyResponse) SetStatusCode(v int32) *DeleteVulScanScheduledStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVulScanScheduledStrategyResponse) SetBody(v *DeleteVulScanScheduledStrategyResponseBody) *DeleteVulScanScheduledStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteVulScanScheduledStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
