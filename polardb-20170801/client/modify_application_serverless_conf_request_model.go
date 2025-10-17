// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationServerlessConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationServerlessConfRequest
	GetApplicationId() *string
	SetServerlessConfList(v []*ModifyApplicationServerlessConfRequestServerlessConfList) *ModifyApplicationServerlessConfRequest
	GetServerlessConfList() []*ModifyApplicationServerlessConfRequestServerlessConfList
}

type ModifyApplicationServerlessConfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	ServerlessConfList []*ModifyApplicationServerlessConfRequestServerlessConfList `json:"ServerlessConfList,omitempty" xml:"ServerlessConfList,omitempty" type:"Repeated"`
}

func (s ModifyApplicationServerlessConfRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationServerlessConfRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationServerlessConfRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationServerlessConfRequest) GetServerlessConfList() []*ModifyApplicationServerlessConfRequestServerlessConfList {
	return s.ServerlessConfList
}

func (s *ModifyApplicationServerlessConfRequest) SetApplicationId(v string) *ModifyApplicationServerlessConfRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationServerlessConfRequest) SetServerlessConfList(v []*ModifyApplicationServerlessConfRequestServerlessConfList) *ModifyApplicationServerlessConfRequest {
	s.ServerlessConfList = v
	return s
}

func (s *ModifyApplicationServerlessConfRequest) Validate() error {
	if s.ServerlessConfList != nil {
		for _, item := range s.ServerlessConfList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyApplicationServerlessConfRequestServerlessConfList struct {
	// example:
	//
	// gateway
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// 16
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
}

func (s ModifyApplicationServerlessConfRequestServerlessConfList) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationServerlessConfRequestServerlessConfList) GoString() string {
	return s.String()
}

func (s *ModifyApplicationServerlessConfRequestServerlessConfList) GetComponentType() *string {
	return s.ComponentType
}

func (s *ModifyApplicationServerlessConfRequestServerlessConfList) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *ModifyApplicationServerlessConfRequestServerlessConfList) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *ModifyApplicationServerlessConfRequestServerlessConfList) SetComponentType(v string) *ModifyApplicationServerlessConfRequestServerlessConfList {
	s.ComponentType = &v
	return s
}

func (s *ModifyApplicationServerlessConfRequestServerlessConfList) SetScaleMax(v string) *ModifyApplicationServerlessConfRequestServerlessConfList {
	s.ScaleMax = &v
	return s
}

func (s *ModifyApplicationServerlessConfRequestServerlessConfList) SetScaleMin(v string) *ModifyApplicationServerlessConfRequestServerlessConfList {
	s.ScaleMin = &v
	return s
}

func (s *ModifyApplicationServerlessConfRequestServerlessConfList) Validate() error {
	return dara.Validate(s)
}
