// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupMetricRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupMetricRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupMetricRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupMetricRulesResponseBody) *CreateGroupMetricRulesResponse
	GetBody() *CreateGroupMetricRulesResponseBody
}

type CreateGroupMetricRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupMetricRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupMetricRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupMetricRulesResponse) GetBody() *CreateGroupMetricRulesResponseBody {
	return s.Body
}

func (s *CreateGroupMetricRulesResponse) SetHeaders(v map[string]*string) *CreateGroupMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupMetricRulesResponse) SetStatusCode(v int32) *CreateGroupMetricRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupMetricRulesResponse) SetBody(v *CreateGroupMetricRulesResponseBody) *CreateGroupMetricRulesResponse {
	s.Body = v
	return s
}

func (s *CreateGroupMetricRulesResponse) Validate() error {
	return dara.Validate(s)
}
