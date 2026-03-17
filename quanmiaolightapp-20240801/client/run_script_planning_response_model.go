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
	SetId(v string) *RunScriptPlanningResponse
	GetId() *string
	SetEvent(v string) *RunScriptPlanningResponse
	GetEvent() *string
	SetBody(v *RunScriptPlanningResponseBody) *RunScriptPlanningResponse
	GetBody() *RunScriptPlanningResponseBody
}

type RunScriptPlanningResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                        `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                        `json:"event,omitempty" xml:"event,omitempty"`
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

func (s *RunScriptPlanningResponse) GetId() *string {
	return s.Id
}

func (s *RunScriptPlanningResponse) GetEvent() *string {
	return s.Event
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

func (s *RunScriptPlanningResponse) SetId(v string) *RunScriptPlanningResponse {
	s.Id = &v
	return s
}

func (s *RunScriptPlanningResponse) SetEvent(v string) *RunScriptPlanningResponse {
	s.Event = &v
	return s
}

func (s *RunScriptPlanningResponse) SetBody(v *RunScriptPlanningResponseBody) *RunScriptPlanningResponse {
	s.Body = v
	return s
}

func (s *RunScriptPlanningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
