// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCloudResourceAuthorizedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CheckCloudResourceAuthorizedResponseBodyData) *CheckCloudResourceAuthorizedResponseBody
	GetData() *CheckCloudResourceAuthorizedResponseBodyData
	SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody
	GetRequestId() *string
}

type CheckCloudResourceAuthorizedResponseBody struct {
	Data *CheckCloudResourceAuthorizedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCloudResourceAuthorizedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponseBody) GetData() *CheckCloudResourceAuthorizedResponseBodyData {
	return s.Data
}

func (s *CheckCloudResourceAuthorizedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetData(v *CheckCloudResourceAuthorizedResponseBodyData) *CheckCloudResourceAuthorizedResponseBody {
	s.Data = v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckCloudResourceAuthorizedResponseBodyData struct {
	// example:
	//
	// 0
	AuthorizationState *string `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// example:
	//
	// acs:ram::123456789012****:role/AliyunRdsInstanceEncryptionDefaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s CheckCloudResourceAuthorizedResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponseBodyData) GetAuthorizationState() *string {
	return s.AuthorizationState
}

func (s *CheckCloudResourceAuthorizedResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CheckCloudResourceAuthorizedResponseBodyData) SetAuthorizationState(v string) *CheckCloudResourceAuthorizedResponseBodyData {
	s.AuthorizationState = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBodyData) SetRoleArn(v string) *CheckCloudResourceAuthorizedResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBodyData) Validate() error {
	return dara.Validate(s)
}
