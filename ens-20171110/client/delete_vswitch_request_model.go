// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVSwitchId(v string) *DeleteVSwitchRequest
	GetVSwitchId() *string
}

type DeleteVSwitchRequest struct {
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-5m9xhl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DeleteVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchRequest) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeleteVSwitchRequest) SetVSwitchId(v string) *DeleteVSwitchRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeleteVSwitchRequest) Validate() error {
	return dara.Validate(s)
}
