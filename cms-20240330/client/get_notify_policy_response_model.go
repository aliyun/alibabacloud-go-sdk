// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNotifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNotifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetNotifyPolicyResponseBody) *GetNotifyPolicyResponse
	GetBody() *GetNotifyPolicyResponseBody
}

type GetNotifyPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNotifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetNotifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNotifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNotifyPolicyResponse) GetBody() *GetNotifyPolicyResponseBody {
	return s.Body
}

func (s *GetNotifyPolicyResponse) SetHeaders(v map[string]*string) *GetNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetNotifyPolicyResponse) SetStatusCode(v int32) *GetNotifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotifyPolicyResponse) SetBody(v *GetNotifyPolicyResponseBody) *GetNotifyPolicyResponse {
	s.Body = v
	return s
}

func (s *GetNotifyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
