// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupportRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupportRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupportRuleListResponseBody) *DescribeSupportRuleListResponse
	GetBody() *DescribeSupportRuleListResponseBody
}

type DescribeSupportRuleListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupportRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupportRuleListResponse) GetBody() *DescribeSupportRuleListResponseBody {
	return s.Body
}

func (s *DescribeSupportRuleListResponse) SetHeaders(v map[string]*string) *DescribeSupportRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportRuleListResponse) SetStatusCode(v int32) *DescribeSupportRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportRuleListResponse) SetBody(v *DescribeSupportRuleListResponseBody) *DescribeSupportRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeSupportRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
