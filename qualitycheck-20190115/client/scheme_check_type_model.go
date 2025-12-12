// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchemeCheckType interface {
	dara.Model
	String() string
	GoString() string
	SetCheckName(v string) *SchemeCheckType
	GetCheckName() *string
	SetCheckType(v int64) *SchemeCheckType
	GetCheckType() *int64
	SetEnable(v int32) *SchemeCheckType
	GetEnable() *int32
	SetSchemeId(v int64) *SchemeCheckType
	GetSchemeId() *int64
	SetSchemeScoreInfoList(v []*SchemeCheckTypeSchemeScoreInfoList) *SchemeCheckType
	GetSchemeScoreInfoList() []*SchemeCheckTypeSchemeScoreInfoList
	SetScore(v int32) *SchemeCheckType
	GetScore() *int32
	SetSourceScore(v int32) *SchemeCheckType
	GetSourceScore() *int32
	SetTaskFlowScoreInfoList(v []*SchemeCheckTypeTaskFlowScoreInfoList) *SchemeCheckType
	GetTaskFlowScoreInfoList() []*SchemeCheckTypeTaskFlowScoreInfoList
}

type SchemeCheckType struct {
	CheckName             *string                                 `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckType             *int64                                  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Enable                *int32                                  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	SchemeId              *int64                                  `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	SchemeScoreInfoList   []*SchemeCheckTypeSchemeScoreInfoList   `json:"SchemeScoreInfoList,omitempty" xml:"SchemeScoreInfoList,omitempty" type:"Repeated"`
	Score                 *int32                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	SourceScore           *int32                                  `json:"SourceScore,omitempty" xml:"SourceScore,omitempty"`
	TaskFlowScoreInfoList []*SchemeCheckTypeTaskFlowScoreInfoList `json:"TaskFlowScoreInfoList,omitempty" xml:"TaskFlowScoreInfoList,omitempty" type:"Repeated"`
}

func (s SchemeCheckType) String() string {
	return dara.Prettify(s)
}

func (s SchemeCheckType) GoString() string {
	return s.String()
}

func (s *SchemeCheckType) GetCheckName() *string {
	return s.CheckName
}

func (s *SchemeCheckType) GetCheckType() *int64 {
	return s.CheckType
}

func (s *SchemeCheckType) GetEnable() *int32 {
	return s.Enable
}

func (s *SchemeCheckType) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *SchemeCheckType) GetSchemeScoreInfoList() []*SchemeCheckTypeSchemeScoreInfoList {
	return s.SchemeScoreInfoList
}

func (s *SchemeCheckType) GetScore() *int32 {
	return s.Score
}

func (s *SchemeCheckType) GetSourceScore() *int32 {
	return s.SourceScore
}

func (s *SchemeCheckType) GetTaskFlowScoreInfoList() []*SchemeCheckTypeTaskFlowScoreInfoList {
	return s.TaskFlowScoreInfoList
}

func (s *SchemeCheckType) SetCheckName(v string) *SchemeCheckType {
	s.CheckName = &v
	return s
}

func (s *SchemeCheckType) SetCheckType(v int64) *SchemeCheckType {
	s.CheckType = &v
	return s
}

func (s *SchemeCheckType) SetEnable(v int32) *SchemeCheckType {
	s.Enable = &v
	return s
}

func (s *SchemeCheckType) SetSchemeId(v int64) *SchemeCheckType {
	s.SchemeId = &v
	return s
}

func (s *SchemeCheckType) SetSchemeScoreInfoList(v []*SchemeCheckTypeSchemeScoreInfoList) *SchemeCheckType {
	s.SchemeScoreInfoList = v
	return s
}

func (s *SchemeCheckType) SetScore(v int32) *SchemeCheckType {
	s.Score = &v
	return s
}

func (s *SchemeCheckType) SetSourceScore(v int32) *SchemeCheckType {
	s.SourceScore = &v
	return s
}

func (s *SchemeCheckType) SetTaskFlowScoreInfoList(v []*SchemeCheckTypeTaskFlowScoreInfoList) *SchemeCheckType {
	s.TaskFlowScoreInfoList = v
	return s
}

