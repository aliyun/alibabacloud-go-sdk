// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCloudResourceAuthorizedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationState(v int32) *CheckCloudResourceAuthorizedResponseBody
	GetAuthorizationState() *int32
	SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *CheckCloudResourceAuthorizedResponseBody
	GetRoleArn() *string
}

type CheckCloudResourceAuthorizedResponseBody struct {
	AuthorizationState *int32  `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleArn            *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckCloudResourceAuthorizedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponseBody) GetAuthorizationState() *int32 {
	return s.AuthorizationState
}

func (s *CheckCloudResourceAuthorizedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCloudResourceAuthorizedResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetAuthorizationState(v int32) *CheckCloudResourceAuthorizedResponseBody {
	s.AuthorizationState = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRoleArn(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RoleArn = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) Validate() error {
	return dara.Validate(s)
}
