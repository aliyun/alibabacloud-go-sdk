// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenterPolicyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenterPolicyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenterPolicyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenterPolicyListResponseBody) *DescribeCenterPolicyListResponse
	GetBody() *DescribeCenterPolicyListResponseBody
}

type DescribeCenterPolicyListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenterPolicyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenterPolicyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenterPolicyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenterPolicyListResponse) GetBody() *DescribeCenterPolicyListResponseBody {
	return s.Body
}

func (s *DescribeCenterPolicyListResponse) SetHeaders(v map[string]*string) *DescribeCenterPolicyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenterPolicyListResponse) SetStatusCode(v int32) *DescribeCenterPolicyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenterPolicyListResponse) SetBody(v *DescribeCenterPolicyListResponseBody) *DescribeCenterPolicyListResponse {
	s.Body = v
	return s
}

func (s *DescribeCenterPolicyListResponse) Validate() error {
	return dara.Validate(s)
}
