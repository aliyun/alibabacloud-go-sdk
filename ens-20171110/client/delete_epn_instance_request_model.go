// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *DeleteEpnInstanceRequest
	GetEPNInstanceId() *string
}

type DeleteEpnInstanceRequest struct {
	// The ID of the EPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-****
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s DeleteEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *DeleteEpnInstanceRequest) SetEPNInstanceId(v string) *DeleteEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *DeleteEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
