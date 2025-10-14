// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackIntelligentProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackIntelligentProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHttpDDoSAttackIntelligentProtectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) *DescribeHttpDDoSAttackIntelligentProtectionResponse
	GetBody() *DescribeHttpDDoSAttackIntelligentProtectionResponseBody
}

type DescribeHttpDDoSAttackIntelligentProtectionResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHttpDDoSAttackIntelligentProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) GetBody() *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	return s.Body
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackIntelligentProtectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) SetStatusCode(v int32) *DescribeHttpDDoSAttackIntelligentProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) SetBody(v *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) *DescribeHttpDDoSAttackIntelligentProtectionResponse {
	s.Body = v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
