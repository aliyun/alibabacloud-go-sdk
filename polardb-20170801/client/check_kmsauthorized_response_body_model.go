// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckKMSAuthorizedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationState(v int32) *CheckKMSAuthorizedResponseBody
	GetAuthorizationState() *int32
	SetDBClusterId(v string) *CheckKMSAuthorizedResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *CheckKMSAuthorizedResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *CheckKMSAuthorizedResponseBody
	GetRoleArn() *string
}

type CheckKMSAuthorizedResponseBody struct {
	// Indicates whether the cluster is authorized to use KMS. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 0
	AuthorizationState *int32 `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role. A RAM role is a virtual identity that you can create within your Alibaba Cloud account. For more information, see [RAM role overview](https://help.aliyun.com/document_detail/93689.html).
	//
	// example:
	//
	// acs:ram::1406926*****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckKMSAuthorizedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckKMSAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckKMSAuthorizedResponseBody) GetAuthorizationState() *int32 {
	return s.AuthorizationState
}

func (s *CheckKMSAuthorizedResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckKMSAuthorizedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckKMSAuthorizedResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CheckKMSAuthorizedResponseBody) SetAuthorizationState(v int32) *CheckKMSAuthorizedResponseBody {
	s.AuthorizationState = &v
	return s
}

func (s *CheckKMSAuthorizedResponseBody) SetDBClusterId(v string) *CheckKMSAuthorizedResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CheckKMSAuthorizedResponseBody) SetRequestId(v string) *CheckKMSAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckKMSAuthorizedResponseBody) SetRoleArn(v string) *CheckKMSAuthorizedResponseBody {
	s.RoleArn = &v
	return s
}

func (s *CheckKMSAuthorizedResponseBody) Validate() error {
	return dara.Validate(s)
}
