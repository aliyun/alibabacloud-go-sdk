// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinVSwitchesToEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *JoinVSwitchesToEpnInstanceRequest
	GetEPNInstanceId() *string
	SetVSwitchesInfo(v string) *JoinVSwitchesToEpnInstanceRequest
	GetVSwitchesInfo() *string
}

type JoinVSwitchesToEpnInstanceRequest struct {
	// The edge-to-edge communication instance ID.
	//
	// The NetworkingModel of the EPN instance must be Connection or SpeedUpAndConnection. The SpeedUp pattern does not allow you to associate vSwitches.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-xxxx
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The internal networking information to be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"VSwitchId":"vs-ixxxx"},{"VSwitchId":"vs-ixxxx"}]
	VSwitchesInfo *string `json:"VSwitchesInfo,omitempty" xml:"VSwitchesInfo,omitempty"`
}

func (s JoinVSwitchesToEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *JoinVSwitchesToEpnInstanceRequest) GetVSwitchesInfo() *string {
	return s.VSwitchesInfo
}

func (s *JoinVSwitchesToEpnInstanceRequest) SetEPNInstanceId(v string) *JoinVSwitchesToEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *JoinVSwitchesToEpnInstanceRequest) SetVSwitchesInfo(v string) *JoinVSwitchesToEpnInstanceRequest {
	s.VSwitchesInfo = &v
	return s
}

func (s *JoinVSwitchesToEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
