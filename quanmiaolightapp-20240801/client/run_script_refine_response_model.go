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
	SetBody(v *RunScriptRefineResponseBody) *RunScriptRefineResponse
	GetBody() *RunScriptRefineResponseBody
}

type RunScriptRefineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
