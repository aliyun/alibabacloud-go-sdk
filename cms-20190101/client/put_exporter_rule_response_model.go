// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutExporterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutExporterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutExporterRuleResponse
	GetStatusCode() *int32
	SetBody(v *PutExporterRuleResponseBody) *PutExporterRuleResponse
	GetBody() *PutExporterRuleResponseBody
}

type PutExporterRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutExporterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutExporterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutExporterRuleResponse) GoString() string {
	return s.String()
}

func (s *PutExporterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutExporterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutExporterRuleResponse) GetBody() *PutExporterRuleResponseBody {
	return s.Body
}

func (s *PutExporterRuleResponse) SetHeaders(v map[string]*string) *PutExporterRuleResponse {
	s.Headers = v
	return s
}

func (s *PutExporterRuleResponse) SetStatusCode(v int32) *PutExporterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutExporterRuleResponse) SetBody(v *PutExporterRuleResponseBody) *PutExporterRuleResponse {
	s.Body = v
	return s
}

func (s *PutExporterRuleResponse) Validate() error {
	return dara.Validate(s)
}
