// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAuthorizationResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAuthorizationResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAuthorizationResourceResponseBody) *DeleteAuthorizationResourceResponse
	GetBody() *DeleteAuthorizationResourceResponseBody
}

type DeleteAuthorizationResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuthorizationResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuthorizationResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAuthorizationResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAuthorizationResourceResponse) GetBody() *DeleteAuthorizationResourceResponseBody {
	return s.Body
}

func (s *DeleteAuthorizationResourceResponse) SetHeaders(v map[string]*string) *DeleteAuthorizationResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthorizationResourceResponse) SetStatusCode(v int32) *DeleteAuthorizationResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthorizationResourceResponse) SetBody(v *DeleteAuthorizationResourceResponseBody) *DeleteAuthorizationResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteAuthorizationResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
