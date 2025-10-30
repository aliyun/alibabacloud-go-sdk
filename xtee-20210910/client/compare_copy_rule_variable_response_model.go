// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareCopyRuleVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareCopyRuleVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareCopyRuleVariableResponse
	GetStatusCode() *int32
	SetBody(v *CompareCopyRuleVariableResponseBody) *CompareCopyRuleVariableResponse
	GetBody() *CompareCopyRuleVariableResponseBody
}

type CompareCopyRuleVariableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareCopyRuleVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareCopyRuleVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponse) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareCopyRuleVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareCopyRuleVariableResponse) GetBody() *CompareCopyRuleVariableResponseBody {
	return s.Body
}

func (s *CompareCopyRuleVariableResponse) SetHeaders(v map[string]*string) *CompareCopyRuleVariableResponse {
	s.Headers = v
	return s
}

func (s *CompareCopyRuleVariableResponse) SetStatusCode(v int32) *CompareCopyRuleVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareCopyRuleVariableResponse) SetBody(v *CompareCopyRuleVariableResponseBody) *CompareCopyRuleVariableResponse {
	s.Body = v
	return s
}

func (s *CompareCopyRuleVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
