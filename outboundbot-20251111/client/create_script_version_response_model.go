// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScriptVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScriptVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateScriptVersionResponseBody) *CreateScriptVersionResponse
	GetBody() *CreateScriptVersionResponseBody
}

type CreateScriptVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScriptVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScriptVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScriptVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScriptVersionResponse) GetBody() *CreateScriptVersionResponseBody {
	return s.Body
}

func (s *CreateScriptVersionResponse) SetHeaders(v map[string]*string) *CreateScriptVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateScriptVersionResponse) SetStatusCode(v int32) *CreateScriptVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScriptVersionResponse) SetBody(v *CreateScriptVersionResponseBody) *CreateScriptVersionResponse {
	s.Body = v
	return s
}

func (s *CreateScriptVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
