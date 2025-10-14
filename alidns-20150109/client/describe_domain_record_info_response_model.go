// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRecordInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRecordInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRecordInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRecordInfoResponseBody) *DescribeDomainRecordInfoResponse
	GetBody() *DescribeDomainRecordInfoResponseBody
}

type DescribeDomainRecordInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRecordInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRecordInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRecordInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRecordInfoResponse) GetBody() *DescribeDomainRecordInfoResponseBody {
	return s.Body
}

func (s *DescribeDomainRecordInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainRecordInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRecordInfoResponse) SetStatusCode(v int32) *DescribeDomainRecordInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRecordInfoResponse) SetBody(v *DescribeDomainRecordInfoResponseBody) *DescribeDomainRecordInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRecordInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
