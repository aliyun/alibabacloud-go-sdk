// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNotifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNotifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNotifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateNotifyPolicyResponseBody) *CreateNotifyPolicyResponse
	GetBody() *CreateNotifyPolicyResponseBody
}

type CreateNotifyPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNotifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateNotifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNotifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNotifyPolicyResponse) GetBody() *CreateNotifyPolicyResponseBody {
	return s.Body
}

func (s *CreateNotifyPolicyResponse) SetHeaders(v map[string]*string) *CreateNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateNotifyPolicyResponse) SetStatusCode(v int32) *CreateNotifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNotifyPolicyResponse) SetBody(v *CreateNotifyPolicyResponseBody) *CreateNotifyPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateNotifyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
