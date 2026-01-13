// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListScenesResponseBody
	GetRequestId() *string
	SetScenes(v []*ListScenesResponseBodyScenes) *ListScenesResponseBody
	GetScenes() []*ListScenesResponseBodyScenes
	SetTotalCount(v int64) *ListScenesResponseBody
	GetTotalCount() *int64
}

type ListScenesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B8987BF7-6028-5B17-80E0-251B7BD67BBA
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scenes    []*ListScenesResponseBodyScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScenesResponseBody) GetScenes() []*ListScenesResponseBodyScenes {
	return s.Scenes
}

func (s *ListScenesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListScenesResponseBody) SetRequestId(v string) *ListScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScenesResponseBody) SetScenes(v []*ListScenesResponseBodyScenes) *ListScenesResponseBody {
	s.Scenes = v
	return s
}

func (s *ListScenesResponseBody) SetTotalCount(v int64) *ListScenesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScenesResponseBody) Validate() error {
	if s.Scenes != nil {
		for _, item := range s.Scenes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScenesResponseBodyScenes struct {
	// example:
	//
	// This is a test.
	Description *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Flows       []*ListScenesResponseBodyScenesFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListScenesResponseBodyScenes) String() string {
	return dara.Prettify(s)
}

func (s ListScenesResponseBodyScenes) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBodyScenes) GetDescription() *string {
	return s.Description
}

func (s *ListScenesResponseBodyScenes) GetFlows() []*ListScenesResponseBodyScenesFlows {
	return s.Flows
}

func (s *ListScenesResponseBodyScenes) GetName() *string {
	return s.Name
}

func (s *ListScenesResponseBodyScenes) GetSceneId() *string {
	return s.SceneId
}

func (s *ListScenesResponseBodyScenes) SetDescription(v string) *ListScenesResponseBodyScenes {
	s.Description = &v
	return s
}

func (s *ListScenesResponseBodyScenes) SetFlows(v []*ListScenesResponseBodyScenesFlows) *ListScenesResponseBodyScenes {
	s.Flows = v
	return s
}

func (s *ListScenesResponseBodyScenes) SetName(v string) *ListScenesResponseBodyScenes {
	s.Name = &v
	return s
}

func (s *ListScenesResponseBodyScenes) SetSceneId(v string) *ListScenesResponseBodyScenes {
	s.SceneId = &v
	return s
}

func (s *ListScenesResponseBodyScenes) Validate() error {
	if s.Flows != nil {
		for _, item := range s.Flows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScenesResponseBodyScenesFlows struct {
	// example:
	//
	// liuliang1
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 流量1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ListScenesResponseBodyScenesFlows) String() string {
	return dara.Prettify(s)
}

func (s ListScenesResponseBodyScenesFlows) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBodyScenesFlows) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ListScenesResponseBodyScenesFlows) GetFlowName() *string {
	return s.FlowName
}

func (s *ListScenesResponseBodyScenesFlows) SetFlowCode(v string) *ListScenesResponseBodyScenesFlows {
	s.FlowCode = &v
	return s
}

func (s *ListScenesResponseBodyScenesFlows) SetFlowName(v string) *ListScenesResponseBodyScenesFlows {
	s.FlowName = &v
	return s
}

func (s *ListScenesResponseBodyScenesFlows) Validate() error {
	return dara.Validate(s)
}
