// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntelligentStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntelligentStrategyList(v []*ListIntelligentStrategyResponseBodyIntelligentStrategyList) *ListIntelligentStrategyResponseBody
	GetIntelligentStrategyList() []*ListIntelligentStrategyResponseBodyIntelligentStrategyList
	SetRequestId(v string) *ListIntelligentStrategyResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListIntelligentStrategyResponseBody
	GetTotal() *int32
}

type ListIntelligentStrategyResponseBody struct {
	IntelligentStrategyList []*ListIntelligentStrategyResponseBodyIntelligentStrategyList `json:"IntelligentStrategyList,omitempty" xml:"IntelligentStrategyList,omitempty" type:"Repeated"`
	RequestId               *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total                   *int32                                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListIntelligentStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntelligentStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntelligentStrategyResponseBody) GetIntelligentStrategyList() []*ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	return s.IntelligentStrategyList
}

func (s *ListIntelligentStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntelligentStrategyResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListIntelligentStrategyResponseBody) SetIntelligentStrategyList(v []*ListIntelligentStrategyResponseBodyIntelligentStrategyList) *ListIntelligentStrategyResponseBody {
	s.IntelligentStrategyList = v
	return s
}

func (s *ListIntelligentStrategyResponseBody) SetRequestId(v string) *ListIntelligentStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntelligentStrategyResponseBody) SetTotal(v int32) *ListIntelligentStrategyResponseBody {
	s.Total = &v
	return s
}

func (s *ListIntelligentStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIntelligentStrategyResponseBodyIntelligentStrategyList struct {
	Conditions    *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExecuteParams *string `json:"ExecuteParams,omitempty" xml:"ExecuteParams,omitempty"`
	MatchLimit    *int64  `json:"MatchLimit,omitempty" xml:"MatchLimit,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority      *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	StrategyId    *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIntelligentStrategyResponseBodyIntelligentStrategyList) String() string {
	return dara.Prettify(s)
}

func (s ListIntelligentStrategyResponseBodyIntelligentStrategyList) GoString() string {
	return s.String()
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetConditions() *string {
	return s.Conditions
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetExecuteParams() *string {
	return s.ExecuteParams
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetMatchLimit() *int64 {
	return s.MatchLimit
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetName() *string {
	return s.Name
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetState() *string {
	return s.State
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) GetType() *string {
	return s.Type
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetConditions(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.Conditions = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetCreationTime(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.CreationTime = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetExecuteParams(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.ExecuteParams = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetMatchLimit(v int64) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.MatchLimit = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetModifiedTime(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.ModifiedTime = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetName(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.Name = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetPriority(v int32) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.Priority = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetStartTime(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.StartTime = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetState(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.State = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetStrategyId(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.StrategyId = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) SetType(v string) *ListIntelligentStrategyResponseBodyIntelligentStrategyList {
	s.Type = &v
	return s
}

func (s *ListIntelligentStrategyResponseBodyIntelligentStrategyList) Validate() error {
	return dara.Validate(s)
}
