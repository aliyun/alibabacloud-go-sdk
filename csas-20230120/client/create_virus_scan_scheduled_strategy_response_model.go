// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanScheduledStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirusScanScheduledStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirusScanScheduledStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirusScanScheduledStrategyResponseBody) *CreateVirusScanScheduledStrategyResponse
	GetBody() *CreateVirusScanScheduledStrategyResponseBody
}

type CreateVirusScanScheduledStrategyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirusScanScheduledStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirusScanScheduledStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanScheduledStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateVirusScanScheduledStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirusScanScheduledStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirusScanScheduledStrategyResponse) GetBody() *CreateVirusScanScheduledStrategyResponseBody {
	return s.Body
}

func (s *CreateVirusScanScheduledStrategyResponse) SetHeaders(v map[string]*string) *CreateVirusScanScheduledStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateVirusScanScheduledStrategyResponse) SetStatusCode(v int32) *CreateVirusScanScheduledStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirusScanScheduledStrategyResponse) SetBody(v *CreateVirusScanScheduledStrategyResponseBody) *CreateVirusScanScheduledStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateVirusScanScheduledStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
