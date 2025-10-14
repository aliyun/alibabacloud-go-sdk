// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdcInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdcInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdcInfoResponseBody) *DescribeCdcInfoResponse
	GetBody() *DescribeCdcInfoResponseBody
}

type DescribeCdcInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdcInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdcInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdcInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdcInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdcInfoResponse) GetBody() *DescribeCdcInfoResponseBody {
	return s.Body
}

func (s *DescribeCdcInfoResponse) SetHeaders(v map[string]*string) *DescribeCdcInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdcInfoResponse) SetStatusCode(v int32) *DescribeCdcInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdcInfoResponse) SetBody(v *DescribeCdcInfoResponseBody) *DescribeCdcInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCdcInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
