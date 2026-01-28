// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iARMSQueryDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *ARMSQueryDataSetRequest
	GetDatasetId() *int64
	SetDateStr(v string) *ARMSQueryDataSetRequest
	GetDateStr() *string
	SetDimensions(v []*ARMSQueryDataSetRequestDimensions) *ARMSQueryDataSetRequest
	GetDimensions() []*ARMSQueryDataSetRequestDimensions
	SetHackerUserId(v string) *ARMSQueryDataSetRequest
	GetHackerUserId() *string
	SetHungryMode(v bool) *ARMSQueryDataSetRequest
	GetHungryMode() *bool
	SetIntervalInSec(v int32) *ARMSQueryDataSetRequest
	GetIntervalInSec() *int32
	SetIsDrillDown(v bool) *ARMSQueryDataSetRequest
	GetIsDrillDown() *bool
	SetLimit(v int32) *ARMSQueryDataSetRequest
	GetLimit() *int32
	SetMaxTime(v int64) *ARMSQueryDataSetRequest
	GetMaxTime() *int64
	SetMeasures(v []*string) *ARMSQueryDataSetRequest
	GetMeasures() []*string
	SetMinTime(v int64) *ARMSQueryDataSetRequest
	GetMinTime() *int64
	SetOptionalDims(v []*ARMSQueryDataSetRequestOptionalDims) *ARMSQueryDataSetRequest
	GetOptionalDims() []*ARMSQueryDataSetRequestOptionalDims
	SetOrderByKey(v string) *ARMSQueryDataSetRequest
	GetOrderByKey() *string
	SetReduceTail(v bool) *ARMSQueryDataSetRequest
	GetReduceTail() *bool
	SetRequiredDims(v []*ARMSQueryDataSetRequestRequiredDims) *ARMSQueryDataSetRequest
	GetRequiredDims() []*ARMSQueryDataSetRequestRequiredDims
	SetSecurityToken(v string) *ARMSQueryDataSetRequest
	GetSecurityToken() *string
}

