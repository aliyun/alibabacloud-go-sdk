// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommodityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommodityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommodityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommodityResponseBody) *DescribeCommodityResponse
	GetBody() *DescribeCommodityResponseBody
}

type DescribeCommodityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommodityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommodityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommodityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommodityResponse) GetBody() *DescribeCommodityResponseBody {
	return s.Body
}

func (s *DescribeCommodityResponse) SetHeaders(v map[string]*string) *DescribeCommodityResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommodityResponse) SetStatusCode(v int32) *DescribeCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommodityResponse) SetBody(v *DescribeCommodityResponseBody) *DescribeCommodityResponse {
	s.Body = v
	return s
}

func (s *DescribeCommodityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
