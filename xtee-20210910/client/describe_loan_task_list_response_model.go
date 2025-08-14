// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoanTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoanTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoanTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoanTaskListResponseBody) *DescribeLoanTaskListResponse
	GetBody() *DescribeLoanTaskListResponseBody
}

type DescribeLoanTaskListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoanTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoanTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoanTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoanTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoanTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoanTaskListResponse) GetBody() *DescribeLoanTaskListResponseBody {
	return s.Body
}

func (s *DescribeLoanTaskListResponse) SetHeaders(v map[string]*string) *DescribeLoanTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoanTaskListResponse) SetStatusCode(v int32) *DescribeLoanTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoanTaskListResponse) SetBody(v *DescribeLoanTaskListResponseBody) *DescribeLoanTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeLoanTaskListResponse) Validate() error {
	return dara.Validate(s)
}
