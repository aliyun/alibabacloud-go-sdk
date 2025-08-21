// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVSwitchesFromEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveVSwitchesFromEpnInstanceResponseBody
	GetRequestId() *string
}

type RemoveVSwitchesFromEpnInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVSwitchesFromEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveVSwitchesFromEpnInstanceResponseBody) SetRequestId(v string) *RemoveVSwitchesFromEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
