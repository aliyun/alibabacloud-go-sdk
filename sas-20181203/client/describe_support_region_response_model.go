// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupportRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupportRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupportRegionResponseBody) *DescribeSupportRegionResponse
	GetBody() *DescribeSupportRegionResponseBody
}

type DescribeSupportRegionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupportRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupportRegionResponse) GetBody() *DescribeSupportRegionResponseBody {
	return s.Body
}

func (s *DescribeSupportRegionResponse) SetHeaders(v map[string]*string) *DescribeSupportRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportRegionResponse) SetStatusCode(v int32) *DescribeSupportRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportRegionResponse) SetBody(v *DescribeSupportRegionResponseBody) *DescribeSupportRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeSupportRegionResponse) Validate() error {
	return dara.Validate(s)
}
