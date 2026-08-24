// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVulScanScheduledStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVulScanScheduledStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVulScanScheduledStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVulScanScheduledStrategyResponseBody) *UpdateVulScanScheduledStrategyResponse
	GetBody() *UpdateVulScanScheduledStrategyResponseBody
}

type UpdateVulScanScheduledStrategyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVulScanScheduledStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVulScanScheduledStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanScheduledStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateVulScanScheduledStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVulScanScheduledStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVulScanScheduledStrategyResponse) GetBody() *UpdateVulScanScheduledStrategyResponseBody {
	return s.Body
}

func (s *UpdateVulScanScheduledStrategyResponse) SetHeaders(v map[string]*string) *UpdateVulScanScheduledStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponse) SetStatusCode(v int32) *UpdateVulScanScheduledStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponse) SetBody(v *UpdateVulScanScheduledStrategyResponseBody) *UpdateVulScanScheduledStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateVulScanScheduledStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
