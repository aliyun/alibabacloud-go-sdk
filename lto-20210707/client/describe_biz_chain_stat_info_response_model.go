// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBizChainStatInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBizChainStatInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBizChainStatInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBizChainStatInfoResponseBody) *DescribeBizChainStatInfoResponse
	GetBody() *DescribeBizChainStatInfoResponseBody
}

type DescribeBizChainStatInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBizChainStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBizChainStatInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizChainStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizChainStatInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBizChainStatInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBizChainStatInfoResponse) GetBody() *DescribeBizChainStatInfoResponseBody {
	return s.Body
}

func (s *DescribeBizChainStatInfoResponse) SetHeaders(v map[string]*string) *DescribeBizChainStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizChainStatInfoResponse) SetStatusCode(v int32) *DescribeBizChainStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBizChainStatInfoResponse) SetBody(v *DescribeBizChainStatInfoResponseBody) *DescribeBizChainStatInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeBizChainStatInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
