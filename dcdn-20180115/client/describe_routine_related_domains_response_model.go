// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineRelatedDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoutineRelatedDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoutineRelatedDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoutineRelatedDomainsResponseBody) *DescribeRoutineRelatedDomainsResponse
	GetBody() *DescribeRoutineRelatedDomainsResponseBody
}

type DescribeRoutineRelatedDomainsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoutineRelatedDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoutineRelatedDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineRelatedDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineRelatedDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoutineRelatedDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoutineRelatedDomainsResponse) GetBody() *DescribeRoutineRelatedDomainsResponseBody {
	return s.Body
}

func (s *DescribeRoutineRelatedDomainsResponse) SetHeaders(v map[string]*string) *DescribeRoutineRelatedDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineRelatedDomainsResponse) SetStatusCode(v int32) *DescribeRoutineRelatedDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoutineRelatedDomainsResponse) SetBody(v *DescribeRoutineRelatedDomainsResponseBody) *DescribeRoutineRelatedDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeRoutineRelatedDomainsResponse) Validate() error {
	return dara.Validate(s)
}
