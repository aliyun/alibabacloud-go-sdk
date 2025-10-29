// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndAttachPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndAttachPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndAttachPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndAttachPolicyResponseBody) *CreateAndAttachPolicyResponse
	GetBody() *CreateAndAttachPolicyResponseBody
}

type CreateAndAttachPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndAttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndAttachPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndAttachPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAndAttachPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndAttachPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndAttachPolicyResponse) GetBody() *CreateAndAttachPolicyResponseBody {
	return s.Body
}

func (s *CreateAndAttachPolicyResponse) SetHeaders(v map[string]*string) *CreateAndAttachPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAndAttachPolicyResponse) SetStatusCode(v int32) *CreateAndAttachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndAttachPolicyResponse) SetBody(v *CreateAndAttachPolicyResponseBody) *CreateAndAttachPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateAndAttachPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
