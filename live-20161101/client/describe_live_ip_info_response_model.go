// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveIpInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveIpInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveIpInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveIpInfoResponseBody) *DescribeLiveIpInfoResponse
	GetBody() *DescribeLiveIpInfoResponseBody
}

type DescribeLiveIpInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveIpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveIpInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveIpInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveIpInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveIpInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveIpInfoResponse) GetBody() *DescribeLiveIpInfoResponseBody {
	return s.Body
}

func (s *DescribeLiveIpInfoResponse) SetHeaders(v map[string]*string) *DescribeLiveIpInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveIpInfoResponse) SetStatusCode(v int32) *DescribeLiveIpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveIpInfoResponse) SetBody(v *DescribeLiveIpInfoResponseBody) *DescribeLiveIpInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveIpInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
