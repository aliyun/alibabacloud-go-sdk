// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutResourceMetricRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutResourceMetricRulesResponse
	GetStatusCode() *int32
	SetBody(v *PutResourceMetricRulesResponseBody) *PutResourceMetricRulesResponse
	GetBody() *PutResourceMetricRulesResponseBody
}

type PutResourceMetricRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutResourceMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutResourceMetricRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutResourceMetricRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutResourceMetricRulesResponse) GetBody() *PutResourceMetricRulesResponseBody {
	return s.Body
}

func (s *PutResourceMetricRulesResponse) SetHeaders(v map[string]*string) *PutResourceMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *PutResourceMetricRulesResponse) SetStatusCode(v int32) *PutResourceMetricRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *PutResourceMetricRulesResponse) SetBody(v *PutResourceMetricRulesResponseBody) *PutResourceMetricRulesResponse {
	s.Body = v
	return s
}

func (s *PutResourceMetricRulesResponse) Validate() error {
	return dara.Validate(s)
}
