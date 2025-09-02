// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanDeleteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *DsgDesensPlanDeleteShrinkRequest
	GetIdsShrink() *string
	SetSceneCode(v string) *DsgDesensPlanDeleteShrinkRequest
	GetSceneCode() *string
}

type DsgDesensPlanDeleteShrinkRequest struct {
	// A collection of data masking rules.
	//
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
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

func (s DsgDesensPlanDeleteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanDeleteShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanDeleteShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DsgDesensPlanDeleteShrinkRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgDesensPlanDeleteShrinkRequest) SetIdsShrink(v string) *DsgDesensPlanDeleteShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DsgDesensPlanDeleteShrinkRequest) SetSceneCode(v string) *DsgDesensPlanDeleteShrinkRequest {
	s.SceneCode = &v
	return s
}

func (s *DsgDesensPlanDeleteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
