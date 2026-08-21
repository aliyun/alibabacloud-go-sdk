// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProhibitedPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProhibitedPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProhibitedPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetProhibitedPolicyResponseBody) *GetProhibitedPolicyResponse
	GetBody() *GetProhibitedPolicyResponseBody
}

type GetProhibitedPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProhibitedPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProhibitedPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetProhibitedPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProhibitedPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProhibitedPolicyResponse) GetBody() *GetProhibitedPolicyResponseBody {
	return s.Body
}

func (s *GetProhibitedPolicyResponse) SetHeaders(v map[string]*string) *GetProhibitedPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetProhibitedPolicyResponse) SetStatusCode(v int32) *GetProhibitedPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProhibitedPolicyResponse) SetBody(v *GetProhibitedPolicyResponseBody) *GetProhibitedPolicyResponse {
	s.Body = v
	return s
}

func (s *GetProhibitedPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
