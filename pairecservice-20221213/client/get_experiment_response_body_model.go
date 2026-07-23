// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliasExperimentId(v string) *GetExperimentResponseBody
	GetAliasExperimentId() *string
	SetBuckets(v string) *GetExperimentResponseBody
	GetBuckets() *string
	SetConfig(v string) *GetExperimentResponseBody
	GetConfig() *string
	SetDebugCrowdId(v string) *GetExperimentResponseBody
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *GetExperimentResponseBody
	GetDebugUsers() *string
	SetDescription(v string) *GetExperimentResponseBody
	GetDescription() *string
	SetExperimentGroupId(v string) *GetExperimentResponseBody
	GetExperimentGroupId() *string
	SetFlowPercent(v int32) *GetExperimentResponseBody
	GetFlowPercent() *int32
	SetGmtCreateTime(v string) *GetExperimentResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetExperimentResponseBody
	GetGmtModifiedTime() *string
	SetLaboratoryId(v string) *GetExperimentResponseBody
	GetLaboratoryId() *string
	SetLayerId(v string) *GetExperimentResponseBody
	GetLayerId() *string
	SetName(v string) *GetExperimentResponseBody
	GetName() *string
	SetRequestId(v string) *GetExperimentResponseBody
	GetRequestId() *string
	SetSceneId(v string) *GetExperimentResponseBody
	GetSceneId() *string
	SetStatus(v string) *GetExperimentResponseBody
	GetStatus() *string
	SetType(v string) *GetExperimentResponseBody
	GetType() *string
}

type GetExperimentResponseBody struct {
	// The alias of the experiment.
	//
	// example:
	//
	// L1#EG1#E1
	AliasExperimentId *string `json:"AliasExperimentId,omitempty" xml:"AliasExperimentId,omitempty"`
	// A comma-separated list of bucket numbers.
	//
	// example:
	//
	// 1,2,3
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// The experiment configuration, in JSON format.
	//
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The debug crowd ID.
	//
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// The UIDs of debug users, which can be the UIDs of an Alibaba Cloud main account or a RAM user. Separate multiple UIDs with a comma (,).
	//
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// The experiment description.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The experiment group ID.
	//
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// The traffic percentage.
	//
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	// The creation time, in ISO 8601 format.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The modification time, in ISO 8601 format.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The laboratory ID.
	//
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// The layer ID.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// The experiment name.
	//
	// example:
	//
	// experiment_test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The status of the experiment. Valid values:<br>● Offline<br>● Online<br><br>
	//
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the experiment. Valid values:<br>● Baseline: a baseline experiment.<br>● Normal: a normal experiment.<br><br>
	//
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBody) GetAliasExperimentId() *string {
	return s.AliasExperimentId
}

func (s *GetExperimentResponseBody) GetBuckets() *string {
	return s.Buckets
}

func (s *GetExperimentResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetExperimentResponseBody) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *GetExperimentResponseBody) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *GetExperimentResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetExperimentResponseBody) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *GetExperimentResponseBody) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *GetExperimentResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetExperimentResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetExperimentResponseBody) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *GetExperimentResponseBody) GetLayerId() *string {
	return s.LayerId
}

func (s *GetExperimentResponseBody) GetName() *string {
	return s.Name
}

func (s *GetExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetExperimentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetExperimentResponseBody) GetType() *string {
	return s.Type
}

func (s *GetExperimentResponseBody) SetAliasExperimentId(v string) *GetExperimentResponseBody {
	s.AliasExperimentId = &v
	return s
}

func (s *GetExperimentResponseBody) SetBuckets(v string) *GetExperimentResponseBody {
	s.Buckets = &v
	return s
}

func (s *GetExperimentResponseBody) SetConfig(v string) *GetExperimentResponseBody {
	s.Config = &v
	return s
}

func (s *GetExperimentResponseBody) SetDebugCrowdId(v string) *GetExperimentResponseBody {
	s.DebugCrowdId = &v
	return s
}

func (s *GetExperimentResponseBody) SetDebugUsers(v string) *GetExperimentResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetExperimentResponseBody) SetDescription(v string) *GetExperimentResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentResponseBody) SetExperimentGroupId(v string) *GetExperimentResponseBody {
	s.ExperimentGroupId = &v
	return s
}

func (s *GetExperimentResponseBody) SetFlowPercent(v int32) *GetExperimentResponseBody {
	s.FlowPercent = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtCreateTime(v string) *GetExperimentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtModifiedTime(v string) *GetExperimentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetLaboratoryId(v string) *GetExperimentResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *GetExperimentResponseBody) SetLayerId(v string) *GetExperimentResponseBody {
	s.LayerId = &v
	return s
}

func (s *GetExperimentResponseBody) SetName(v string) *GetExperimentResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentResponseBody) SetRequestId(v string) *GetExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentResponseBody) SetSceneId(v string) *GetExperimentResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetExperimentResponseBody) SetStatus(v string) *GetExperimentResponseBody {
	s.Status = &v
	return s
}

func (s *GetExperimentResponseBody) SetType(v string) *GetExperimentResponseBody {
	s.Type = &v
	return s
}

func (s *GetExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
