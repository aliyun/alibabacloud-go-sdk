// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulScanScheduledStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulScanScheduledStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulScanScheduledStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetVulScanScheduledStrategyResponseBody) *GetVulScanScheduledStrategyResponse
	GetBody() *GetVulScanScheduledStrategyResponseBody
}

type GetVulScanScheduledStrategyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulScanScheduledStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulScanScheduledStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanScheduledStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetVulScanScheduledStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulScanScheduledStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulScanScheduledStrategyResponse) GetBody() *GetVulScanScheduledStrategyResponseBody {
	return s.Body
}

func (s *GetVulScanScheduledStrategyResponse) SetHeaders(v map[string]*string) *GetVulScanScheduledStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetVulScanScheduledStrategyResponse) SetStatusCode(v int32) *GetVulScanScheduledStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulScanScheduledStrategyResponse) SetBody(v *GetVulScanScheduledStrategyResponseBody) *GetVulScanScheduledStrategyResponse {
	s.Body = v
	return s
}

func (s *GetVulScanScheduledStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
