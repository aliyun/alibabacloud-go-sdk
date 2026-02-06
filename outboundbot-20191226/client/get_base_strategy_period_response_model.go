// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaseStrategyPeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBaseStrategyPeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBaseStrategyPeriodResponse
	GetStatusCode() *int32
	SetBody(v *GetBaseStrategyPeriodResponseBody) *GetBaseStrategyPeriodResponse
	GetBody() *GetBaseStrategyPeriodResponseBody
}

type GetBaseStrategyPeriodResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaseStrategyPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaseStrategyPeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBaseStrategyPeriodResponse) GoString() string {
	return s.String()
}

func (s *GetBaseStrategyPeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBaseStrategyPeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBaseStrategyPeriodResponse) GetBody() *GetBaseStrategyPeriodResponseBody {
	return s.Body
}

func (s *GetBaseStrategyPeriodResponse) SetHeaders(v map[string]*string) *GetBaseStrategyPeriodResponse {
	s.Headers = v
	return s
}

func (s *GetBaseStrategyPeriodResponse) SetStatusCode(v int32) *GetBaseStrategyPeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaseStrategyPeriodResponse) SetBody(v *GetBaseStrategyPeriodResponseBody) *GetBaseStrategyPeriodResponse {
	s.Body = v
	return s
}

func (s *GetBaseStrategyPeriodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
