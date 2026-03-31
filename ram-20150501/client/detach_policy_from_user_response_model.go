// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachPolicyFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachPolicyFromUserResponse
	GetStatusCode() *int32
	SetBody(v *DetachPolicyFromUserResponseBody) *DetachPolicyFromUserResponse
	GetBody() *DetachPolicyFromUserResponseBody
}

type DetachPolicyFromUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromUserResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachPolicyFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachPolicyFromUserResponse) GetBody() *DetachPolicyFromUserResponseBody {
	return s.Body
}

func (s *DetachPolicyFromUserResponse) SetHeaders(v map[string]*string) *DetachPolicyFromUserResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyFromUserResponse) SetStatusCode(v int32) *DetachPolicyFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyFromUserResponse) SetBody(v *DetachPolicyFromUserResponseBody) *DetachPolicyFromUserResponse {
	s.Body = v
	return s
}

func (s *DetachPolicyFromUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
