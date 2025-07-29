// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptPlanningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunScriptPlanningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunScriptPlanningResponse
	GetStatusCode() *int32
	SetBody(v *RunScriptPlanningResponseBody) *RunScriptPlanningResponse
	GetBody() *RunScriptPlanningResponseBody
}

type RunScriptPlanningResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunScriptPlanningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptPlanningResponse) String() string {
	return dara.Prettify(s)
}

func (s RunScriptPlanningResponse) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunScriptPlanningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunScriptPlanningResponse) GetBody() *RunScriptPlanningResponseBody {
	return s.Body
}

func (s *RunScriptPlanningResponse) SetHeaders(v map[string]*string) *RunScriptPlanningResponse {
	s.Headers = v
	return s
}

func (s *RunScriptPlanningResponse) SetStatusCode(v int32) *RunScriptPlanningResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptPlanningResponse) SetBody(v *RunScriptPlanningResponseBody) *RunScriptPlanningResponse {
	s.Body = v
	return s
}

func (s *RunScriptPlanningResponse) Validate() error {
	return dara.Validate(s)
}
