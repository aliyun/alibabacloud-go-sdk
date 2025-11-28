// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberBizChainStatInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMemberBizChainStatInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMemberBizChainStatInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMemberBizChainStatInfoResponseBody) *DescribeMemberBizChainStatInfoResponse
	GetBody() *DescribeMemberBizChainStatInfoResponseBody
}

type DescribeMemberBizChainStatInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMemberBizChainStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMemberBizChainStatInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberBizChainStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMemberBizChainStatInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMemberBizChainStatInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMemberBizChainStatInfoResponse) GetBody() *DescribeMemberBizChainStatInfoResponseBody {
	return s.Body
}

func (s *DescribeMemberBizChainStatInfoResponse) SetHeaders(v map[string]*string) *DescribeMemberBizChainStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponse) SetStatusCode(v int32) *DescribeMemberBizChainStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponse) SetBody(v *DescribeMemberBizChainStatInfoResponseBody) *DescribeMemberBizChainStatInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
