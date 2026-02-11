// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *QueryDatasetRequest
	GetDatasetId() *int64
	SetDateStr(v string) *QueryDatasetRequest
	GetDateStr() *string
	SetDimensions(v []*QueryDatasetRequestDimensions) *QueryDatasetRequest
	GetDimensions() []*QueryDatasetRequestDimensions
	SetHungryMode(v bool) *QueryDatasetRequest
	GetHungryMode() *bool
	SetIntervalInSec(v int32) *QueryDatasetRequest
	GetIntervalInSec() *int32
	SetIsDrillDown(v bool) *QueryDatasetRequest
	GetIsDrillDown() *bool
	SetLimit(v int32) *QueryDatasetRequest
	GetLimit() *int32
	SetMaxTime(v int64) *QueryDatasetRequest
	GetMaxTime() *int64
	SetMeasures(v []*string) *QueryDatasetRequest
	GetMeasures() []*string
	SetMinTime(v int64) *QueryDatasetRequest
	GetMinTime() *int64
	SetOptionalDims(v []*QueryDatasetRequestOptionalDims) *QueryDatasetRequest
	GetOptionalDims() []*QueryDatasetRequestOptionalDims
	SetOrderByKey(v string) *QueryDatasetRequest
	GetOrderByKey() *string
	SetProxyUserId(v string) *QueryDatasetRequest
	GetProxyUserId() *string
	SetReduceTail(v bool) *QueryDatasetRequest
	GetReduceTail() *bool
	SetRequiredDims(v []*QueryDatasetRequestRequiredDims) *QueryDatasetRequest
	GetRequiredDims() []*QueryDatasetRequestRequiredDims
}

type QueryDatasetRequest struct {
	// This parameter is required.
	DatasetId  *int64                           `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DateStr    *string                          `json:"DateStr,omitempty" xml:"DateStr,omitempty"`
	Dimensions []*QueryDatasetRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	HungryMode *bool                            `json:"HungryMode,omitempty" xml:"HungryMode,omitempty"`
	// This parameter is required.
	IntervalInSec *int32 `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	IsDrillDown   *bool  `json:"IsDrillDown,omitempty" xml:"IsDrillDown,omitempty"`
	Limit         *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is required.
	MaxTime  *int64    `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	Measures []*string `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	// This parameter is required.
	MinTime      *int64                             `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	OptionalDims []*QueryDatasetRequestOptionalDims `json:"OptionalDims,omitempty" xml:"OptionalDims,omitempty" type:"Repeated"`
	OrderByKey   *string                            `json:"OrderByKey,omitempty" xml:"OrderByKey,omitempty"`
	ProxyUserId  *string                            `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ReduceTail   *bool                              `json:"ReduceTail,omitempty" xml:"ReduceTail,omitempty"`
	RequiredDims []*QueryDatasetRequestRequiredDims `json:"RequiredDims,omitempty" xml:"RequiredDims,omitempty" type:"Repeated"`
}

func (s QueryDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *QueryDatasetRequest) GetDateStr() *string {
	return s.DateStr
}

func (s *QueryDatasetRequest) GetDimensions() []*QueryDatasetRequestDimensions {
	return s.Dimensions
}

func (s *QueryDatasetRequest) GetHungryMode() *bool {
	return s.HungryMode
}

func (s *QueryDatasetRequest) GetIntervalInSec() *int32 {
	return s.IntervalInSec
}

func (s *QueryDatasetRequest) GetIsDrillDown() *bool {
	return s.IsDrillDown
}

func (s *QueryDatasetRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryDatasetRequest) GetMaxTime() *int64 {
	return s.MaxTime
}

func (s *QueryDatasetRequest) GetMeasures() []*string {
	return s.Measures
}

func (s *QueryDatasetRequest) GetMinTime() *int64 {
	return s.MinTime
}

func (s *QueryDatasetRequest) GetOptionalDims() []*QueryDatasetRequestOptionalDims {
	return s.OptionalDims
}

func (s *QueryDatasetRequest) GetOrderByKey() *string {
	return s.OrderByKey
}

func (s *QueryDatasetRequest) GetProxyUserId() *string {
	return s.ProxyUserId
}

func (s *QueryDatasetRequest) GetReduceTail() *bool {
	return s.ReduceTail
}

func (s *QueryDatasetRequest) GetRequiredDims() []*QueryDatasetRequestRequiredDims {
	return s.RequiredDims
}

func (s *QueryDatasetRequest) SetDatasetId(v int64) *QueryDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetRequest) SetDateStr(v string) *QueryDatasetRequest {
	s.DateStr = &v
	return s
}

