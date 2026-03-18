// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosOriginInstanceBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosOriginInstanceBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosOriginInstanceBillResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosOriginInstanceBillResponseBody) *DescribeDdosOriginInstanceBillResponse
	GetBody() *DescribeDdosOriginInstanceBillResponseBody
}

type DescribeDdosOriginInstanceBillResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosOriginInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosOriginInstanceBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosOriginInstanceBillResponse) GetBody() *DescribeDdosOriginInstanceBillResponseBody {
	return s.Body
}

func (s *DescribeDdosOriginInstanceBillResponse) SetHeaders(v map[string]*string) *DescribeDdosOriginInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponse) SetStatusCode(v int32) *DescribeDdosOriginInstanceBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponse) SetBody(v *DescribeDdosOriginInstanceBillResponseBody) *DescribeDdosOriginInstanceBillResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
