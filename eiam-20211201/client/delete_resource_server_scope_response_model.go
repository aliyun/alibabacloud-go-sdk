// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceServerScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceServerScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceServerScopeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceServerScopeResponseBody) *DeleteResourceServerScopeResponse
	GetBody() *DeleteResourceServerScopeResponseBody
}

type DeleteResourceServerScopeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceServerScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceServerScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceServerScopeResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceServerScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceServerScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceServerScopeResponse) GetBody() *DeleteResourceServerScopeResponseBody {
	return s.Body
}

func (s *DeleteResourceServerScopeResponse) SetHeaders(v map[string]*string) *DeleteResourceServerScopeResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceServerScopeResponse) SetStatusCode(v int32) *DeleteResourceServerScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceServerScopeResponse) SetBody(v *DeleteResourceServerScopeResponseBody) *DeleteResourceServerScopeResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceServerScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
