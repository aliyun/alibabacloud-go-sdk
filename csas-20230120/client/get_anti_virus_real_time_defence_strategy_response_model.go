// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAntiVirusRealTimeDefenceStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAntiVirusRealTimeDefenceStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAntiVirusRealTimeDefenceStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetAntiVirusRealTimeDefenceStrategyResponseBody) *GetAntiVirusRealTimeDefenceStrategyResponse
	GetBody() *GetAntiVirusRealTimeDefenceStrategyResponseBody
}

type GetAntiVirusRealTimeDefenceStrategyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAntiVirusRealTimeDefenceStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAntiVirusRealTimeDefenceStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAntiVirusRealTimeDefenceStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponse) GetBody() *GetAntiVirusRealTimeDefenceStrategyResponseBody {
	return s.Body
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponse) SetHeaders(v map[string]*string) *GetAntiVirusRealTimeDefenceStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponse) SetStatusCode(v int32) *GetAntiVirusRealTimeDefenceStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponse) SetBody(v *GetAntiVirusRealTimeDefenceStrategyResponseBody) *GetAntiVirusRealTimeDefenceStrategyResponse {
	s.Body = v
	return s
}

func (s *GetAntiVirusRealTimeDefenceStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
