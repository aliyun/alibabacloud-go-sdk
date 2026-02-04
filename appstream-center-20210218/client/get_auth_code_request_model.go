// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateUser(v bool) *GetAuthCodeRequest
	GetAutoCreateUser() *bool
	SetEndUserId(v string) *GetAuthCodeRequest
	GetEndUserId() *string
	SetExternalUserId(v string) *GetAuthCodeRequest
	GetExternalUserId() *string
	SetPolicy(v string) *GetAuthCodeRequest
	GetPolicy() *string
}

type GetAuthCodeRequest struct {
	AutoCreateUser *bool `json:"AutoCreateUser,omitempty" xml:"AutoCreateUser,omitempty"`
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// alice
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// example:
	//
	// {
	//
	//       "Version": "1",
	//
	//       "Resource": {
	//
	//             "Type": "AppInstanceGroup",
	//
	//             "Id": "aig-9ciijz60n4xsv****"
	//
	//       }
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s GetAuthCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAuthCodeRequest) GetAutoCreateUser() *bool {
	return s.AutoCreateUser
}

func (s *GetAuthCodeRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetAuthCodeRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *GetAuthCodeRequest) GetPolicy() *string {
	return s.Policy
}

func (s *GetAuthCodeRequest) SetAutoCreateUser(v bool) *GetAuthCodeRequest {
	s.AutoCreateUser = &v
	return s
}

func (s *GetAuthCodeRequest) SetEndUserId(v string) *GetAuthCodeRequest {
	s.EndUserId = &v
	return s
}

func (s *GetAuthCodeRequest) SetExternalUserId(v string) *GetAuthCodeRequest {
	s.ExternalUserId = &v
	return s
}

func (s *GetAuthCodeRequest) SetPolicy(v string) *GetAuthCodeRequest {
	s.Policy = &v
	return s
}

func (s *GetAuthCodeRequest) Validate() error {
	return dara.Validate(s)
}
