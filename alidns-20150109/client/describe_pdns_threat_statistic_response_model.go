// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsThreatStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsThreatStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsThreatStatisticResponseBody) *DescribePdnsThreatStatisticResponse
	GetBody() *DescribePdnsThreatStatisticResponseBody
}

type DescribePdnsThreatStatisticResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsThreatStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsThreatStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsThreatStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsThreatStatisticResponse) GetBody() *DescribePdnsThreatStatisticResponseBody {
	return s.Body
}

func (s *DescribePdnsThreatStatisticResponse) SetHeaders(v map[string]*string) *DescribePdnsThreatStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsThreatStatisticResponse) SetStatusCode(v int32) *DescribePdnsThreatStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponse) SetBody(v *DescribePdnsThreatStatisticResponseBody) *DescribePdnsThreatStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsThreatStatisticResponse) Validate() error {
	return dara.Validate(s)
}
