// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakebaseS3AccountPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLakebaseS3AccountPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLakebaseS3AccountPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLakebaseS3AccountPolicyResponseBody) *ModifyLakebaseS3AccountPolicyResponse
	GetBody() *ModifyLakebaseS3AccountPolicyResponseBody
}

type ModifyLakebaseS3AccountPolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLakebaseS3AccountPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLakebaseS3AccountPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakebaseS3AccountPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLakebaseS3AccountPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLakebaseS3AccountPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLakebaseS3AccountPolicyResponse) GetBody() *ModifyLakebaseS3AccountPolicyResponseBody {
	return s.Body
}

func (s *ModifyLakebaseS3AccountPolicyResponse) SetHeaders(v map[string]*string) *ModifyLakebaseS3AccountPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyResponse) SetStatusCode(v int32) *ModifyLakebaseS3AccountPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyResponse) SetBody(v *ModifyLakebaseS3AccountPolicyResponseBody) *ModifyLakebaseS3AccountPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
