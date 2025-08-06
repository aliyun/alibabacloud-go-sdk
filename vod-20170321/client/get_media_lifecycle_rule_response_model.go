// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaLifecycleRuleResponseBody) *GetMediaLifecycleRuleResponse
	GetBody() *GetMediaLifecycleRuleResponseBody
}

type GetMediaLifecycleRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *GetMediaLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaLifecycleRuleResponse) GetBody() *GetMediaLifecycleRuleResponseBody {
	return s.Body
}

func (s *GetMediaLifecycleRuleResponse) SetHeaders(v map[string]*string) *GetMediaLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *GetMediaLifecycleRuleResponse) SetStatusCode(v int32) *GetMediaLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaLifecycleRuleResponse) SetBody(v *GetMediaLifecycleRuleResponseBody) *GetMediaLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *GetMediaLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
