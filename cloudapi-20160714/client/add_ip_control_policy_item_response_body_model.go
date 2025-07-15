// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpControlPolicyItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyItemId(v string) *AddIpControlPolicyItemResponseBody
	GetPolicyItemId() *string
	SetRequestId(v string) *AddIpControlPolicyItemResponseBody
	GetRequestId() *string
}

type AddIpControlPolicyItemResponseBody struct {
	// The ID of the policy. The ID is unique.
	//
	// example:
	//
	// P151617000829241
	PolicyItemId *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpControlPolicyItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIpControlPolicyItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpControlPolicyItemResponseBody) GetPolicyItemId() *string {
	return s.PolicyItemId
}

func (s *AddIpControlPolicyItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIpControlPolicyItemResponseBody) SetPolicyItemId(v string) *AddIpControlPolicyItemResponseBody {
	s.PolicyItemId = &v
	return s
}

func (s *AddIpControlPolicyItemResponseBody) SetRequestId(v string) *AddIpControlPolicyItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIpControlPolicyItemResponseBody) Validate() error {
	return dara.Validate(s)
}
