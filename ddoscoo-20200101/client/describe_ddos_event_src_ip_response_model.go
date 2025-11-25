// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventSrcIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDosEventSrcIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDosEventSrcIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDosEventSrcIpResponseBody) *DescribeDDosEventSrcIpResponse
	GetBody() *DescribeDDosEventSrcIpResponseBody
}

type DescribeDDosEventSrcIpResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDosEventSrcIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDosEventSrcIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventSrcIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDosEventSrcIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDosEventSrcIpResponse) GetBody() *DescribeDDosEventSrcIpResponseBody {
	return s.Body
}

func (s *DescribeDDosEventSrcIpResponse) SetHeaders(v map[string]*string) *DescribeDDosEventSrcIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventSrcIpResponse) SetStatusCode(v int32) *DescribeDDosEventSrcIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponse) SetBody(v *DescribeDDosEventSrcIpResponseBody) *DescribeDDosEventSrcIpResponse {
	s.Body = v
	return s
}

func (s *DescribeDDosEventSrcIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
