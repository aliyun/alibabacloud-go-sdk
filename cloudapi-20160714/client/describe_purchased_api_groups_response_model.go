// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedApiGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePurchasedApiGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePurchasedApiGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePurchasedApiGroupsResponseBody) *DescribePurchasedApiGroupsResponse
	GetBody() *DescribePurchasedApiGroupsResponseBody
}

type DescribePurchasedApiGroupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedApiGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePurchasedApiGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePurchasedApiGroupsResponse) GetBody() *DescribePurchasedApiGroupsResponseBody {
	return s.Body
}

func (s *DescribePurchasedApiGroupsResponse) SetHeaders(v map[string]*string) *DescribePurchasedApiGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApiGroupsResponse) SetStatusCode(v int32) *DescribePurchasedApiGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponse) SetBody(v *DescribePurchasedApiGroupsResponseBody) *DescribePurchasedApiGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribePurchasedApiGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
