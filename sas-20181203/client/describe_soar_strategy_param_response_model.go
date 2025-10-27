// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarStrategyParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarStrategyParamResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarStrategyParamResponseBody) *DescribeSoarStrategyParamResponse
	GetBody() *DescribeSoarStrategyParamResponseBody
}

type DescribeSoarStrategyParamResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarStrategyParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarStrategyParamResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyParamResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarStrategyParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarStrategyParamResponse) GetBody() *DescribeSoarStrategyParamResponseBody {
	return s.Body
}

func (s *DescribeSoarStrategyParamResponse) SetHeaders(v map[string]*string) *DescribeSoarStrategyParamResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarStrategyParamResponse) SetStatusCode(v int32) *DescribeSoarStrategyParamResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarStrategyParamResponse) SetBody(v *DescribeSoarStrategyParamResponseBody) *DescribeSoarStrategyParamResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarStrategyParamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
