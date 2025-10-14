// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExporterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExporterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExporterRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExporterRuleResponseBody) *DeleteExporterRuleResponse
	GetBody() *DeleteExporterRuleResponseBody
}

type DeleteExporterRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExporterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExporterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExporterRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteExporterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExporterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExporterRuleResponse) GetBody() *DeleteExporterRuleResponseBody {
	return s.Body
}

func (s *DeleteExporterRuleResponse) SetHeaders(v map[string]*string) *DeleteExporterRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteExporterRuleResponse) SetStatusCode(v int32) *DeleteExporterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExporterRuleResponse) SetBody(v *DeleteExporterRuleResponseBody) *DeleteExporterRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteExporterRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
