// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAntiVirusRealTimeDefenceStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAntiVirusRealTimeDefenceStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAntiVirusRealTimeDefenceStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) *UpdateAntiVirusRealTimeDefenceStrategyResponse
	GetBody() *UpdateAntiVirusRealTimeDefenceStrategyResponseBody
}

type UpdateAntiVirusRealTimeDefenceStrategyResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntiVirusRealTimeDefenceStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAntiVirusRealTimeDefenceStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAntiVirusRealTimeDefenceStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponse) GetBody() *UpdateAntiVirusRealTimeDefenceStrategyResponseBody {
	return s.Body
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponse) SetHeaders(v map[string]*string) *UpdateAntiVirusRealTimeDefenceStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponse) SetStatusCode(v int32) *UpdateAntiVirusRealTimeDefenceStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponse) SetBody(v *UpdateAntiVirusRealTimeDefenceStrategyResponseBody) *UpdateAntiVirusRealTimeDefenceStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateAntiVirusRealTimeDefenceStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
