// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchPolarClawConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *PatchPolarClawConfigRequest
	GetApplicationId() *string
	SetConfigPatch(v map[string]interface{}) *PatchPolarClawConfigRequest
	GetConfigPatch() map[string]interface{}
	SetRestart(v bool) *PatchPolarClawConfigRequest
	GetRestart() *bool
}

type PatchPolarClawConfigRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// A JSON merge-patch object.
	//
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
	ConfigPatch map[string]interface{} `json:"ConfigPatch,omitempty" xml:"ConfigPatch,omitempty"`
	// Specifies whether to restart the gateway after applying the patch. The default is `true`.
	//
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s PatchPolarClawConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchPolarClawConfigRequest) GoString() string {
	return s.String()
}

func (s *PatchPolarClawConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *PatchPolarClawConfigRequest) GetConfigPatch() map[string]interface{} {
	return s.ConfigPatch
}

func (s *PatchPolarClawConfigRequest) GetRestart() *bool {
	return s.Restart
}

func (s *PatchPolarClawConfigRequest) SetApplicationId(v string) *PatchPolarClawConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *PatchPolarClawConfigRequest) SetConfigPatch(v map[string]interface{}) *PatchPolarClawConfigRequest {
	s.ConfigPatch = v
	return s
}

func (s *PatchPolarClawConfigRequest) SetRestart(v bool) *PatchPolarClawConfigRequest {
	s.Restart = &v
	return s
}

func (s *PatchPolarClawConfigRequest) Validate() error {
	return dara.Validate(s)
}
