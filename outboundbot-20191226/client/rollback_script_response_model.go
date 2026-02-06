// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackScriptResponse
	GetStatusCode() *int32
	SetBody(v *RollbackScriptResponseBody) *RollbackScriptResponse
	GetBody() *RollbackScriptResponseBody
}

type RollbackScriptResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackScriptResponse) GoString() string {
	return s.String()
}

func (s *RollbackScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackScriptResponse) GetBody() *RollbackScriptResponseBody {
	return s.Body
}

func (s *RollbackScriptResponse) SetHeaders(v map[string]*string) *RollbackScriptResponse {
	s.Headers = v
	return s
}

func (s *RollbackScriptResponse) SetStatusCode(v int32) *RollbackScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackScriptResponse) SetBody(v *RollbackScriptResponseBody) *RollbackScriptResponse {
	s.Body = v
	return s
}

func (s *RollbackScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
