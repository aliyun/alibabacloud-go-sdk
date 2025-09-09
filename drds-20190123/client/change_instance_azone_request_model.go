// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeInstanceAzoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeVSwitch(v bool) *ChangeInstanceAzoneRequest
	GetChangeVSwitch() *bool
	SetDrdsInstanceId(v string) *ChangeInstanceAzoneRequest
	GetDrdsInstanceId() *string
	SetDrdsRegionId(v string) *ChangeInstanceAzoneRequest
	GetDrdsRegionId() *string
	SetNewVSwitch(v string) *ChangeInstanceAzoneRequest
	GetNewVSwitch() *string
	SetOriginAzoneId(v string) *ChangeInstanceAzoneRequest
	GetOriginAzoneId() *string
	SetTargetAzoneId(v string) *ChangeInstanceAzoneRequest
	GetTargetAzoneId() *string
}

type ChangeInstanceAzoneRequest struct {
	ChangeVSwitch *bool `json:"ChangeVSwitch,omitempty" xml:"ChangeVSwitch,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsjiii1b49****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DrdsRegionId *string `json:"DrdsRegionId,omitempty" xml:"DrdsRegionId,omitempty"`
	NewVSwitch   *string `json:"NewVSwitch,omitempty" xml:"NewVSwitch,omitempty"`
	// The source zone of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-B
	OriginAzoneId *string `json:"OriginAzoneId,omitempty" xml:"OriginAzoneId,omitempty"`
	// The destination zone to which you want to modify
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-A
	TargetAzoneId *string `json:"TargetAzoneId,omitempty" xml:"TargetAzoneId,omitempty"`
}

func (s ChangeInstanceAzoneRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeInstanceAzoneRequest) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneRequest) GetChangeVSwitch() *bool {
	return s.ChangeVSwitch
}

func (s *ChangeInstanceAzoneRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ChangeInstanceAzoneRequest) GetDrdsRegionId() *string {
	return s.DrdsRegionId
}

func (s *ChangeInstanceAzoneRequest) GetNewVSwitch() *string {
	return s.NewVSwitch
}

func (s *ChangeInstanceAzoneRequest) GetOriginAzoneId() *string {
	return s.OriginAzoneId
}

func (s *ChangeInstanceAzoneRequest) GetTargetAzoneId() *string {
	return s.TargetAzoneId
}

func (s *ChangeInstanceAzoneRequest) SetChangeVSwitch(v bool) *ChangeInstanceAzoneRequest {
	s.ChangeVSwitch = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetDrdsInstanceId(v string) *ChangeInstanceAzoneRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetDrdsRegionId(v string) *ChangeInstanceAzoneRequest {
	s.DrdsRegionId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetNewVSwitch(v string) *ChangeInstanceAzoneRequest {
	s.NewVSwitch = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetOriginAzoneId(v string) *ChangeInstanceAzoneRequest {
	s.OriginAzoneId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) SetTargetAzoneId(v string) *ChangeInstanceAzoneRequest {
	s.TargetAzoneId = &v
	return s
}

func (s *ChangeInstanceAzoneRequest) Validate() error {
	return dara.Validate(s)
}
