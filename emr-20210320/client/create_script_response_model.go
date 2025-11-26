// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScriptResponse
	GetStatusCode() *int32
	SetBody(v *CreateScriptResponseBody) *CreateScriptResponse
	GetBody() *CreateScriptResponseBody
}

type CreateScriptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptResponse) GoString() string {
	return s.String()
}

func (s *CreateScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScriptResponse) GetBody() *CreateScriptResponseBody {
	return s.Body
}

func (s *CreateScriptResponse) SetHeaders(v map[string]*string) *CreateScriptResponse {
	s.Headers = v
	return s
}

func (s *CreateScriptResponse) SetStatusCode(v int32) *CreateScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScriptResponse) SetBody(v *CreateScriptResponseBody) *CreateScriptResponse {
	s.Body = v
	return s
}

func (s *CreateScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
