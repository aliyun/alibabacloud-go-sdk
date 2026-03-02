// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpGreyServiceGroupCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *PdpGreyServiceGroupCreateCmd
	GetEnv() *string
	SetServiceId(v int64) *PdpGreyServiceGroupCreateCmd
	GetServiceId() *int64
}

type PdpGreyServiceGroupCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s PdpGreyServiceGroupCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpGreyServiceGroupCreateCmd) GoString() string {
	return s.String()
}

func (s *PdpGreyServiceGroupCreateCmd) GetEnv() *string {
	return s.Env
}

func (s *PdpGreyServiceGroupCreateCmd) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpGreyServiceGroupCreateCmd) SetEnv(v string) *PdpGreyServiceGroupCreateCmd {
	s.Env = &v
	return s
}

func (s *PdpGreyServiceGroupCreateCmd) SetServiceId(v int64) *PdpGreyServiceGroupCreateCmd {
	s.ServiceId = &v
	return s
}

func (s *PdpGreyServiceGroupCreateCmd) Validate() error {
	return dara.Validate(s)
}
