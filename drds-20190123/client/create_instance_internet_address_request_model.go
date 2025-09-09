// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInternetAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *CreateInstanceInternetAddressRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *CreateInstanceInternetAddressRequest
	GetRegionId() *string
}

type CreateInstanceInternetAddressRequest struct {
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds****************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region to which the DRDS instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstanceInternetAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInternetAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *CreateInstanceInternetAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceInternetAddressRequest) SetDrdsInstanceId(v string) *CreateInstanceInternetAddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateInstanceInternetAddressRequest) SetRegionId(v string) *CreateInstanceInternetAddressRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceInternetAddressRequest) Validate() error {
	return dara.Validate(s)
}
