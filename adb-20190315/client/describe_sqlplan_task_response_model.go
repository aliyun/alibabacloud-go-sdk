// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPlanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLPlanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLPlanTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLPlanTaskResponseBody) *DescribeSQLPlanTaskResponse
	GetBody() *DescribeSQLPlanTaskResponseBody
}

type DescribeSQLPlanTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLPlanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLPlanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPlanTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLPlanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLPlanTaskResponse) GetBody() *DescribeSQLPlanTaskResponseBody {
	return s.Body
}

func (s *DescribeSQLPlanTaskResponse) SetHeaders(v map[string]*string) *DescribeSQLPlanTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlanTaskResponse) SetStatusCode(v int32) *DescribeSQLPlanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPlanTaskResponse) SetBody(v *DescribeSQLPlanTaskResponseBody) *DescribeSQLPlanTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLPlanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
