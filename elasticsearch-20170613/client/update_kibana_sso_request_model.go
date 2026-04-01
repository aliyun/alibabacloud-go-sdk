// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaSsoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *UpdateKibanaSsoRequest
	GetEnable() *string
	SetNetworkType(v string) *UpdateKibanaSsoRequest
	GetNetworkType() *string
}

type UpdateKibanaSsoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	Enable *string `json:"enable,omitempty" xml:"enable,omitempty"`
	// PUBLIC, PRIVATE
	//
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC, PRIVATE
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s UpdateKibanaSsoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaSsoRequest) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSsoRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateKibanaSsoRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateKibanaSsoRequest) SetEnable(v string) *UpdateKibanaSsoRequest {
	s.Enable = &v
	return s
}

func (s *UpdateKibanaSsoRequest) SetNetworkType(v string) *UpdateKibanaSsoRequest {
	s.NetworkType = &v
	return s
}

func (s *UpdateKibanaSsoRequest) Validate() error {
	return dara.Validate(s)
}
