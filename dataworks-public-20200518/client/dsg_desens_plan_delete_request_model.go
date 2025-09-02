// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int32) *DsgDesensPlanDeleteRequest
	GetIds() []*int32
	SetSceneCode(v string) *DsgDesensPlanDeleteRequest
	GetSceneCode() *string
}

type DsgDesensPlanDeleteRequest struct {
	// A collection of data masking rules.
	//
	// This parameter is required.
	Ids []*int32 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The code of the level-1 data masking scenario to which the rule belongs. Valid values:
	//
	// 	- dataworks_display_desense_code: masking of displayed data in DataStudio and Data Map
	//
	// 	- maxcompute_desense_code: data masking at the MaxCompute compute engine layer
	//
	// 	- maxcompute_new_desense_code: data masking at the MaxCompute compute engine layer (new)
	//
	// 	- hologres_display_desense_code: data masking at the Hologres compute engine layer
	//
	// 	- dataworks_data_integration_desense_code: static data masking in Data Integration
	//
	// 	- dataworks_analysis_desense_code: masking of displayed data in DataAnalysis
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks_display_desense_code
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DsgDesensPlanDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanDeleteRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanDeleteRequest) GetIds() []*int32 {
	return s.Ids
}

func (s *DsgDesensPlanDeleteRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgDesensPlanDeleteRequest) SetIds(v []*int32) *DsgDesensPlanDeleteRequest {
	s.Ids = v
	return s
}

func (s *DsgDesensPlanDeleteRequest) SetSceneCode(v string) *DsgDesensPlanDeleteRequest {
	s.SceneCode = &v
	return s
}

func (s *DsgDesensPlanDeleteRequest) Validate() error {
	return dara.Validate(s)
}
