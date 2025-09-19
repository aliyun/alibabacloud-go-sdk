// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarStrategyTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarStrategyTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarStrategyTaskResultResponseBody) *DescribeSoarStrategyTaskResultResponse
	GetBody() *DescribeSoarStrategyTaskResultResponseBody
}

type DescribeSoarStrategyTaskResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarStrategyTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarStrategyTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarStrategyTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarStrategyTaskResultResponse) GetBody() *DescribeSoarStrategyTaskResultResponseBody {
	return s.Body
}

func (s *DescribeSoarStrategyTaskResultResponse) SetHeaders(v map[string]*string) *DescribeSoarStrategyTaskResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarStrategyTaskResultResponse) SetStatusCode(v int32) *DescribeSoarStrategyTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarStrategyTaskResultResponse) SetBody(v *DescribeSoarStrategyTaskResultResponseBody) *DescribeSoarStrategyTaskResultResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarStrategyTaskResultResponse) Validate() error {
	return dara.Validate(s)
}
