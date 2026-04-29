// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchPolarClawConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *PatchPolarClawConfigShrinkRequest
	GetApplicationId() *string
	SetConfigPatchShrink(v string) *PatchPolarClawConfigShrinkRequest
	GetConfigPatchShrink() *string
	SetRestart(v bool) *PatchPolarClawConfigShrinkRequest
	GetRestart() *bool
}

type PatchPolarClawConfigShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// {
	//
	//     "tools": {
	//
	//         "web": {
	//
	//             "search": {
	//
	//                 "enabled": false
	//
	//             }
	//
	//         }
	//
	//     }
	//
	// }
	ConfigPatchShrink *string `json:"ConfigPatch,omitempty" xml:"ConfigPatch,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s PatchPolarClawConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchPolarClawConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *PatchPolarClawConfigShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *PatchPolarClawConfigShrinkRequest) GetConfigPatchShrink() *string {
	return s.ConfigPatchShrink
}

func (s *PatchPolarClawConfigShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *PatchPolarClawConfigShrinkRequest) SetApplicationId(v string) *PatchPolarClawConfigShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *PatchPolarClawConfigShrinkRequest) SetConfigPatchShrink(v string) *PatchPolarClawConfigShrinkRequest {
	s.ConfigPatchShrink = &v
	return s
}

func (s *PatchPolarClawConfigShrinkRequest) SetRestart(v bool) *PatchPolarClawConfigShrinkRequest {
	s.Restart = &v
	return s
}

func (s *PatchPolarClawConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
