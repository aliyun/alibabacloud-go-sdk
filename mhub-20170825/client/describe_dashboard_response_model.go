// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDashboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDashboardResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDashboardResponseBody) *DescribeDashboardResponse
	GetBody() *DescribeDashboardResponseBody
}

type DescribeDashboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDashboardResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDashboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDashboardResponse) GetBody() *DescribeDashboardResponseBody {
	return s.Body
}

func (s *DescribeDashboardResponse) SetHeaders(v map[string]*string) *DescribeDashboardResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardResponse) SetStatusCode(v int32) *DescribeDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardResponse) SetBody(v *DescribeDashboardResponseBody) *DescribeDashboardResponse {
	s.Body = v
	return s
}

func (s *DescribeDashboardResponse) Validate() error {
	return dara.Validate(s)
}
