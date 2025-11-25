// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnLinkageRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnLinkageRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnLinkageRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnLinkageRulesResponseBody) *DescribeCdnLinkageRulesResponse
	GetBody() *DescribeCdnLinkageRulesResponseBody
}

type DescribeCdnLinkageRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnLinkageRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnLinkageRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnLinkageRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnLinkageRulesResponse) GetBody() *DescribeCdnLinkageRulesResponseBody {
	return s.Body
}

func (s *DescribeCdnLinkageRulesResponse) SetHeaders(v map[string]*string) *DescribeCdnLinkageRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnLinkageRulesResponse) SetStatusCode(v int32) *DescribeCdnLinkageRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnLinkageRulesResponse) SetBody(v *DescribeCdnLinkageRulesResponseBody) *DescribeCdnLinkageRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnLinkageRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
