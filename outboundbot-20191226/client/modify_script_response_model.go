// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScriptResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScriptResponseBody) *ModifyScriptResponse
	GetBody() *ModifyScriptResponseBody
}

type ModifyScriptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptResponse) GoString() string {
	return s.String()
}

func (s *ModifyScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScriptResponse) GetBody() *ModifyScriptResponseBody {
	return s.Body
}

func (s *ModifyScriptResponse) SetHeaders(v map[string]*string) *ModifyScriptResponse {
	s.Headers = v
	return s
}

func (s *ModifyScriptResponse) SetStatusCode(v int32) *ModifyScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScriptResponse) SetBody(v *ModifyScriptResponseBody) *ModifyScriptResponse {
	s.Body = v
	return s
}

func (s *ModifyScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
