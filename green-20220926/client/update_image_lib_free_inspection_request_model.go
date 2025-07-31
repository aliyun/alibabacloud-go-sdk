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
}

type UpdateImageLibFreeInspectionRequest struct {
	Config map[string]*int32 `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *UpdateImageLibFreeInspectionRequest) SetConfig(v map[string]*int32) *UpdateImageLibFreeInspectionRequest {
	s.Config = v
	return s
}

func (s *UpdateImageLibFreeInspectionRequest) SetRegionId(v string) *UpdateImageLibFreeInspectionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateImageLibFreeInspectionRequest) Validate() error {
	return dara.Validate(s)
}
