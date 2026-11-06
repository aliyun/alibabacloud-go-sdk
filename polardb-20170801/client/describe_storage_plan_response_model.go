// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStoragePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStoragePlanResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStoragePlanResponseBody) *DescribeStoragePlanResponse
	GetBody() *DescribeStoragePlanResponseBody
}

type DescribeStoragePlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStoragePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStoragePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoragePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStoragePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStoragePlanResponse) GetBody() *DescribeStoragePlanResponseBody {
	return s.Body
}

func (s *DescribeStoragePlanResponse) SetHeaders(v map[string]*string) *DescribeStoragePlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoragePlanResponse) SetStatusCode(v int32) *DescribeStoragePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoragePlanResponse) SetBody(v *DescribeStoragePlanResponseBody) *DescribeStoragePlanResponse {
	s.Body = v
	return s
}

func (s *DescribeStoragePlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
