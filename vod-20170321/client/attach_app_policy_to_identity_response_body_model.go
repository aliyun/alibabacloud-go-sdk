// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAppPolicyToIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedPolicyNames(v []*string) *AttachAppPolicyToIdentityResponseBody
	GetFailedPolicyNames() []*string
	SetNonExistPolicyNames(v []*string) *AttachAppPolicyToIdentityResponseBody
	GetNonExistPolicyNames() []*string
	SetRequestId(v string) *AttachAppPolicyToIdentityResponseBody
	GetRequestId() *string
}

type AttachAppPolicyToIdentityResponseBody struct {
	// The names of the policies that failed to be granted to the RAM user or RAM role.
	FailedPolicyNames []*string `json:"FailedPolicyNames,omitempty" xml:"FailedPolicyNames,omitempty" type:"Repeated"`
	// The names of the policies that were not found.
	NonExistPolicyNames []*string `json:"NonExistPolicyNames,omitempty" xml:"NonExistPolicyNames,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-****-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAppPolicyToIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachAppPolicyToIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAppPolicyToIdentityResponseBody) GetFailedPolicyNames() []*string {
	return s.FailedPolicyNames
}

func (s *AttachAppPolicyToIdentityResponseBody) GetNonExistPolicyNames() []*string {
	return s.NonExistPolicyNames
}

func (s *AttachAppPolicyToIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachAppPolicyToIdentityResponseBody) SetFailedPolicyNames(v []*string) *AttachAppPolicyToIdentityResponseBody {
	s.FailedPolicyNames = v
	return s
}

func (s *AttachAppPolicyToIdentityResponseBody) SetNonExistPolicyNames(v []*string) *AttachAppPolicyToIdentityResponseBody {
	s.NonExistPolicyNames = v
	return s
}

func (s *AttachAppPolicyToIdentityResponseBody) SetRequestId(v string) *AttachAppPolicyToIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachAppPolicyToIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
