// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *StopEpnInstanceRequest
	GetEPNInstanceId() *string
}

type StopEpnInstanceRequest struct {
	// The ID of the EPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-20201014152822q2S9tQ
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s StopEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *StopEpnInstanceRequest) SetEPNInstanceId(v string) *StopEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *StopEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
