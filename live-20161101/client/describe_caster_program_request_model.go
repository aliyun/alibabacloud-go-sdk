// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterProgramRequest
	GetCasterId() *string
	SetEndTime(v string) *DescribeCasterProgramRequest
	GetEndTime() *string
	SetEpisodeId(v string) *DescribeCasterProgramRequest
	GetEpisodeId() *string
	SetEpisodeType(v string) *DescribeCasterProgramRequest
	GetEpisodeType() *string
	SetOwnerId(v int64) *DescribeCasterProgramRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeCasterProgramRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeCasterProgramRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCasterProgramRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeCasterProgramRequest
	GetStartTime() *string
	SetStatus(v int32) *DescribeCasterProgramRequest
	GetStatus() *int32
}

type DescribeCasterProgramRequest struct {
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T10:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the episode.
	//
	// example:
	//
	// 1872639A-F203-4EC5-8E43-CB92E68F****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The type of the episode. Valid values:
	//
	// 	- **Resource**: a video resource
	//
	// 	- **Component**: a component
	//
	// example:
	//
	// Resource
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 5
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the episode. Valid values:
	//
	// 	- **0**: The episode is not played.
	//
	// 	- **1**: The episode is being played.
	//
	// 	- **2**: The playback of the episode is complete.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCasterProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterProgramRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCasterProgramRequest) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *DescribeCasterProgramRequest) GetEpisodeType() *string {
	return s.EpisodeType
}

func (s *DescribeCasterProgramRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterProgramRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeCasterProgramRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCasterProgramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterProgramRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCasterProgramRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCasterProgramRequest) SetCasterId(v string) *DescribeCasterProgramRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetEndTime(v string) *DescribeCasterProgramRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetEpisodeId(v string) *DescribeCasterProgramRequest {
	s.EpisodeId = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetEpisodeType(v string) *DescribeCasterProgramRequest {
	s.EpisodeType = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetOwnerId(v int64) *DescribeCasterProgramRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetPageNum(v int32) *DescribeCasterProgramRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetPageSize(v int32) *DescribeCasterProgramRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetRegionId(v string) *DescribeCasterProgramRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetStartTime(v string) *DescribeCasterProgramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetStatus(v int32) *DescribeCasterProgramRequest {
	s.Status = &v
	return s
}

func (s *DescribeCasterProgramRequest) Validate() error {
	return dara.Validate(s)
}
