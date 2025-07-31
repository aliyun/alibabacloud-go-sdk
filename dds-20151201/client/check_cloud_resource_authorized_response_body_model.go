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
	// Indicates whether KMS keys are authorized to ApsaraDB for MongoDB instances. Valid values:
	//
	// 	- **0**: KMS keys are not authorized.
	//
	// 	- **1**: KMS keys are authorized.
	//
	// 	- **2**: KMS is not enabled.
	//
	// example:
	//
	// 1
	AuthorizationState *int32 `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A0181AC4-F186-46ED-87CA-100C70B86729
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role information of the authorized Alibaba Resource Name (ARN).
	//
	// >  This parameter is returned only when the value of the **AuthorizationState*	- parameter is **1**.
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
