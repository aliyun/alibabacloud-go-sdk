// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScanRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScanRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetScanRuleResponseBody) *GetScanRuleResponse
	GetBody() *GetScanRuleResponseBody
}

type GetScanRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScanRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScanRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScanRuleResponse) GoString() string {
	return s.String()
}

func (s *GetScanRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScanRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScanRuleResponse) GetBody() *GetScanRuleResponseBody {
	return s.Body
}

func (s *GetScanRuleResponse) SetHeaders(v map[string]*string) *GetScanRuleResponse {
	s.Headers = v
	return s
}

func (s *GetScanRuleResponse) SetStatusCode(v int32) *GetScanRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScanRuleResponse) SetBody(v *GetScanRuleResponseBody) *GetScanRuleResponse {
	s.Body = v
	return s
}

func (s *GetScanRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
