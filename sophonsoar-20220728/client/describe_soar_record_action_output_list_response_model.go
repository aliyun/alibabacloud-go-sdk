// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordActionOutputListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarRecordActionOutputListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarRecordActionOutputListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarRecordActionOutputListResponseBody) *DescribeSoarRecordActionOutputListResponse
	GetBody() *DescribeSoarRecordActionOutputListResponseBody
}

type DescribeSoarRecordActionOutputListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarRecordActionOutputListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarRecordActionOutputListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarRecordActionOutputListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarRecordActionOutputListResponse) GetBody() *DescribeSoarRecordActionOutputListResponseBody {
	return s.Body
}

func (s *DescribeSoarRecordActionOutputListResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordActionOutputListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponse) SetStatusCode(v int32) *DescribeSoarRecordActionOutputListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponse) SetBody(v *DescribeSoarRecordActionOutputListResponseBody) *DescribeSoarRecordActionOutputListResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
