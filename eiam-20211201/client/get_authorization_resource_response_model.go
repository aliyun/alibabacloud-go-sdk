// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthorizationResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthorizationResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthorizationResourceResponseBody) *GetAuthorizationResourceResponse
	GetBody() *GetAuthorizationResourceResponseBody
}

type GetAuthorizationResourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorizationResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorizationResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationResourceResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthorizationResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthorizationResourceResponse) GetBody() *GetAuthorizationResourceResponseBody {
	return s.Body
}

func (s *GetAuthorizationResourceResponse) SetHeaders(v map[string]*string) *GetAuthorizationResourceResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationResourceResponse) SetStatusCode(v int32) *GetAuthorizationResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorizationResourceResponse) SetBody(v *GetAuthorizationResourceResponseBody) *GetAuthorizationResourceResponse {
	s.Body = v
	return s
}

func (s *GetAuthorizationResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
