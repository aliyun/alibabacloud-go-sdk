// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantSagVbrRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGrantSagVbrRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGrantSagVbrRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGrantSagVbrRulesResponseBody) *DescribeGrantSagVbrRulesResponse
	GetBody() *DescribeGrantSagVbrRulesResponseBody
}

type DescribeGrantSagVbrRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGrantSagVbrRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGrantSagVbrRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagVbrRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagVbrRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGrantSagVbrRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGrantSagVbrRulesResponse) GetBody() *DescribeGrantSagVbrRulesResponseBody {
	return s.Body
}

func (s *DescribeGrantSagVbrRulesResponse) SetHeaders(v map[string]*string) *DescribeGrantSagVbrRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGrantSagVbrRulesResponse) SetStatusCode(v int32) *DescribeGrantSagVbrRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGrantSagVbrRulesResponse) SetBody(v *DescribeGrantSagVbrRulesResponseBody) *DescribeGrantSagVbrRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeGrantSagVbrRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
