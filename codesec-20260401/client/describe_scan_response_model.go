// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScanResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScanResponseBody) *DescribeScanResponse
	GetBody() *DescribeScanResponseBody
}

type DescribeScanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScanResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScanResponse) GetBody() *DescribeScanResponseBody {
	return s.Body
}

func (s *DescribeScanResponse) SetHeaders(v map[string]*string) *DescribeScanResponse {
	s.Headers = v
	return s
}

func (s *DescribeScanResponse) SetStatusCode(v int32) *DescribeScanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScanResponse) SetBody(v *DescribeScanResponseBody) *DescribeScanResponse {
	s.Body = v
	return s
}

func (s *DescribeScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
