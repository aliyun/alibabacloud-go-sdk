// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarRecordsResponseBody) *DescribeSoarRecordsResponse
	GetBody() *DescribeSoarRecordsResponseBody
}

type DescribeSoarRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarRecordsResponse) GetBody() *DescribeSoarRecordsResponseBody {
	return s.Body
}

func (s *DescribeSoarRecordsResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordsResponse) SetStatusCode(v int32) *DescribeSoarRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordsResponse) SetBody(v *DescribeSoarRecordsResponseBody) *DescribeSoarRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
