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
	// 	- **1**: authorized
	//
	// 	- **0**: not authorized
	//
	// example:
	//
	// 1
	AuthorizationState *int32 `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role. A RAM role is a virtual identity that you can create within your Alibaba Cloud account. For more information, see [RAM role overview](https://help.aliyun.com/document_detail/93689.html).
	//
	// example:
	//
	// acs:ram::1406926****:role/aliyunrdsinstanceencryptiondefaultrole
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
