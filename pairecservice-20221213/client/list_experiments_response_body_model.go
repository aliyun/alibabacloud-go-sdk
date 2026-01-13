// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperiments(v []*ListExperimentsResponseBodyExperiments) *ListExperimentsResponseBody
	GetExperiments() []*ListExperimentsResponseBodyExperiments
	SetRequestId(v string) *ListExperimentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListExperimentsResponseBody
	GetTotalCount() *int64
}

type ListExperimentsResponseBody struct {
	Experiments []*ListExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 68075085-1A7D-55C2-B51D-7AD9B02A6DD6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBody) GetExperiments() []*ListExperimentsResponseBodyExperiments {
	return s.Experiments
}

func (s *ListExperimentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExperimentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExperimentsResponseBody) SetExperiments(v []*ListExperimentsResponseBodyExperiments) *ListExperimentsResponseBody {
	s.Experiments = v
	return s
}

func (s *ListExperimentsResponseBody) SetRequestId(v string) *ListExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentsResponseBody) SetTotalCount(v int64) *ListExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExperimentsResponseBody) Validate() error {
	if s.Experiments != nil {
		for _, item := range s.Experiments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExperimentsResponseBodyExperiments struct {
	// example:
	//
	// L1#EG1#E1
	AliasExperimentId *string `json:"AliasExperimentId,omitempty" xml:"AliasExperimentId,omitempty"`
	// example:
	//
	// 1,2,3
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// uid1,uid2,uid3
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// example:
	//
	// 3
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// experiment_test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListExperimentsResponseBodyExperiments) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyExperiments) GetAliasExperimentId() *string {
	return s.AliasExperimentId
}

func (s *ListExperimentsResponseBodyExperiments) GetBuckets() *string {
	return s.Buckets
}

func (s *ListExperimentsResponseBodyExperiments) GetConfig() *string {
	return s.Config
}

func (s *ListExperimentsResponseBodyExperiments) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *ListExperimentsResponseBodyExperiments) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *ListExperimentsResponseBodyExperiments) GetDescription() *string {
	return s.Description
}

func (s *ListExperimentsResponseBodyExperiments) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *ListExperimentsResponseBodyExperiments) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListExperimentsResponseBodyExperiments) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *ListExperimentsResponseBodyExperiments) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListExperimentsResponseBodyExperiments) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListExperimentsResponseBodyExperiments) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *ListExperimentsResponseBodyExperiments) GetLayerId() *string {
	return s.LayerId
}

func (s *ListExperimentsResponseBodyExperiments) GetName() *string {
	return s.Name
}

func (s *ListExperimentsResponseBodyExperiments) GetSceneId() *string {
	return s.SceneId
}

func (s *ListExperimentsResponseBodyExperiments) GetStatus() *string {
	return s.Status
}

func (s *ListExperimentsResponseBodyExperiments) GetType() *string {
	return s.Type
}

func (s *ListExperimentsResponseBodyExperiments) SetAliasExperimentId(v string) *ListExperimentsResponseBodyExperiments {
	s.AliasExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetBuckets(v string) *ListExperimentsResponseBodyExperiments {
	s.Buckets = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetConfig(v string) *ListExperimentsResponseBodyExperiments {
	s.Config = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDebugCrowdId(v string) *ListExperimentsResponseBodyExperiments {
	s.DebugCrowdId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDebugUsers(v string) *ListExperimentsResponseBodyExperiments {
	s.DebugUsers = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDescription(v string) *ListExperimentsResponseBodyExperiments {
	s.Description = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetExperimentGroupId(v string) *ListExperimentsResponseBodyExperiments {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetFlowPercent(v int32) *ListExperimentsResponseBodyExperiments {
	s.FlowPercent = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtCreateTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtCreateTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtModifiedTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetLaboratoryId(v string) *ListExperimentsResponseBodyExperiments {
	s.LaboratoryId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetLayerId(v string) *ListExperimentsResponseBodyExperiments {
	s.LayerId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetName(v string) *ListExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetSceneId(v string) *ListExperimentsResponseBodyExperiments {
	s.SceneId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetStatus(v string) *ListExperimentsResponseBodyExperiments {
	s.Status = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetType(v string) *ListExperimentsResponseBodyExperiments {
	s.Type = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) Validate() error {
	return dara.Validate(s)
}
