// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceServerScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceServerScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceServerScopeResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceServerScopeResponseBody) *CreateResourceServerScopeResponse
	GetBody() *CreateResourceServerScopeResponseBody
}

type CreateResourceServerScopeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceServerScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceServerScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceServerScopeResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceServerScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceServerScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceServerScopeResponse) GetBody() *CreateResourceServerScopeResponseBody {
	return s.Body
}

func (s *CreateResourceServerScopeResponse) SetHeaders(v map[string]*string) *CreateResourceServerScopeResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceServerScopeResponse) SetStatusCode(v int32) *CreateResourceServerScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceServerScopeResponse) SetBody(v *CreateResourceServerScopeResponseBody) *CreateResourceServerScopeResponse {
	s.Body = v
	return s
}

func (s *CreateResourceServerScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
