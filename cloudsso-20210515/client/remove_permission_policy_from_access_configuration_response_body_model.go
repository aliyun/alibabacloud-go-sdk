// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePermissionPolicyFromAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemovePermissionPolicyFromAccessConfigurationResponseBody
	GetRequestId() *string
}

type RemovePermissionPolicyFromAccessConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B13E4EE-3853-5852-9165-597C32AD8FB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePermissionPolicyFromAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePermissionPolicyFromAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponseBody) SetRequestId(v string) *RemovePermissionPolicyFromAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
