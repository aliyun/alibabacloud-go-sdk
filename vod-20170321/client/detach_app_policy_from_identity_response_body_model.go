// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAppPolicyFromIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedPolicyNames(v []*string) *DetachAppPolicyFromIdentityResponseBody
	GetFailedPolicyNames() []*string
	SetNonExistPolicyNames(v []*string) *DetachAppPolicyFromIdentityResponseBody
	GetNonExistPolicyNames() []*string
	SetRequestId(v string) *DetachAppPolicyFromIdentityResponseBody
	GetRequestId() *string
}

type DetachAppPolicyFromIdentityResponseBody struct {
	// The names of the policies that failed to be granted to the RAM user or RAM role.
	FailedPolicyNames []*string `json:"FailedPolicyNames,omitempty" xml:"FailedPolicyNames,omitempty" type:"Repeated"`
	// The name of the policy that was not found.
	NonExistPolicyNames []*string `json:"NonExistPolicyNames,omitempty" xml:"NonExistPolicyNames,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-****-D7393642C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachAppPolicyFromIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachAppPolicyFromIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAppPolicyFromIdentityResponseBody) GetFailedPolicyNames() []*string {
	return s.FailedPolicyNames
}

func (s *DetachAppPolicyFromIdentityResponseBody) GetNonExistPolicyNames() []*string {
	return s.NonExistPolicyNames
}

func (s *DetachAppPolicyFromIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachAppPolicyFromIdentityResponseBody) SetFailedPolicyNames(v []*string) *DetachAppPolicyFromIdentityResponseBody {
	s.FailedPolicyNames = v
	return s
}

func (s *DetachAppPolicyFromIdentityResponseBody) SetNonExistPolicyNames(v []*string) *DetachAppPolicyFromIdentityResponseBody {
	s.NonExistPolicyNames = v
	return s
}

func (s *DetachAppPolicyFromIdentityResponseBody) SetRequestId(v string) *DetachAppPolicyFromIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachAppPolicyFromIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
