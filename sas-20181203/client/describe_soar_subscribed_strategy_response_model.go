// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarSubscribedStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarSubscribedStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarSubscribedStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarSubscribedStrategyResponseBody) *DescribeSoarSubscribedStrategyResponse
	GetBody() *DescribeSoarSubscribedStrategyResponseBody
}

type DescribeSoarSubscribedStrategyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarSubscribedStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarSubscribedStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarSubscribedStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarSubscribedStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarSubscribedStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarSubscribedStrategyResponse) GetBody() *DescribeSoarSubscribedStrategyResponseBody {
	return s.Body
}

func (s *DescribeSoarSubscribedStrategyResponse) SetHeaders(v map[string]*string) *DescribeSoarSubscribedStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponse) SetStatusCode(v int32) *DescribeSoarSubscribedStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponse) SetBody(v *DescribeSoarSubscribedStrategyResponseBody) *DescribeSoarSubscribedStrategyResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
