// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulScanScheduledStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVulScanScheduledStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVulScanScheduledStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateVulScanScheduledStrategyResponseBody) *CreateVulScanScheduledStrategyResponse
	GetBody() *CreateVulScanScheduledStrategyResponseBody
}

type CreateVulScanScheduledStrategyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVulScanScheduledStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVulScanScheduledStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVulScanScheduledStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateVulScanScheduledStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVulScanScheduledStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVulScanScheduledStrategyResponse) GetBody() *CreateVulScanScheduledStrategyResponseBody {
	return s.Body
}

func (s *CreateVulScanScheduledStrategyResponse) SetHeaders(v map[string]*string) *CreateVulScanScheduledStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateVulScanScheduledStrategyResponse) SetStatusCode(v int32) *CreateVulScanScheduledStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVulScanScheduledStrategyResponse) SetBody(v *CreateVulScanScheduledStrategyResponseBody) *CreateVulScanScheduledStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateVulScanScheduledStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
