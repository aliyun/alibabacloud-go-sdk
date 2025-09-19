// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarStrategyTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarStrategyTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarStrategyTaskDetailResponseBody) *DescribeSoarStrategyTaskDetailResponse
	GetBody() *DescribeSoarStrategyTaskDetailResponseBody
}

type DescribeSoarStrategyTaskDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarStrategyTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarStrategyTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarStrategyTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarStrategyTaskDetailResponse) GetBody() *DescribeSoarStrategyTaskDetailResponseBody {
	return s.Body
}

func (s *DescribeSoarStrategyTaskDetailResponse) SetHeaders(v map[string]*string) *DescribeSoarStrategyTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponse) SetStatusCode(v int32) *DescribeSoarStrategyTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponse) SetBody(v *DescribeSoarStrategyTaskDetailResponseBody) *DescribeSoarStrategyTaskDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarStrategyTaskDetailResponse) Validate() error {
	return dara.Validate(s)
}
