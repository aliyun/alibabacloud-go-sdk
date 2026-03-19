// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBasicStatisticsResponseBody
	GetCode() *string
	SetGlobalStatistics(v *GetBasicStatisticsResponseBodyGlobalStatistics) *GetBasicStatisticsResponseBody
	GetGlobalStatistics() *GetBasicStatisticsResponseBodyGlobalStatistics
	SetMessage(v string) *GetBasicStatisticsResponseBody
	GetMessage() *string
	SetRegionStatistics(v []*GetBasicStatisticsResponseBodyRegionStatistics) *GetBasicStatisticsResponseBody
	GetRegionStatistics() []*GetBasicStatisticsResponseBodyRegionStatistics
	SetRequestId(v string) *GetBasicStatisticsResponseBody
	GetRequestId() *string
	SetSourceType(v string) *GetBasicStatisticsResponseBody
	GetSourceType() *string
	SetSuccess(v bool) *GetBasicStatisticsResponseBody
	GetSuccess() *bool
}

type GetBasicStatisticsResponseBody struct {
	// example:
	//
	// 200
	Code             *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	GlobalStatistics *GetBasicStatisticsResponseBodyGlobalStatistics `json:"GlobalStatistics,omitempty" xml:"GlobalStatistics,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message          *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionStatistics []*GetBasicStatisticsResponseBodyRegionStatistics `json:"RegionStatistics,omitempty" xml:"RegionStatistics,omitempty" type:"Repeated"`
	// example:
	//
	// EB526A5D-1FE2-51C1-B790-1732C1DBA969
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBasicStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBasicStatisticsResponseBody) GetGlobalStatistics() *GetBasicStatisticsResponseBodyGlobalStatistics {
	return s.GlobalStatistics
}

func (s *GetBasicStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBasicStatisticsResponseBody) GetRegionStatistics() []*GetBasicStatisticsResponseBodyRegionStatistics {
	return s.RegionStatistics
}

func (s *GetBasicStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicStatisticsResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetBasicStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBasicStatisticsResponseBody) SetCode(v string) *GetBasicStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetBasicStatisticsResponseBody) SetGlobalStatistics(v *GetBasicStatisticsResponseBodyGlobalStatistics) *GetBasicStatisticsResponseBody {
	s.GlobalStatistics = v
	return s
}

func (s *GetBasicStatisticsResponseBody) SetMessage(v string) *GetBasicStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetBasicStatisticsResponseBody) SetRegionStatistics(v []*GetBasicStatisticsResponseBodyRegionStatistics) *GetBasicStatisticsResponseBody {
	s.RegionStatistics = v
	return s
}

func (s *GetBasicStatisticsResponseBody) SetRequestId(v string) *GetBasicStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicStatisticsResponseBody) SetSourceType(v string) *GetBasicStatisticsResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetBasicStatisticsResponseBody) SetSuccess(v bool) *GetBasicStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetBasicStatisticsResponseBody) Validate() error {
	if s.GlobalStatistics != nil {
		if err := s.GlobalStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.RegionStatistics != nil {
		for _, item := range s.RegionStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBasicStatisticsResponseBodyGlobalStatistics struct {
	// example:
	//
	// 42949672960
	ProtectedDataSize *int64 `json:"ProtectedDataSize,omitempty" xml:"ProtectedDataSize,omitempty"`
	// example:
	//
	// 5
	ProtectedResourceCount *int32 `json:"ProtectedResourceCount,omitempty" xml:"ProtectedResourceCount,omitempty"`
}

func (s GetBasicStatisticsResponseBodyGlobalStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetBasicStatisticsResponseBodyGlobalStatistics) GoString() string {
	return s.String()
}

func (s *GetBasicStatisticsResponseBodyGlobalStatistics) GetProtectedDataSize() *int64 {
	return s.ProtectedDataSize
}

func (s *GetBasicStatisticsResponseBodyGlobalStatistics) GetProtectedResourceCount() *int32 {
	return s.ProtectedResourceCount
}

func (s *GetBasicStatisticsResponseBodyGlobalStatistics) SetProtectedDataSize(v int64) *GetBasicStatisticsResponseBodyGlobalStatistics {
	s.ProtectedDataSize = &v
	return s
}

func (s *GetBasicStatisticsResponseBodyGlobalStatistics) SetProtectedResourceCount(v int32) *GetBasicStatisticsResponseBodyGlobalStatistics {
	s.ProtectedResourceCount = &v
	return s
}

func (s *GetBasicStatisticsResponseBodyGlobalStatistics) Validate() error {
	return dara.Validate(s)
}

type GetBasicStatisticsResponseBodyRegionStatistics struct {
	// example:
	//
	// 42949672960
	ProtectedDataSize *int64 `json:"ProtectedDataSize,omitempty" xml:"ProtectedDataSize,omitempty"`
	// example:
	//
	// 5
	ProtectedResourceCount *int32 `json:"ProtectedResourceCount,omitempty" xml:"ProtectedResourceCount,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicStatisticsResponseBodyRegionStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetBasicStatisticsResponseBodyRegionStatistics) GoString() string {
	return s.String()
}

func (s *GetBasicStatisticsResponseBodyRegionStatistics) GetProtectedDataSize() *int64 {
	return s.ProtectedDataSize
}

func (s *GetBasicStatisticsResponseBodyRegionStatistics) GetProtectedResourceCount() *int32 {
	return s.ProtectedResourceCount
}

func (s *GetBasicStatisticsResponseBodyRegionStatistics) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicStatisticsResponseBodyRegionStatistics) SetProtectedDataSize(v int64) *GetBasicStatisticsResponseBodyRegionStatistics {
	s.ProtectedDataSize = &v
	return s
}

func (s *GetBasicStatisticsResponseBodyRegionStatistics) SetProtectedResourceCount(v int32) *GetBasicStatisticsResponseBodyRegionStatistics {
	s.ProtectedResourceCount = &v
	return s
}

func (s *GetBasicStatisticsResponseBodyRegionStatistics) SetRegionId(v string) *GetBasicStatisticsResponseBodyRegionStatistics {
	s.RegionId = &v
	return s
}

func (s *GetBasicStatisticsResponseBodyRegionStatistics) Validate() error {
	return dara.Validate(s)
}
