// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetRunConfigurationRequest
	GetRegionId() *string
}

type GetRunConfigurationRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetRunConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRunConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetRunConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRunConfigurationRequest) SetRegionId(v string) *GetRunConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *GetRunConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
