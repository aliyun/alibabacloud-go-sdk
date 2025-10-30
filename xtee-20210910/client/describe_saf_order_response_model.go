// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSafOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSafOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSafOrderResponseBody) *DescribeSafOrderResponse
	GetBody() *DescribeSafOrderResponseBody
}

type DescribeSafOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSafOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSafOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeSafOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSafOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSafOrderResponse) GetBody() *DescribeSafOrderResponseBody {
	return s.Body
}

func (s *DescribeSafOrderResponse) SetHeaders(v map[string]*string) *DescribeSafOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeSafOrderResponse) SetStatusCode(v int32) *DescribeSafOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSafOrderResponse) SetBody(v *DescribeSafOrderResponseBody) *DescribeSafOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeSafOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
