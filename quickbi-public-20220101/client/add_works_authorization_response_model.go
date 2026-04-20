// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorksAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorksAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorksAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *AddWorksAuthorizationResponseBody) *AddWorksAuthorizationResponse
	GetBody() *AddWorksAuthorizationResponseBody
}

type AddWorksAuthorizationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorksAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorksAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorksAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *AddWorksAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorksAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorksAuthorizationResponse) GetBody() *AddWorksAuthorizationResponseBody {
	return s.Body
}

func (s *AddWorksAuthorizationResponse) SetHeaders(v map[string]*string) *AddWorksAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *AddWorksAuthorizationResponse) SetStatusCode(v int32) *AddWorksAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorksAuthorizationResponse) SetBody(v *AddWorksAuthorizationResponseBody) *AddWorksAuthorizationResponse {
	s.Body = v
	return s
}

func (s *AddWorksAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
