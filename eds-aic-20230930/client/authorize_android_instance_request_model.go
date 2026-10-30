// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeAndroidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *AuthorizeAndroidInstanceRequest
	GetAndroidInstanceIds() []*string
	SetAuthorizeUserId(v string) *AuthorizeAndroidInstanceRequest
	GetAuthorizeUserId() *string
	SetUnAuthorizeUserId(v string) *AuthorizeAndroidInstanceRequest
	GetUnAuthorizeUserId() *string
}

type AuthorizeAndroidInstanceRequest struct {
	// The list of instance IDs.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// The ID of the user to assign.
	//
	// example:
	//
	// test
	AuthorizeUserId *string `json:"AuthorizeUserId,omitempty" xml:"AuthorizeUserId,omitempty"`
	// The ID of the user to unassign.
	//
	// example:
	//
	// test
	UnAuthorizeUserId *string `json:"UnAuthorizeUserId,omitempty" xml:"UnAuthorizeUserId,omitempty"`
}

func (s AuthorizeAndroidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeAndroidInstanceRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *AuthorizeAndroidInstanceRequest) GetAuthorizeUserId() *string {
	return s.AuthorizeUserId
}

func (s *AuthorizeAndroidInstanceRequest) GetUnAuthorizeUserId() *string {
	return s.UnAuthorizeUserId
}

func (s *AuthorizeAndroidInstanceRequest) SetAndroidInstanceIds(v []*string) *AuthorizeAndroidInstanceRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *AuthorizeAndroidInstanceRequest) SetAuthorizeUserId(v string) *AuthorizeAndroidInstanceRequest {
	s.AuthorizeUserId = &v
	return s
}

func (s *AuthorizeAndroidInstanceRequest) SetUnAuthorizeUserId(v string) *AuthorizeAndroidInstanceRequest {
	s.UnAuthorizeUserId = &v
	return s
}

func (s *AuthorizeAndroidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
