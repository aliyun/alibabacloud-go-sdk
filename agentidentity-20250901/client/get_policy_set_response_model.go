// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicySetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicySetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicySetResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicySetResponseBody) *GetPolicySetResponse
	GetBody() *GetPolicySetResponseBody
}

type GetPolicySetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicySetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicySetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicySetResponse) GoString() string {
	return s.String()
}

func (s *GetPolicySetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicySetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicySetResponse) GetBody() *GetPolicySetResponseBody {
	return s.Body
}

func (s *GetPolicySetResponse) SetHeaders(v map[string]*string) *GetPolicySetResponse {
	s.Headers = v
	return s
}

func (s *GetPolicySetResponse) SetStatusCode(v int32) *GetPolicySetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicySetResponse) SetBody(v *GetPolicySetResponseBody) *GetPolicySetResponse {
	s.Body = v
	return s
}

func (s *GetPolicySetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
