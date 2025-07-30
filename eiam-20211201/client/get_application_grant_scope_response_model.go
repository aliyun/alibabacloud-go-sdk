// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationGrantScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationGrantScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationGrantScopeResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationGrantScopeResponseBody) *GetApplicationGrantScopeResponse
	GetBody() *GetApplicationGrantScopeResponseBody
}

type GetApplicationGrantScopeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationGrantScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationGrantScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGrantScopeResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationGrantScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationGrantScopeResponse) GetBody() *GetApplicationGrantScopeResponseBody {
	return s.Body
}

func (s *GetApplicationGrantScopeResponse) SetHeaders(v map[string]*string) *GetApplicationGrantScopeResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationGrantScopeResponse) SetStatusCode(v int32) *GetApplicationGrantScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationGrantScopeResponse) SetBody(v *GetApplicationGrantScopeResponseBody) *GetApplicationGrantScopeResponse {
	s.Body = v
	return s
}

func (s *GetApplicationGrantScopeResponse) Validate() error {
	return dara.Validate(s)
}
