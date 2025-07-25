// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsRequestStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsRequestStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsRequestStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsRequestStatisticResponseBody) *DescribePdnsRequestStatisticResponse
	GetBody() *DescribePdnsRequestStatisticResponseBody
}

type DescribePdnsRequestStatisticResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsRequestStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsRequestStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsRequestStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsRequestStatisticResponse) GetBody() *DescribePdnsRequestStatisticResponseBody {
	return s.Body
}

func (s *DescribePdnsRequestStatisticResponse) SetHeaders(v map[string]*string) *DescribePdnsRequestStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsRequestStatisticResponse) SetStatusCode(v int32) *DescribePdnsRequestStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponse) SetBody(v *DescribePdnsRequestStatisticResponseBody) *DescribePdnsRequestStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsRequestStatisticResponse) Validate() error {
	return dara.Validate(s)
}
