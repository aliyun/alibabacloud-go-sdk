// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLPlanResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLPlanResponseBody) *DescribeSQLPlanResponse
	GetBody() *DescribeSQLPlanResponseBody
}

type DescribeSQLPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLPlanResponse) GetBody() *DescribeSQLPlanResponseBody {
	return s.Body
}

func (s *DescribeSQLPlanResponse) SetHeaders(v map[string]*string) *DescribeSQLPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlanResponse) SetStatusCode(v int32) *DescribeSQLPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPlanResponse) SetBody(v *DescribeSQLPlanResponseBody) *DescribeSQLPlanResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
