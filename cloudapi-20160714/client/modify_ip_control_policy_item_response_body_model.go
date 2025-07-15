// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpControlPolicyItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpControlPolicyItemResponseBody
	GetRequestId() *string
}

type ModifyIpControlPolicyItemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpControlPolicyItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpControlPolicyItemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpControlPolicyItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpControlPolicyItemResponseBody) SetRequestId(v string) *ModifyIpControlPolicyItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpControlPolicyItemResponseBody) Validate() error {
	return dara.Validate(s)
}
