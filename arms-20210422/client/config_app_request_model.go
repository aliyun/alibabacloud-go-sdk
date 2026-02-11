// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *ConfigAppRequest
	GetAppIds() *string
	SetEnable(v string) *ConfigAppRequest
	GetEnable() *string
	SetRegionId(v string) *ConfigAppRequest
	GetRegionId() *string
}

type ConfigAppRequest struct {
	// This parameter is required.
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigAppRequest) GoString() string {
	return s.String()
}

func (s *ConfigAppRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *ConfigAppRequest) GetEnable() *string {
	return s.Enable
}

func (s *ConfigAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigAppRequest) SetAppIds(v string) *ConfigAppRequest {
	s.AppIds = &v
	return s
}

func (s *ConfigAppRequest) SetEnable(v string) *ConfigAppRequest {
	s.Enable = &v
	return s
}

func (s *ConfigAppRequest) SetRegionId(v string) *ConfigAppRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigAppRequest) Validate() error {
	return dara.Validate(s)
}
