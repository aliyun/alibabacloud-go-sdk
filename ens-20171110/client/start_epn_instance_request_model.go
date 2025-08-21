// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *StartEpnInstanceRequest
	GetEPNInstanceId() *string
}

type StartEpnInstanceRequest struct {
	// The ID of the EPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-****
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s StartEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *StartEpnInstanceRequest) SetEPNInstanceId(v string) *StartEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *StartEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
