// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarStrategyTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarStrategyTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarStrategyTasksResponseBody) *DescribeSoarStrategyTasksResponse
	GetBody() *DescribeSoarStrategyTasksResponseBody
}

type DescribeSoarStrategyTasksResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarStrategyTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarStrategyTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarStrategyTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarStrategyTasksResponse) GetBody() *DescribeSoarStrategyTasksResponseBody {
	return s.Body
}

func (s *DescribeSoarStrategyTasksResponse) SetHeaders(v map[string]*string) *DescribeSoarStrategyTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarStrategyTasksResponse) SetStatusCode(v int32) *DescribeSoarStrategyTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponse) SetBody(v *DescribeSoarStrategyTasksResponseBody) *DescribeSoarStrategyTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarStrategyTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
