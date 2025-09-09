// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassicExpiredDays(v int32) *UpdateInstanceNetworkRequest
	GetClassicExpiredDays() *int32
	SetDrdsInstanceId(v string) *UpdateInstanceNetworkRequest
	GetDrdsInstanceId() *string
	SetRetainClassic(v bool) *UpdateInstanceNetworkRequest
	GetRetainClassic() *bool
	SetSrcInstanceNetworkType(v string) *UpdateInstanceNetworkRequest
	GetSrcInstanceNetworkType() *string
}

type UpdateInstanceNetworkRequest struct {
	// Specifies the retention period of the classic network endpoint. Unit: days.
	//
	// example:
	//
	// 30
	ClassicExpiredDays *int32 `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds******
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// Specifies whether to retain the classic network endpoint.
	//
	// example:
	//
	// true
	RetainClassic *bool `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	// The network type of the PolarDB-X 1.0 instance. Valid values:
	//
	// 	- vpc: Virtual Private Cloud (VPC)
	//
	// 	- classic: classic network
	//
	// This parameter is required.
	//
	// example:
	//
	// classic
	SrcInstanceNetworkType *string `json:"SrcInstanceNetworkType,omitempty" xml:"SrcInstanceNetworkType,omitempty"`
}

func (s UpdateInstanceNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkRequest) GetClassicExpiredDays() *int32 {
	return s.ClassicExpiredDays
}

func (s *UpdateInstanceNetworkRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *UpdateInstanceNetworkRequest) GetRetainClassic() *bool {
	return s.RetainClassic
}

func (s *UpdateInstanceNetworkRequest) GetSrcInstanceNetworkType() *string {
	return s.SrcInstanceNetworkType
}

func (s *UpdateInstanceNetworkRequest) SetClassicExpiredDays(v int32) *UpdateInstanceNetworkRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetDrdsInstanceId(v string) *UpdateInstanceNetworkRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetRetainClassic(v bool) *UpdateInstanceNetworkRequest {
	s.RetainClassic = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) SetSrcInstanceNetworkType(v string) *UpdateInstanceNetworkRequest {
	s.SrcInstanceNetworkType = &v
	return s
}

func (s *UpdateInstanceNetworkRequest) Validate() error {
	return dara.Validate(s)
}
