// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetricRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetricRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetricRulesResponseBody) *DeleteMetricRulesResponse
	GetBody() *DeleteMetricRulesResponseBody
}

type DeleteMetricRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetricRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetricRulesResponse) GetBody() *DeleteMetricRulesResponseBody {
	return s.Body
}

func (s *DeleteMetricRulesResponse) SetHeaders(v map[string]*string) *DeleteMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRulesResponse) SetStatusCode(v int32) *DeleteMetricRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricRulesResponse) SetBody(v *DeleteMetricRulesResponseBody) *DeleteMetricRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteMetricRulesResponse) Validate() error {
	return dara.Validate(s)
}
