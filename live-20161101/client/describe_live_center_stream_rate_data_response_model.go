// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCenterStreamRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveCenterStreamRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveCenterStreamRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveCenterStreamRateDataResponseBody) *DescribeLiveCenterStreamRateDataResponse
	GetBody() *DescribeLiveCenterStreamRateDataResponseBody
}

type DescribeLiveCenterStreamRateDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveCenterStreamRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveCenterStreamRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterStreamRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterStreamRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveCenterStreamRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveCenterStreamRateDataResponse) GetBody() *DescribeLiveCenterStreamRateDataResponseBody {
	return s.Body
}

func (s *DescribeLiveCenterStreamRateDataResponse) SetHeaders(v map[string]*string) *DescribeLiveCenterStreamRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponse) SetStatusCode(v int32) *DescribeLiveCenterStreamRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponse) SetBody(v *DescribeLiveCenterStreamRateDataResponseBody) *DescribeLiveCenterStreamRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveCenterStreamRateDataResponse) Validate() error {
	return dara.Validate(s)
}
