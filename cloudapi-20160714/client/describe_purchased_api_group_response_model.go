// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePurchasedApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePurchasedApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribePurchasedApiGroupResponseBody) *DescribePurchasedApiGroupResponse
	GetBody() *DescribePurchasedApiGroupResponseBody
}

type DescribePurchasedApiGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePurchasedApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePurchasedApiGroupResponse) GetBody() *DescribePurchasedApiGroupResponseBody {
	return s.Body
}

func (s *DescribePurchasedApiGroupResponse) SetHeaders(v map[string]*string) *DescribePurchasedApiGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApiGroupResponse) SetStatusCode(v int32) *DescribePurchasedApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedApiGroupResponse) SetBody(v *DescribePurchasedApiGroupResponseBody) *DescribePurchasedApiGroupResponse {
	s.Body = v
	return s
}

func (s *DescribePurchasedApiGroupResponse) Validate() error {
	return dara.Validate(s)
}
