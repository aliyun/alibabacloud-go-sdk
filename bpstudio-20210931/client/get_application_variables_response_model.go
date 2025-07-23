// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationVariablesResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationVariablesResponseBody) *GetApplicationVariablesResponse
	GetBody() *GetApplicationVariablesResponseBody
}

type GetApplicationVariablesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariablesResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationVariablesResponse) GetBody() *GetApplicationVariablesResponseBody {
	return s.Body
}

func (s *GetApplicationVariablesResponse) SetHeaders(v map[string]*string) *GetApplicationVariablesResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationVariablesResponse) SetStatusCode(v int32) *GetApplicationVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationVariablesResponse) SetBody(v *GetApplicationVariablesResponseBody) *GetApplicationVariablesResponse {
	s.Body = v
	return s
}

func (s *GetApplicationVariablesResponse) Validate() error {
	return dara.Validate(s)
}
