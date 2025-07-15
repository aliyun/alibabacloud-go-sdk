// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVSwitchResponseBody
	GetRequestId() *string
	SetVSwitchId(v string) *CreateVSwitchResponseBody
	GetVSwitchId() *string
}

type CreateVSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-25navfgbue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateVSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVSwitchResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateVSwitchResponseBody) SetRequestId(v string) *CreateVSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVSwitchResponseBody) SetVSwitchId(v string) *CreateVSwitchResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateVSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
