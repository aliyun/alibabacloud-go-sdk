// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnUserBillTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnUserBillTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnUserBillTypeResponseBody) *DescribeCdnUserBillTypeResponse
	GetBody() *DescribeCdnUserBillTypeResponseBody
}

type DescribeCdnUserBillTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnUserBillTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnUserBillTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnUserBillTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnUserBillTypeResponse) GetBody() *DescribeCdnUserBillTypeResponseBody {
	return s.Body
}

func (s *DescribeCdnUserBillTypeResponse) SetHeaders(v map[string]*string) *DescribeCdnUserBillTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserBillTypeResponse) SetStatusCode(v int32) *DescribeCdnUserBillTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponse) SetBody(v *DescribeCdnUserBillTypeResponseBody) *DescribeCdnUserBillTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnUserBillTypeResponse) Validate() error {
	return dara.Validate(s)
}
