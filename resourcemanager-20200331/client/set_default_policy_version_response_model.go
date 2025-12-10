// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultPolicyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultPolicyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultPolicyVersionResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultPolicyVersionResponseBody) *SetDefaultPolicyVersionResponse
	GetBody() *SetDefaultPolicyVersionResponseBody
}

type SetDefaultPolicyVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultPolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultPolicyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultPolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultPolicyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultPolicyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultPolicyVersionResponse) GetBody() *SetDefaultPolicyVersionResponseBody {
	return s.Body
}

func (s *SetDefaultPolicyVersionResponse) SetHeaders(v map[string]*string) *SetDefaultPolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultPolicyVersionResponse) SetStatusCode(v int32) *SetDefaultPolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultPolicyVersionResponse) SetBody(v *SetDefaultPolicyVersionResponseBody) *SetDefaultPolicyVersionResponse {
	s.Body = v
	return s
}

func (s *SetDefaultPolicyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
