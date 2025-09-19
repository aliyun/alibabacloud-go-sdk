// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarStrategyTaskParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarStrategyTaskParamsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarStrategyTaskParamsResponseBody) *DescribeSoarStrategyTaskParamsResponse
	GetBody() *DescribeSoarStrategyTaskParamsResponseBody
}

type DescribeSoarStrategyTaskParamsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarStrategyTaskParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarStrategyTaskParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskParamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarStrategyTaskParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarStrategyTaskParamsResponse) GetBody() *DescribeSoarStrategyTaskParamsResponseBody {
	return s.Body
}

func (s *DescribeSoarStrategyTaskParamsResponse) SetHeaders(v map[string]*string) *DescribeSoarStrategyTaskParamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarStrategyTaskParamsResponse) SetStatusCode(v int32) *DescribeSoarStrategyTaskParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarStrategyTaskParamsResponse) SetBody(v *DescribeSoarStrategyTaskParamsResponseBody) *DescribeSoarStrategyTaskParamsResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarStrategyTaskParamsResponse) Validate() error {
	return dara.Validate(s)
}
