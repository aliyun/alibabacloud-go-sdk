// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafDeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSafDeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSafDeOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSafDeOrderResponseBody) *DescribeSafDeOrderResponse
	GetBody() *DescribeSafDeOrderResponseBody
}

type DescribeSafDeOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSafDeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSafDeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafDeOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeSafDeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSafDeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSafDeOrderResponse) GetBody() *DescribeSafDeOrderResponseBody {
	return s.Body
}

func (s *DescribeSafDeOrderResponse) SetHeaders(v map[string]*string) *DescribeSafDeOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeSafDeOrderResponse) SetStatusCode(v int32) *DescribeSafDeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSafDeOrderResponse) SetBody(v *DescribeSafDeOrderResponseBody) *DescribeSafDeOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeSafDeOrderResponse) Validate() error {
	return dara.Validate(s)
}