func (s *QueryDatasetRequest) SetDimensions(v []*QueryDatasetRequestDimensions) *QueryDatasetRequest {
	s.Dimensions = v
	return s
}

func (s *QueryDatasetRequest) SetHungryMode(v bool) *QueryDatasetRequest {
	s.HungryMode = &v
	return s
}

func (s *QueryDatasetRequest) SetIntervalInSec(v int32) *QueryDatasetRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryDatasetRequest) SetIsDrillDown(v bool) *QueryDatasetRequest {
	s.IsDrillDown = &v
	return s
}

func (s *QueryDatasetRequest) SetLimit(v int32) *QueryDatasetRequest {
	s.Limit = &v
	return s
}

func (s *QueryDatasetRequest) SetMaxTime(v int64) *QueryDatasetRequest {
	s.MaxTime = &v
	return s
}

func (s *QueryDatasetRequest) SetMeasures(v []*string) *QueryDatasetRequest {
	s.Measures = v
	return s
}

func (s *QueryDatasetRequest) SetMinTime(v int64) *QueryDatasetRequest {
	s.MinTime = &v
	return s
}

func (s *QueryDatasetRequest) SetOptionalDims(v []*QueryDatasetRequestOptionalDims) *QueryDatasetRequest {
	s.OptionalDims = v
	return s
}

func (s *QueryDatasetRequest) SetOrderByKey(v string) *QueryDatasetRequest {
	s.OrderByKey = &v
	return s
}

func (s *QueryDatasetRequest) SetProxyUserId(v string) *QueryDatasetRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryDatasetRequest) SetReduceTail(v bool) *QueryDatasetRequest {
	s.ReduceTail = &v
	return s
}

func (s *QueryDatasetRequest) SetRequiredDims(v []*QueryDatasetRequestRequiredDims) *QueryDatasetRequest {
	s.RequiredDims = v
	return s
}

func (s *QueryDatasetRequest) Validate() error {
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

type QueryDatasetRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDatasetRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetRequestDimensions) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *QueryDatasetRequestDimensions) GetType() *string {
	return s.Type
}

func (s *QueryDatasetRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *QueryDatasetRequestDimensions) SetKey(v string) *QueryDatasetRequestDimensions {
	s.Key = &v
	return s
}

func (s *QueryDatasetRequestDimensions) SetType(v string) *QueryDatasetRequestDimensions {
	s.Type = &v
	return s
}

func (s *QueryDatasetRequestDimensions) SetValue(v string) *QueryDatasetRequestDimensions {
	s.Value = &v
	return s
}

func (s *QueryDatasetRequestDimensions) Validate() error {
	return dara.Validate(s)
}

type QueryDatasetRequestOptionalDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDatasetRequestOptionalDims) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetRequestOptionalDims) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequestOptionalDims) GetKey() *string {
	return s.Key
}

func (s *QueryDatasetRequestOptionalDims) GetType() *string {
	return s.Type
}

func (s *QueryDatasetRequestOptionalDims) GetValue() *string {
	return s.Value
}

func (s *QueryDatasetRequestOptionalDims) SetKey(v string) *QueryDatasetRequestOptionalDims {
	s.Key = &v
	return s
}

func (s *QueryDatasetRequestOptionalDims) SetType(v string) *QueryDatasetRequestOptionalDims {
	s.Type = &v
	return s
}

func (s *QueryDatasetRequestOptionalDims) SetValue(v string) *QueryDatasetRequestOptionalDims {
	s.Value = &v
	return s
}

func (s *QueryDatasetRequestOptionalDims) Validate() error {
	return dara.Validate(s)
}

type QueryDatasetRequestRequiredDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDatasetRequestRequiredDims) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetRequestRequiredDims) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequestRequiredDims) GetKey() *string {
	return s.Key
}

func (s *QueryDatasetRequestRequiredDims) GetType() *string {
	return s.Type
}

func (s *QueryDatasetRequestRequiredDims) GetValue() *string {
	return s.Value
}

func (s *QueryDatasetRequestRequiredDims) SetKey(v string) *QueryDatasetRequestRequiredDims {
	s.Key = &v
	return s
}

func (s *QueryDatasetRequestRequiredDims) SetType(v string) *QueryDatasetRequestRequiredDims {
	s.Type = &v
	return s
}

func (s *QueryDatasetRequestRequiredDims) SetValue(v string) *QueryDatasetRequestRequiredDims {
	s.Value = &v
	return s
}

func (s *QueryDatasetRequestRequiredDims) Validate() error {
	return dara.Validate(s)
}
