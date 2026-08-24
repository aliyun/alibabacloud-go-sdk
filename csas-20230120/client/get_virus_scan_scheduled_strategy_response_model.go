// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanScheduledStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVirusScanScheduledStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVirusScanScheduledStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetVirusScanScheduledStrategyResponseBody) *GetVirusScanScheduledStrategyResponse
	GetBody() *GetVirusScanScheduledStrategyResponseBody
}

type GetVirusScanScheduledStrategyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVirusScanScheduledStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVirusScanScheduledStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanScheduledStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetVirusScanScheduledStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVirusScanScheduledStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVirusScanScheduledStrategyResponse) GetBody() *GetVirusScanScheduledStrategyResponseBody {
	return s.Body
}

func (s *GetVirusScanScheduledStrategyResponse) SetHeaders(v map[string]*string) *GetVirusScanScheduledStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponse) SetStatusCode(v int32) *GetVirusScanScheduledStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVirusScanScheduledStrategyResponse) SetBody(v *GetVirusScanScheduledStrategyResponseBody) *GetVirusScanScheduledStrategyResponse {
	s.Body = v
	return s
}

func (s *GetVirusScanScheduledStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
