// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageLibFreeInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]*int32) *UpdateImageLibFreeInspectionRequest
	GetConfig() map[string]*int32
	SetRegionId(v string) *UpdateImageLibFreeInspectionRequest
	GetRegionId() *string
	SetServiceCode(v string) *UpdateImageLibFreeInspectionRequest
	GetServiceCode() *string
}

type UpdateImageLibFreeInspectionRequest struct {
	// Configuration.
	Config map[string]*int32 `json:"Config,omitempty" xml:"Config,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s UpdateImageLibFreeInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageLibFreeInspectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionRequest) GetConfig() map[string]*int32 {
	return s.Config
}

func (s *UpdateImageLibFreeInspectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateImageLibFreeInspectionRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *UpdateImageLibFreeInspectionRequest) SetConfig(v map[string]*int32) *UpdateImageLibFreeInspectionRequest {
	s.Config = v
	return s
}

func (s *UpdateImageLibFreeInspectionRequest) SetRegionId(v string) *UpdateImageLibFreeInspectionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateImageLibFreeInspectionRequest) SetServiceCode(v string) *UpdateImageLibFreeInspectionRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateImageLibFreeInspectionRequest) Validate() error {
	return dara.Validate(s)
}
