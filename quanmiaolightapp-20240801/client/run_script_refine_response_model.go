// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptRefineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunScriptRefineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunScriptRefineResponse
	GetStatusCode() *int32
	SetId(v string) *RunScriptRefineResponse
	GetId() *string
	SetEvent(v string) *RunScriptRefineResponse
	GetEvent() *string
	SetBody(v *RunScriptRefineResponseBody) *RunScriptRefineResponse
	GetBody() *RunScriptRefineResponseBody
}

type RunScriptRefineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                      `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                      `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunScriptRefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptRefineResponse) String() string {
	return dara.Prettify(s)
}

func (s RunScriptRefineResponse) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunScriptRefineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunScriptRefineResponse) GetId() *string {
	return s.Id
}

func (s *RunScriptRefineResponse) GetEvent() *string {
	return s.Event
}

func (s *RunScriptRefineResponse) GetBody() *RunScriptRefineResponseBody {
	return s.Body
}

func (s *RunScriptRefineResponse) SetHeaders(v map[string]*string) *RunScriptRefineResponse {
	s.Headers = v
	return s
}

func (s *RunScriptRefineResponse) SetStatusCode(v int32) *RunScriptRefineResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptRefineResponse) SetId(v string) *RunScriptRefineResponse {
	s.Id = &v
	return s
}

func (s *RunScriptRefineResponse) SetEvent(v string) *RunScriptRefineResponse {
	s.Event = &v
	return s
}

func (s *RunScriptRefineResponse) SetBody(v *RunScriptRefineResponseBody) *RunScriptRefineResponse {
	s.Body = v
	return s
}

func (s *RunScriptRefineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
