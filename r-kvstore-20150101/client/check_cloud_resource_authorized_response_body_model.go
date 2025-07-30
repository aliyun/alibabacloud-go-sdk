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
}

type CheckCloudResourceAuthorizedResponseBody struct {
	// Indicates whether the instance is authorized to use KMS. Valid values:
	//
	// 	- **0**: The instance is authorized to use KMS.
	//
	// 	- **1**: The instance is not authorized to use KMS.
	//
	// 	- **2**: KMS is not activated. For more information, see [Activate KMS](https://help.aliyun.com/document_detail/153781.html).
	//
	// example:
	//
	// 1
	AuthorizationState *int32 `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// The ID of the request.
	//
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

func (s *CheckCloudResourceAuthorizedResponseBody) GetAuthorizationState() *int32 {
	return s.AuthorizationState
}

func (s *CheckCloudResourceAuthorizedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetAuthorizationState(v int32) *CheckCloudResourceAuthorizedResponseBody {
	s.AuthorizationState = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) SetRequestId(v string) *CheckCloudResourceAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponseBody) Validate() error {
	return dara.Validate(s)
}
