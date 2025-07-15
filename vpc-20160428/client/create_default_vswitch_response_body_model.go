// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultVSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDefaultVSwitchResponseBody
	GetRequestId() *string
	SetVSwitchId(v string) *CreateDefaultVSwitchResponseBody
	GetVSwitchId() *string
}

type CreateDefaultVSwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the default vSwitch.
	//
	// example:
	//
	// vsw-bp1a4b5qhmxftjimq****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateDefaultVSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultVSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefaultVSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDefaultVSwitchResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDefaultVSwitchResponseBody) SetRequestId(v string) *CreateDefaultVSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefaultVSwitchResponseBody) SetVSwitchId(v string) *CreateDefaultVSwitchResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateDefaultVSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
