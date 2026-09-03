// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrinterEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrinterEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrinterEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrinterEventsResponseBody) *DescribePrinterEventsResponse
	GetBody() *DescribePrinterEventsResponseBody
}

type DescribePrinterEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrinterEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrinterEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrinterEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribePrinterEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrinterEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrinterEventsResponse) GetBody() *DescribePrinterEventsResponseBody {
	return s.Body
}

func (s *DescribePrinterEventsResponse) SetHeaders(v map[string]*string) *DescribePrinterEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribePrinterEventsResponse) SetStatusCode(v int32) *DescribePrinterEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrinterEventsResponse) SetBody(v *DescribePrinterEventsResponseBody) *DescribePrinterEventsResponse {
	s.Body = v
	return s
}

func (s *DescribePrinterEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
