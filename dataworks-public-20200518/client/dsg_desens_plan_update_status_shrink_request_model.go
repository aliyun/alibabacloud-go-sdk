// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanUpdateStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *DsgDesensPlanUpdateStatusShrinkRequest
	GetIdsShrink() *string
	SetSceneCode(v string) *DsgDesensPlanUpdateStatusShrinkRequest
	GetSceneCode() *string
	SetStatus(v int32) *DsgDesensPlanUpdateStatusShrinkRequest
	GetStatus() *int32
}

type DsgDesensPlanUpdateStatusShrinkRequest struct {
	// A collection of IDs of the data masking rules of which the status you want to modify.
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
	// The status of the data masking rule. Valid values:
	//
	// 	- 0: expired
	//
	// 	- 1: effective
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DsgDesensPlanUpdateStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanUpdateStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanUpdateStatusShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DsgDesensPlanUpdateStatusShrinkRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgDesensPlanUpdateStatusShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DsgDesensPlanUpdateStatusShrinkRequest) SetIdsShrink(v string) *DsgDesensPlanUpdateStatusShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusShrinkRequest) SetSceneCode(v string) *DsgDesensPlanUpdateStatusShrinkRequest {
	s.SceneCode = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusShrinkRequest) SetStatus(v int32) *DsgDesensPlanUpdateStatusShrinkRequest {
	s.Status = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
