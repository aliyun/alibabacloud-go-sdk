// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleInfoResponseBody) *DescribeSampleInfoResponse
	GetBody() *DescribeSampleInfoResponseBody
}

type DescribeSampleInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleInfoResponse) GetBody() *DescribeSampleInfoResponseBody {
	return s.Body
}

func (s *DescribeSampleInfoResponse) SetHeaders(v map[string]*string) *DescribeSampleInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleInfoResponse) SetStatusCode(v int32) *DescribeSampleInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleInfoResponse) SetBody(v *DescribeSampleInfoResponseBody) *DescribeSampleInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
