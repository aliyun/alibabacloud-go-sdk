// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyProjectAuthorizationResponseBody) *ModifyProjectAuthorizationResponse
	GetBody() *ModifyProjectAuthorizationResponseBody
}

type ModifyProjectAuthorizationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ModifyProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyProjectAuthorizationResponse) GetBody() *ModifyProjectAuthorizationResponseBody {
	return s.Body
}

func (s *ModifyProjectAuthorizationResponse) SetHeaders(v map[string]*string) *ModifyProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ModifyProjectAuthorizationResponse) SetStatusCode(v int32) *ModifyProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProjectAuthorizationResponse) SetBody(v *ModifyProjectAuthorizationResponseBody) *ModifyProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ModifyProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
