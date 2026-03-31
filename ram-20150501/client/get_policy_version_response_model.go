// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyVersionResponseBody) *GetPolicyVersionResponse
	GetBody() *GetPolicyVersionResponseBody
}

type GetPolicyVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyVersionResponse) GetBody() *GetPolicyVersionResponseBody {
	return s.Body
}

func (s *GetPolicyVersionResponse) SetHeaders(v map[string]*string) *GetPolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyVersionResponse) SetStatusCode(v int32) *GetPolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyVersionResponse) SetBody(v *GetPolicyVersionResponseBody) *GetPolicyVersionResponse {
	s.Body = v
	return s
}

func (s *GetPolicyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
