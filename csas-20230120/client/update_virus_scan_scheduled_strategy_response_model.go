// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirusScanScheduledStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVirusScanScheduledStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVirusScanScheduledStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVirusScanScheduledStrategyResponseBody) *UpdateVirusScanScheduledStrategyResponse
	GetBody() *UpdateVirusScanScheduledStrategyResponseBody
}

type UpdateVirusScanScheduledStrategyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirusScanScheduledStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVirusScanScheduledStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirusScanScheduledStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirusScanScheduledStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVirusScanScheduledStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVirusScanScheduledStrategyResponse) GetBody() *UpdateVirusScanScheduledStrategyResponseBody {
	return s.Body
}

func (s *UpdateVirusScanScheduledStrategyResponse) SetHeaders(v map[string]*string) *UpdateVirusScanScheduledStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirusScanScheduledStrategyResponse) SetStatusCode(v int32) *UpdateVirusScanScheduledStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirusScanScheduledStrategyResponse) SetBody(v *UpdateVirusScanScheduledStrategyResponseBody) *UpdateVirusScanScheduledStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateVirusScanScheduledStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
