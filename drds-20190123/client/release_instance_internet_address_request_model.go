// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceInternetAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *ReleaseInstanceInternetAddressRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *ReleaseInstanceInternetAddressRequest
	GetRegionId() *string
}

type ReleaseInstanceInternetAddressRequest struct {
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The region where the instance is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstanceInternetAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceInternetAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ReleaseInstanceInternetAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseInstanceInternetAddressRequest) SetDrdsInstanceId(v string) *ReleaseInstanceInternetAddressRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ReleaseInstanceInternetAddressRequest) SetRegionId(v string) *ReleaseInstanceInternetAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseInstanceInternetAddressRequest) Validate() error {
	return dara.Validate(s)
}