func (s *SchemeCheckType) Validate() error {
	if s.SchemeScoreInfoList != nil {
		for _, item := range s.SchemeScoreInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskFlowScoreInfoList != nil {
		for _, item := range s.TaskFlowScoreInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SchemeCheckTypeSchemeScoreInfoList struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid              *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	ScoreNum         *int32  `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType     *int32  `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType *int32  `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreType        *int32  `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	TaskFlowId       *int64  `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowName     *string `json:"TaskFlowName,omitempty" xml:"TaskFlowName,omitempty"`
}

func (s SchemeCheckTypeSchemeScoreInfoList) String() string {
	return dara.Prettify(s)
}

func (s SchemeCheckTypeSchemeScoreInfoList) GoString() string {
	return s.String()
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetName() *string {
	return s.Name
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetRid() *int64 {
	return s.Rid
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetScoreNumType() *int32 {
	return s.ScoreNumType
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetScoreRuleHitType() *int32 {
	return s.ScoreRuleHitType
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetScoreType() *int32 {
	return s.ScoreType
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetTaskFlowId() *int64 {
	return s.TaskFlowId
}

func (s *SchemeCheckTypeSchemeScoreInfoList) GetTaskFlowName() *string {
	return s.TaskFlowName
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetName(v string) *SchemeCheckTypeSchemeScoreInfoList {
	s.Name = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetRid(v int64) *SchemeCheckTypeSchemeScoreInfoList {
	s.Rid = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreNum(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreNum = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreNumType(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreNumType = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreRuleHitType(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreRuleHitType = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetScoreType(v int32) *SchemeCheckTypeSchemeScoreInfoList {
	s.ScoreType = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetTaskFlowId(v int64) *SchemeCheckTypeSchemeScoreInfoList {
	s.TaskFlowId = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) SetTaskFlowName(v string) *SchemeCheckTypeSchemeScoreInfoList {
	s.TaskFlowName = &v
	return s
}

func (s *SchemeCheckTypeSchemeScoreInfoList) Validate() error {
	return dara.Validate(s)
}

type SchemeCheckTypeTaskFlowScoreInfoList struct {
	SchemeScoreInfoList []*SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList `json:"SchemeScoreInfoList,omitempty" xml:"SchemeScoreInfoList,omitempty" type:"Repeated"`
	TaskFlowId          *int64                                                     `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowName        *string                                                    `json:"TaskFlowName,omitempty" xml:"TaskFlowName,omitempty"`
	TaskFlowType        *int32                                                     `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
}

func (s SchemeCheckTypeTaskFlowScoreInfoList) String() string {
	return dara.Prettify(s)
}

func (s SchemeCheckTypeTaskFlowScoreInfoList) GoString() string {
	return s.String()
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) GetSchemeScoreInfoList() []*SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	return s.SchemeScoreInfoList
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) GetTaskFlowId() *int64 {
	return s.TaskFlowId
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) GetTaskFlowName() *string {
	return s.TaskFlowName
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) GetTaskFlowType() *int32 {
	return s.TaskFlowType
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetSchemeScoreInfoList(v []*SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.SchemeScoreInfoList = v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetTaskFlowId(v int64) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.TaskFlowId = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetTaskFlowName(v string) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.TaskFlowName = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) SetTaskFlowType(v int32) *SchemeCheckTypeTaskFlowScoreInfoList {
	s.TaskFlowType = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoList) Validate() error {
	if s.SchemeScoreInfoList != nil {
		for _, item := range s.SchemeScoreInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid              *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	ScoreNum         *int32  `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType     *int32  `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreRuleHitType *int32  `json:"ScoreRuleHitType,omitempty" xml:"ScoreRuleHitType,omitempty"`
	ScoreType        *int32  `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	TaskFlowId       *int64  `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	TaskFlowName     *string `json:"TaskFlowName,omitempty" xml:"TaskFlowName,omitempty"`
}

func (s SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) String() string {
	return dara.Prettify(s)
}

func (s SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GoString() string {
	return s.String()
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetName() *string {
	return s.Name
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetRid() *int64 {
	return s.Rid
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetScoreNumType() *int32 {
	return s.ScoreNumType
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetScoreRuleHitType() *int32 {
	return s.ScoreRuleHitType
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetScoreType() *int32 {
	return s.ScoreType
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetTaskFlowId() *int64 {
	return s.TaskFlowId
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) GetTaskFlowName() *string {
	return s.TaskFlowName
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetName(v string) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.Name = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetRid(v int64) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.Rid = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreNum(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreNum = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreNumType(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreNumType = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreRuleHitType(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreRuleHitType = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetScoreType(v int32) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.ScoreType = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetTaskFlowId(v int64) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.TaskFlowId = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) SetTaskFlowName(v string) *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList {
	s.TaskFlowName = &v
	return s
}

func (s *SchemeCheckTypeTaskFlowScoreInfoListSchemeScoreInfoList) Validate() error {
	return dara.Validate(s)
}
