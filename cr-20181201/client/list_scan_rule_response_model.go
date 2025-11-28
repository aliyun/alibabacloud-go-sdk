// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScanRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScanRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListScanRuleResponseBody) *ListScanRuleResponse
	GetBody() *ListScanRuleResponseBody
}

type ListScanRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScanRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScanRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScanRuleResponse) GoString() string {
	return s.String()
}

func (s *ListScanRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScanRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScanRuleResponse) GetBody() *ListScanRuleResponseBody {
	return s.Body
}

func (s *ListScanRuleResponse) SetHeaders(v map[string]*string) *ListScanRuleResponse {
	s.Headers = v
	return s
}

func (s *ListScanRuleResponse) SetStatusCode(v int32) *ListScanRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScanRuleResponse) SetBody(v *ListScanRuleResponseBody) *ListScanRuleResponse {
	s.Body = v
	return s
}

func (s *ListScanRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
