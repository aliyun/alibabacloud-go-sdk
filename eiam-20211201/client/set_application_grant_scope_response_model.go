// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationGrantScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApplicationGrantScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApplicationGrantScopeResponse
	GetStatusCode() *int32
	SetBody(v *SetApplicationGrantScopeResponseBody) *SetApplicationGrantScopeResponse
	GetBody() *SetApplicationGrantScopeResponseBody
}

type SetApplicationGrantScopeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationGrantScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationGrantScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationGrantScopeResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationGrantScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApplicationGrantScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApplicationGrantScopeResponse) GetBody() *SetApplicationGrantScopeResponseBody {
	return s.Body
}

func (s *SetApplicationGrantScopeResponse) SetHeaders(v map[string]*string) *SetApplicationGrantScopeResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationGrantScopeResponse) SetStatusCode(v int32) *SetApplicationGrantScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationGrantScopeResponse) SetBody(v *SetApplicationGrantScopeResponseBody) *SetApplicationGrantScopeResponse {
	s.Body = v
	return s
}

func (s *SetApplicationGrantScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
