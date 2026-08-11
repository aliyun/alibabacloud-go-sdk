// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScriptVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScriptVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScriptVersionResponseBody) *UpdateScriptVersionResponse
	GetBody() *UpdateScriptVersionResponseBody
}

type UpdateScriptVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScriptVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScriptVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScriptVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScriptVersionResponse) GetBody() *UpdateScriptVersionResponseBody {
	return s.Body
}

func (s *UpdateScriptVersionResponse) SetHeaders(v map[string]*string) *UpdateScriptVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateScriptVersionResponse) SetStatusCode(v int32) *UpdateScriptVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScriptVersionResponse) SetBody(v *UpdateScriptVersionResponseBody) *UpdateScriptVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateScriptVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
