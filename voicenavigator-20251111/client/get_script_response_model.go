// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScriptResponse
	GetStatusCode() *int32
	SetBody(v *GetScriptResponseBody) *GetScriptResponse
	GetBody() *GetScriptResponseBody
}

type GetScriptResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScriptResponse) GoString() string {
	return s.String()
}

func (s *GetScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScriptResponse) GetBody() *GetScriptResponseBody {
	return s.Body
}

func (s *GetScriptResponse) SetHeaders(v map[string]*string) *GetScriptResponse {
	s.Headers = v
	return s
}

func (s *GetScriptResponse) SetStatusCode(v int32) *GetScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScriptResponse) SetBody(v *GetScriptResponseBody) *GetScriptResponse {
	s.Body = v
	return s
}

func (s *GetScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
