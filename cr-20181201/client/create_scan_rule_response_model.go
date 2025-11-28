// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScanRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScanRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateScanRuleResponseBody) *CreateScanRuleResponse
	GetBody() *CreateScanRuleResponseBody
}

type CreateScanRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScanRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScanRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScanRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateScanRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScanRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScanRuleResponse) GetBody() *CreateScanRuleResponseBody {
	return s.Body
}

func (s *CreateScanRuleResponse) SetHeaders(v map[string]*string) *CreateScanRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateScanRuleResponse) SetStatusCode(v int32) *CreateScanRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScanRuleResponse) SetBody(v *CreateScanRuleResponseBody) *CreateScanRuleResponse {
	s.Body = v
	return s
}

func (s *CreateScanRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