type ARMSQueryDataSetRequest struct {
	// This parameter is required.
	DatasetId    *int64                               `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DateStr      *string                              `json:"DateStr,omitempty" xml:"DateStr,omitempty"`
	Dimensions   []*ARMSQueryDataSetRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	HackerUserId *string                              `json:"HackerUserId,omitempty" xml:"HackerUserId,omitempty"`
	HungryMode   *bool                                `json:"HungryMode,omitempty" xml:"HungryMode,omitempty"`
	// This parameter is required.
	IntervalInSec *int32 `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	IsDrillDown   *bool  `json:"IsDrillDown,omitempty" xml:"IsDrillDown,omitempty"`
	Limit         *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is required.
	MaxTime  *int64    `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	Measures []*string `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	// This parameter is required.
	MinTime       *int64                                 `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	OptionalDims  []*ARMSQueryDataSetRequestOptionalDims `json:"OptionalDims,omitempty" xml:"OptionalDims,omitempty" type:"Repeated"`
	OrderByKey    *string                                `json:"OrderByKey,omitempty" xml:"OrderByKey,omitempty"`
	ReduceTail    *bool                                  `json:"ReduceTail,omitempty" xml:"ReduceTail,omitempty"`
	RequiredDims  []*ARMSQueryDataSetRequestRequiredDims `json:"RequiredDims,omitempty" xml:"RequiredDims,omitempty" type:"Repeated"`
	SecurityToken *string                                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ARMSQueryDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s ARMSQueryDataSetRequest) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ARMSQueryDataSetRequest) GetDateStr() *string {
	return s.DateStr
}

func (s *ARMSQueryDataSetRequest) GetDimensions() []*ARMSQueryDataSetRequestDimensions {
	return s.Dimensions
}

func (s *ARMSQueryDataSetRequest) GetHackerUserId() *string {
	return s.HackerUserId
}

func (s *ARMSQueryDataSetRequest) GetHungryMode() *bool {
	return s.HungryMode
}

func (s *ARMSQueryDataSetRequest) GetIntervalInSec() *int32 {
	return s.IntervalInSec
}

func (s *ARMSQueryDataSetRequest) GetIsDrillDown() *bool {
	return s.IsDrillDown
}

func (s *ARMSQueryDataSetRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ARMSQueryDataSetRequest) GetMaxTime() *int64 {
	return s.MaxTime
}

func (s *ARMSQueryDataSetRequest) GetMeasures() []*string {
	return s.Measures
}

func (s *ARMSQueryDataSetRequest) GetMinTime() *int64 {
	return s.MinTime
}

func (s *ARMSQueryDataSetRequest) GetOptionalDims() []*ARMSQueryDataSetRequestOptionalDims {
	return s.OptionalDims
}

func (s *ARMSQueryDataSetRequest) GetOrderByKey() *string {
	return s.OrderByKey
}

func (s *ARMSQueryDataSetRequest) GetReduceTail() *bool {
	return s.ReduceTail
}

func (s *ARMSQueryDataSetRequest) GetRequiredDims() []*ARMSQueryDataSetRequestRequiredDims {
	return s.RequiredDims
}

func (s *ARMSQueryDataSetRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ARMSQueryDataSetRequest) SetDatasetId(v int64) *ARMSQueryDataSetRequest {
	s.DatasetId = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetDateStr(v string) *ARMSQueryDataSetRequest {
	s.DateStr = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetDimensions(v []*ARMSQueryDataSetRequestDimensions) *ARMSQueryDataSetRequest {
	s.Dimensions = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetHackerUserId(v string) *ARMSQueryDataSetRequest {
	s.HackerUserId = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetHungryMode(v bool) *ARMSQueryDataSetRequest {
	s.HungryMode = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetIntervalInSec(v int32) *ARMSQueryDataSetRequest {
	s.IntervalInSec = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetIsDrillDown(v bool) *ARMSQueryDataSetRequest {
	s.IsDrillDown = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetLimit(v int32) *ARMSQueryDataSetRequest {
	s.Limit = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMaxTime(v int64) *ARMSQueryDataSetRequest {
	s.MaxTime = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMeasures(v []*string) *ARMSQueryDataSetRequest {
	s.Measures = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMinTime(v int64) *ARMSQueryDataSetRequest {
	s.MinTime = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetOptionalDims(v []*ARMSQueryDataSetRequestOptionalDims) *ARMSQueryDataSetRequest {
	s.OptionalDims = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetOrderByKey(v string) *ARMSQueryDataSetRequest {
	s.OrderByKey = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetReduceTail(v bool) *ARMSQueryDataSetRequest {
	s.ReduceTail = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetRequiredDims(v []*ARMSQueryDataSetRequestRequiredDims) *ARMSQueryDataSetRequest {
	s.RequiredDims = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetSecurityToken(v string) *ARMSQueryDataSetRequest {
	s.SecurityToken = &v
	return s
}

func (s *ARMSQueryDataSetRequest) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OptionalDims != nil {
		for _, item := range s.OptionalDims {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RequiredDims != nil {
		for _, item := range s.RequiredDims {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ARMSQueryDataSetRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ARMSQueryDataSetRequestDimensions) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *ARMSQueryDataSetRequestDimensions) GetType() *string {
	return s.Type
}

func (s *ARMSQueryDataSetRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *ARMSQueryDataSetRequestDimensions) SetKey(v string) *ARMSQueryDataSetRequestDimensions {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestDimensions) SetType(v string) *ARMSQueryDataSetRequestDimensions {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestDimensions) SetValue(v string) *ARMSQueryDataSetRequestDimensions {
	s.Value = &v
	return s
}

func (s *ARMSQueryDataSetRequestDimensions) Validate() error {
	return dara.Validate(s)
}

type ARMSQueryDataSetRequestOptionalDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestOptionalDims) String() string {
	return dara.Prettify(s)
}

func (s ARMSQueryDataSetRequestOptionalDims) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestOptionalDims) GetKey() *string {
	return s.Key
}

func (s *ARMSQueryDataSetRequestOptionalDims) GetType() *string {
	return s.Type
}

func (s *ARMSQueryDataSetRequestOptionalDims) GetValue() *string {
	return s.Value
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetKey(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetType(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetValue(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Value = &v
	return s
}

func (s *ARMSQueryDataSetRequestOptionalDims) Validate() error {
	return dara.Validate(s)
}

type ARMSQueryDataSetRequestRequiredDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestRequiredDims) String() string {
	return dara.Prettify(s)
}

func (s ARMSQueryDataSetRequestRequiredDims) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestRequiredDims) GetKey() *string {
	return s.Key
}

func (s *ARMSQueryDataSetRequestRequiredDims) GetType() *string {
	return s.Type
}

func (s *ARMSQueryDataSetRequestRequiredDims) GetValue() *string {
	return s.Value
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetKey(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetType(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetValue(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Value = &v
	return s
}

func (s *ARMSQueryDataSetRequestRequiredDims) Validate() error {
	return dara.Validate(s)
}
