// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBaseStrategyPeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBaseStrategyPeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBaseStrategyPeriodResponse
	GetStatusCode() *int32
	SetBody(v *SaveBaseStrategyPeriodResponseBody) *SaveBaseStrategyPeriodResponse
	GetBody() *SaveBaseStrategyPeriodResponseBody
}

type SaveBaseStrategyPeriodResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBaseStrategyPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBaseStrategyPeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBaseStrategyPeriodResponse) GoString() string {
	return s.String()
}

func (s *SaveBaseStrategyPeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBaseStrategyPeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBaseStrategyPeriodResponse) GetBody() *SaveBaseStrategyPeriodResponseBody {
	return s.Body
}

func (s *SaveBaseStrategyPeriodResponse) SetHeaders(v map[string]*string) *SaveBaseStrategyPeriodResponse {
	s.Headers = v
	return s
}

func (s *SaveBaseStrategyPeriodResponse) SetStatusCode(v int32) *SaveBaseStrategyPeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBaseStrategyPeriodResponse) SetBody(v *SaveBaseStrategyPeriodResponseBody) *SaveBaseStrategyPeriodResponse {
	s.Body = v
	return s
}

func (s *SaveBaseStrategyPeriodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
