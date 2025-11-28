// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScanRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScanRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScanRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScanRuleResponseBody) *DeleteScanRuleResponse
	GetBody() *DeleteScanRuleResponseBody
}

type DeleteScanRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScanRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScanRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScanRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScanRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScanRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScanRuleResponse) GetBody() *DeleteScanRuleResponseBody {
	return s.Body
}

func (s *DeleteScanRuleResponse) SetHeaders(v map[string]*string) *DeleteScanRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScanRuleResponse) SetStatusCode(v int32) *DeleteScanRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScanRuleResponse) SetBody(v *DeleteScanRuleResponseBody) *DeleteScanRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteScanRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
