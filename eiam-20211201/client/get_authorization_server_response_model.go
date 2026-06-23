// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthorizationServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthorizationServerResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthorizationServerResponseBody) *GetAuthorizationServerResponse
	GetBody() *GetAuthorizationServerResponseBody
}

type GetAuthorizationServerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorizationServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorizationServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationServerResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthorizationServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthorizationServerResponse) GetBody() *GetAuthorizationServerResponseBody {
	return s.Body
}

func (s *GetAuthorizationServerResponse) SetHeaders(v map[string]*string) *GetAuthorizationServerResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationServerResponse) SetStatusCode(v int32) *GetAuthorizationServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorizationServerResponse) SetBody(v *GetAuthorizationServerResponseBody) *GetAuthorizationServerResponse {
	s.Body = v
	return s
}

func (s *GetAuthorizationServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
