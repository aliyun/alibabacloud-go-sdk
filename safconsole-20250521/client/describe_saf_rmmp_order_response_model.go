// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafRmmpOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSafRmmpOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSafRmmpOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSafRmmpOrderResponseBody) *DescribeSafRmmpOrderResponse
	GetBody() *DescribeSafRmmpOrderResponseBody
}

type DescribeSafRmmpOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSafRmmpOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSafRmmpOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafRmmpOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeSafRmmpOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSafRmmpOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSafRmmpOrderResponse) GetBody() *DescribeSafRmmpOrderResponseBody {
	return s.Body
}

func (s *DescribeSafRmmpOrderResponse) SetHeaders(v map[string]*string) *DescribeSafRmmpOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeSafRmmpOrderResponse) SetStatusCode(v int32) *DescribeSafRmmpOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSafRmmpOrderResponse) SetBody(v *DescribeSafRmmpOrderResponseBody) *DescribeSafRmmpOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeSafRmmpOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
