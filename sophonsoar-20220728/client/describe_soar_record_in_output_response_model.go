// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordInOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarRecordInOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarRecordInOutputResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarRecordInOutputResponseBody) *DescribeSoarRecordInOutputResponse
	GetBody() *DescribeSoarRecordInOutputResponseBody
}

type DescribeSoarRecordInOutputResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarRecordInOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarRecordInOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordInOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarRecordInOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarRecordInOutputResponse) GetBody() *DescribeSoarRecordInOutputResponseBody {
	return s.Body
}

func (s *DescribeSoarRecordInOutputResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordInOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordInOutputResponse) SetStatusCode(v int32) *DescribeSoarRecordInOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordInOutputResponse) SetBody(v *DescribeSoarRecordInOutputResponseBody) *DescribeSoarRecordInOutputResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarRecordInOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
