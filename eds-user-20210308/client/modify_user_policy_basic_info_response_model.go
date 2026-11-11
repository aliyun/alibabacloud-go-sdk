// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPolicyBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserPolicyBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserPolicyBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserPolicyBasicInfoResponseBody) *ModifyUserPolicyBasicInfoResponse
	GetBody() *ModifyUserPolicyBasicInfoResponseBody
}

type ModifyUserPolicyBasicInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserPolicyBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserPolicyBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPolicyBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserPolicyBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserPolicyBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserPolicyBasicInfoResponse) GetBody() *ModifyUserPolicyBasicInfoResponseBody {
	return s.Body
}

func (s *ModifyUserPolicyBasicInfoResponse) SetHeaders(v map[string]*string) *ModifyUserPolicyBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserPolicyBasicInfoResponse) SetStatusCode(v int32) *ModifyUserPolicyBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserPolicyBasicInfoResponse) SetBody(v *ModifyUserPolicyBasicInfoResponseBody) *ModifyUserPolicyBasicInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyUserPolicyBasicInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
