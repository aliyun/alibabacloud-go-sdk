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
	// The authorization status. Valid values:
	//
	// - **0**: Not authorized.
	//
	// - **1**: Authorized.
	//
	// - **2**: The user has not enabled KMS.
	//
	// example:
	//
	// 1
	AuthorizationState *int32 `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A0181AC4-F186-46ED-87CA-100C70B86729
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the role that is authorized to access Alibaba Cloud resources.
	//
	// > The response includes this parameter only when the return value of **AuthorizationState*	- is **1**.
	//
	// example:
	//
	// acs:ram::140****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
