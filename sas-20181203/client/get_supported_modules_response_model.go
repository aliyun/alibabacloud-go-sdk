// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupportedModulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupportedModulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupportedModulesResponse
	GetStatusCode() *int32
	SetBody(v *GetSupportedModulesResponseBody) *GetSupportedModulesResponse
	GetBody() *GetSupportedModulesResponseBody
}

type GetSupportedModulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupportedModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupportedModulesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedModulesResponse) GoString() string {
	return s.String()
}

func (s *GetSupportedModulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupportedModulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupportedModulesResponse) GetBody() *GetSupportedModulesResponseBody {
	return s.Body
}

func (s *GetSupportedModulesResponse) SetHeaders(v map[string]*string) *GetSupportedModulesResponse {
	s.Headers = v
	return s
}

func (s *GetSupportedModulesResponse) SetStatusCode(v int32) *GetSupportedModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupportedModulesResponse) SetBody(v *GetSupportedModulesResponseBody) *GetSupportedModulesResponse {
	s.Body = v
	return s
}

func (s *GetSupportedModulesResponse) Validate() error {
	return dara.Validate(s)
}
