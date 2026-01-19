// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceServerScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceServerScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceServerScopeResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceServerScopeResponseBody) *GetResourceServerScopeResponse
	GetBody() *GetResourceServerScopeResponseBody
}

type GetResourceServerScopeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceServerScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceServerScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceServerScopeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceServerScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceServerScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceServerScopeResponse) GetBody() *GetResourceServerScopeResponseBody {
	return s.Body
}

func (s *GetResourceServerScopeResponse) SetHeaders(v map[string]*string) *GetResourceServerScopeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceServerScopeResponse) SetStatusCode(v int32) *GetResourceServerScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceServerScopeResponse) SetBody(v *GetResourceServerScopeResponseBody) *GetResourceServerScopeResponse {
	s.Body = v
	return s
}

func (s *GetResourceServerScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
