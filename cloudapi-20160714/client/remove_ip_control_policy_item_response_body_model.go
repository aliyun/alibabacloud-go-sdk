// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpControlPolicyItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveIpControlPolicyItemResponseBody
	GetRequestId() *string
}

type RemoveIpControlPolicyItemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIpControlPolicyItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpControlPolicyItemResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIpControlPolicyItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveIpControlPolicyItemResponseBody) SetRequestId(v string) *RemoveIpControlPolicyItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveIpControlPolicyItemResponseBody) Validate() error {
	return dara.Validate(s)
}
