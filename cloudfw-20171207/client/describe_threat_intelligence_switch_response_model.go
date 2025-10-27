// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatIntelligenceSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeThreatIntelligenceSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeThreatIntelligenceSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeThreatIntelligenceSwitchResponseBody) *DescribeThreatIntelligenceSwitchResponse
	GetBody() *DescribeThreatIntelligenceSwitchResponseBody
}

type DescribeThreatIntelligenceSwitchResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeThreatIntelligenceSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeThreatIntelligenceSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatIntelligenceSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeThreatIntelligenceSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeThreatIntelligenceSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeThreatIntelligenceSwitchResponse) GetBody() *DescribeThreatIntelligenceSwitchResponseBody {
	return s.Body
}

func (s *DescribeThreatIntelligenceSwitchResponse) SetHeaders(v map[string]*string) *DescribeThreatIntelligenceSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponse) SetStatusCode(v int32) *DescribeThreatIntelligenceSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponse) SetBody(v *DescribeThreatIntelligenceSwitchResponseBody) *DescribeThreatIntelligenceSwitchResponse {
	s.Body = v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
