// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *UpdateSwimLaneRequest
	GetGroupName() *string
	SetLabelSelectorKey(v string) *UpdateSwimLaneRequest
	GetLabelSelectorKey() *string
	SetLabelSelectorValue(v string) *UpdateSwimLaneRequest
	GetLabelSelectorValue() *string
	SetServiceMeshId(v string) *UpdateSwimLaneRequest
	GetServiceMeshId() *string
	SetServicesList(v string) *UpdateSwimLaneRequest
	GetServicesList() *string
	SetSwimLaneName(v string) *UpdateSwimLaneRequest
	GetSwimLaneName() *string
}

type UpdateSwimLaneRequest struct {
	// The name of the lane group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The label key of the associated service workload. Set the value to `ASM_TRAFFIC_TAG`.
	//
	// example:
	//
	// ASM_TRAFFIC_TAG
	LabelSelectorKey *string `json:"LabelSelectorKey,omitempty" xml:"LabelSelectorKey,omitempty"`
	// The label value of the associated service workload.``
	//
	// example:
	//
	// v1
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// A list of services associated with the lane.
	//
	// example:
	//
	// ["sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mocka","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockb","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockc"]
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// s1
	SwimLaneName *string `json:"SwimLaneName,omitempty" xml:"SwimLaneName,omitempty"`
}

func (s UpdateSwimLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimLaneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateSwimLaneRequest) GetLabelSelectorKey() *string {
	return s.LabelSelectorKey
}

func (s *UpdateSwimLaneRequest) GetLabelSelectorValue() *string {
	return s.LabelSelectorValue
}

func (s *UpdateSwimLaneRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateSwimLaneRequest) GetServicesList() *string {
	return s.ServicesList
}

func (s *UpdateSwimLaneRequest) GetSwimLaneName() *string {
	return s.SwimLaneName
}

func (s *UpdateSwimLaneRequest) SetGroupName(v string) *UpdateSwimLaneRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetLabelSelectorKey(v string) *UpdateSwimLaneRequest {
	s.LabelSelectorKey = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetLabelSelectorValue(v string) *UpdateSwimLaneRequest {
	s.LabelSelectorValue = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetServiceMeshId(v string) *UpdateSwimLaneRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetServicesList(v string) *UpdateSwimLaneRequest {
	s.ServicesList = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetSwimLaneName(v string) *UpdateSwimLaneRequest {
	s.SwimLaneName = &v
	return s
}

func (s *UpdateSwimLaneRequest) Validate() error {
	return dara.Validate(s)
}
