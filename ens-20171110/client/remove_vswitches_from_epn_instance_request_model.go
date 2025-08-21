// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVSwitchesFromEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *RemoveVSwitchesFromEpnInstanceRequest
	GetEPNInstanceId() *string
	SetVSwitchesInfo(v string) *RemoveVSwitchesFromEpnInstanceRequest
	GetVSwitchesInfo() *string
}

type RemoveVSwitchesFromEpnInstanceRequest struct {
	// The ID of theEPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-****
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The internal networking information that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// VSwitchesInfo=[{"VSwitchId":"vs-ixxxx"},{"VSwitchId":"vs-ixxxx"}]
	VSwitchesInfo *string `json:"VSwitchesInfo,omitempty" xml:"VSwitchesInfo,omitempty"`
}

func (s RemoveVSwitchesFromEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) GetVSwitchesInfo() *string {
	return s.VSwitchesInfo
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) SetEPNInstanceId(v string) *RemoveVSwitchesFromEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) SetVSwitchesInfo(v string) *RemoveVSwitchesFromEpnInstanceRequest {
	s.VSwitchesInfo = &v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
