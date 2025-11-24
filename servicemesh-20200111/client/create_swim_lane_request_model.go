// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSwimLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CreateSwimLaneRequest
	GetGroupName() *string
	SetLabelSelectorKey(v string) *CreateSwimLaneRequest
	GetLabelSelectorKey() *string
	SetLabelSelectorValue(v string) *CreateSwimLaneRequest
	GetLabelSelectorValue() *string
	SetServiceMeshId(v string) *CreateSwimLaneRequest
	GetServiceMeshId() *string
	SetServicesList(v string) *CreateSwimLaneRequest
	GetServicesList() *string
	SetSwimLaneName(v string) *CreateSwimLaneRequest
	GetSwimLaneName() *string
}

type CreateSwimLaneRequest struct {
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
	// v3
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// *****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The list of services associated with the lane. The value is a JSON array. The format of a single service is `$Cluster name/$Cluster ID/$Namespace/$Service name`.
	//
	// example:
	//
	// [\\"sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mocka\\",\\"sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockb\\",\\"sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockc\\"]
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// s3
	SwimLaneName *string `json:"SwimLaneName,omitempty" xml:"SwimLaneName,omitempty"`
}

func (s CreateSwimLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSwimLaneRequest) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateSwimLaneRequest) GetLabelSelectorKey() *string {
	return s.LabelSelectorKey
}

func (s *CreateSwimLaneRequest) GetLabelSelectorValue() *string {
	return s.LabelSelectorValue
}

func (s *CreateSwimLaneRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateSwimLaneRequest) GetServicesList() *string {
	return s.ServicesList
}

func (s *CreateSwimLaneRequest) GetSwimLaneName() *string {
	return s.SwimLaneName
}

func (s *CreateSwimLaneRequest) SetGroupName(v string) *CreateSwimLaneRequest {
	s.GroupName = &v
	return s
}

func (s *CreateSwimLaneRequest) SetLabelSelectorKey(v string) *CreateSwimLaneRequest {
	s.LabelSelectorKey = &v
	return s
}

func (s *CreateSwimLaneRequest) SetLabelSelectorValue(v string) *CreateSwimLaneRequest {
	s.LabelSelectorValue = &v
	return s
}

func (s *CreateSwimLaneRequest) SetServiceMeshId(v string) *CreateSwimLaneRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateSwimLaneRequest) SetServicesList(v string) *CreateSwimLaneRequest {
	s.ServicesList = &v
	return s
}

func (s *CreateSwimLaneRequest) SetSwimLaneName(v string) *CreateSwimLaneRequest {
	s.SwimLaneName = &v
	return s
}

func (s *CreateSwimLaneRequest) Validate() error {
	return dara.Validate(s)
}
