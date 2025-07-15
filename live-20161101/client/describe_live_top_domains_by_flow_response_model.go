// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTopDomainsByFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveTopDomainsByFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveTopDomainsByFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveTopDomainsByFlowResponseBody) *DescribeLiveTopDomainsByFlowResponse
	GetBody() *DescribeLiveTopDomainsByFlowResponseBody
}

type DescribeLiveTopDomainsByFlowResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveTopDomainsByFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveTopDomainsByFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveTopDomainsByFlowResponse) GetBody() *DescribeLiveTopDomainsByFlowResponseBody {
	return s.Body
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeLiveTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetStatusCode(v int32) *DescribeLiveTopDomainsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetBody(v *DescribeLiveTopDomainsByFlowResponseBody) *DescribeLiveTopDomainsByFlowResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) Validate() error {
	return dara.Validate(s)
}
