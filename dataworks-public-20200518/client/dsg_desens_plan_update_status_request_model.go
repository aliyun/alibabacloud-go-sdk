// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanUpdateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int32) *DsgDesensPlanUpdateStatusRequest
	GetIds() []*int32
	SetSceneCode(v string) *DsgDesensPlanUpdateStatusRequest
	GetSceneCode() *string
	SetStatus(v int32) *DsgDesensPlanUpdateStatusRequest
	GetStatus() *int32
}

type DsgDesensPlanUpdateStatusRequest struct {
	// A collection of IDs of the data masking rules of which the status you want to modify.
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

func (s DsgDesensPlanUpdateStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanUpdateStatusRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanUpdateStatusRequest) GetIds() []*int32 {
	return s.Ids
}

func (s *DsgDesensPlanUpdateStatusRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgDesensPlanUpdateStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DsgDesensPlanUpdateStatusRequest) SetIds(v []*int32) *DsgDesensPlanUpdateStatusRequest {
	s.Ids = v
	return s
}

func (s *DsgDesensPlanUpdateStatusRequest) SetSceneCode(v string) *DsgDesensPlanUpdateStatusRequest {
	s.SceneCode = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusRequest) SetStatus(v int32) *DsgDesensPlanUpdateStatusRequest {
	s.Status = &v
	return s
}

func (s *DsgDesensPlanUpdateStatusRequest) Validate() error {
	return dara.Validate(s)
}
