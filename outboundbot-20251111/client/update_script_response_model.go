// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScriptResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScriptResponseBody) *UpdateScriptResponse
	GetBody() *UpdateScriptResponseBody
}

type UpdateScriptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptResponse) GoString() string {
	return s.String()
}

func (s *UpdateScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScriptResponse) GetBody() *UpdateScriptResponseBody {
	return s.Body
}

func (s *UpdateScriptResponse) SetHeaders(v map[string]*string) *UpdateScriptResponse {
	s.Headers = v
	return s
}

func (s *UpdateScriptResponse) SetStatusCode(v int32) *UpdateScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScriptResponse) SetBody(v *UpdateScriptResponseBody) *UpdateScriptResponse {
	s.Body = v
	return s
}

func (s *UpdateScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
