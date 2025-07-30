// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarningStrategyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWarningStrategyConfigResponseBody
	GetCode() *string
	SetData(v *GetWarningStrategyConfigResponseBodyData) *GetWarningStrategyConfigResponseBody
	GetData() *GetWarningStrategyConfigResponseBodyData
	SetMessage(v string) *GetWarningStrategyConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWarningStrategyConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWarningStrategyConfigResponseBody
	GetSuccess() *bool
}

type GetWarningStrategyConfigResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetWarningStrategyConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWarningStrategyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWarningStrategyConfigResponseBody) GetData() *GetWarningStrategyConfigResponseBodyData {
	return s.Data
}

func (s *GetWarningStrategyConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWarningStrategyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWarningStrategyConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWarningStrategyConfigResponseBody) SetCode(v string) *GetWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetData(v *GetWarningStrategyConfigResponseBodyData) *GetWarningStrategyConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetMessage(v string) *GetWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetRequestId(v string) *GetWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetSuccess(v bool) *GetWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWarningStrategyConfigResponseBodyData struct {
	Id                  *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	IntervalTime        *int64                                                       `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	Lambda              *string                                                      `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Level               *int64                                                       `json:"Level,omitempty" xml:"Level,omitempty"`
	MaxNumber           *int64                                                       `json:"MaxNumber,omitempty" xml:"MaxNumber,omitempty"`
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	WarningStrategyList *GetWarningStrategyConfigResponseBodyDataWarningStrategyList `json:"WarningStrategyList,omitempty" xml:"WarningStrategyList,omitempty" type:"Struct"`
}

func (s GetWarningStrategyConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetWarningStrategyConfigResponseBodyData) GetIntervalTime() *int64 {
	return s.IntervalTime
}

func (s *GetWarningStrategyConfigResponseBodyData) GetLambda() *string {
	return s.Lambda
}

func (s *GetWarningStrategyConfigResponseBodyData) GetLevel() *int64 {
	return s.Level
}

func (s *GetWarningStrategyConfigResponseBodyData) GetMaxNumber() *int64 {
	return s.MaxNumber
}

func (s *GetWarningStrategyConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetWarningStrategyConfigResponseBodyData) GetWarningStrategyList() *GetWarningStrategyConfigResponseBodyDataWarningStrategyList {
	return s.WarningStrategyList
}

func (s *GetWarningStrategyConfigResponseBodyData) SetId(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetIntervalTime(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.IntervalTime = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetLambda(v string) *GetWarningStrategyConfigResponseBodyData {
	s.Lambda = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetLevel(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetMaxNumber(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.MaxNumber = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetName(v string) *GetWarningStrategyConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetWarningStrategyList(v *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) *GetWarningStrategyConfigResponseBodyData {
	s.WarningStrategyList = v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyList struct {
	WarningStrategyList []*GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList `json:"warningStrategyList,omitempty" xml:"warningStrategyList,omitempty" type:"Repeated"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyList) String() string {
	return dara.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyList) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) GetWarningStrategyList() []*GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	return s.WarningStrategyList
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) SetWarningStrategyList(v []*GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) *GetWarningStrategyConfigResponseBodyDataWarningStrategyList {
	s.WarningStrategyList = v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) Validate() error {
	return dara.Validate(s)
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList struct {
	Code                *string                                                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Duration            *int64                                                                               `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationExpression  *int64                                                                               `json:"DurationExpression,omitempty" xml:"DurationExpression,omitempty"`
	HitNumber           *int64                                                                               `json:"HitNumber,omitempty" xml:"HitNumber,omitempty"`
	HitNumberExpression *int64                                                                               `json:"HitNumberExpression,omitempty" xml:"HitNumberExpression,omitempty"`
	HitRuleList         *string                                                                              `json:"HitRuleList,omitempty" xml:"HitRuleList,omitempty"`
	HitType             *int64                                                                               `json:"HitType,omitempty" xml:"HitType,omitempty"`
	Id                  *int64                                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Range               *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Status              *int64                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) String() string {
	return dara.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetCode() *string {
	return s.Code
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetDuration() *int64 {
	return s.Duration
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetDurationExpression() *int64 {
	return s.DurationExpression
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetHitNumber() *int64 {
	return s.HitNumber
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetHitNumberExpression() *int64 {
	return s.HitNumberExpression
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetHitRuleList() *string {
	return s.HitRuleList
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetHitType() *int64 {
	return s.HitType
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetId() *int64 {
	return s.Id
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetRange() *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange {
	return s.Range
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GetStatus() *int64 {
	return s.Status
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetCode(v string) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Code = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetDuration(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Duration = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetDurationExpression(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.DurationExpression = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitNumber(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitNumber = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitNumberExpression(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitNumberExpression = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitRuleList(v string) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitRuleList = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitType(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitType = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetId(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Id = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetRange(v *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Range = v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetStatus(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Status = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) Validate() error {
	return dara.Validate(s)
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange struct {
	RangeNum *int64 `json:"RangeNum,omitempty" xml:"RangeNum,omitempty"`
	Type     *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) String() string {
	return dara.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) GetRangeNum() *int64 {
	return s.RangeNum
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) GetType() *int64 {
	return s.Type
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) SetRangeNum(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange {
	s.RangeNum = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) SetType(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange {
	s.Type = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) Validate() error {
	return dara.Validate(s)
}
