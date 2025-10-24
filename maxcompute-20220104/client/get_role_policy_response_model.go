// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRolePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRolePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRolePolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetRolePolicyResponseBody) *GetRolePolicyResponse
	GetBody() *GetRolePolicyResponseBody
}

type GetRolePolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRolePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRolePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRolePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetRolePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRolePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRolePolicyResponse) GetBody() *GetRolePolicyResponseBody {
	return s.Body
}

func (s *GetRolePolicyResponse) SetHeaders(v map[string]*string) *GetRolePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetRolePolicyResponse) SetStatusCode(v int32) *GetRolePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRolePolicyResponse) SetBody(v *GetRolePolicyResponseBody) *GetRolePolicyResponse {
	s.Body = v
	return s
}

func (s *GetRolePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
