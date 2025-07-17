// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumDataForPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroup(v string) *GetRumDataForPageRequest
	GetAppGroup() *string
	SetCurrentPage(v int32) *GetRumDataForPageRequest
	GetCurrentPage() *int32
	SetEndTime(v int32) *GetRumDataForPageRequest
	GetEndTime() *int32
	SetPageSize(v int32) *GetRumDataForPageRequest
	GetPageSize() *int32
	SetPid(v string) *GetRumDataForPageRequest
	GetPid() *string
	SetQuery(v string) *GetRumDataForPageRequest
	GetQuery() *string
	SetRegionId(v string) *GetRumDataForPageRequest
	GetRegionId() *string
	SetStartTime(v int32) *GetRumDataForPageRequest
	GetStartTime() *int32
}

type GetRumDataForPageRequest struct {
	// The group to which the application belongs.
	//
	// example:
	//
	// default
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The beginning of the time range to query. The time is accurate to seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713774233
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The application ID.
	//
	// example:
	//
	// iixxxjcnuk@582846f37******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// A query statement that complies with the query syntax of Simple Log Service Logstore. For more information, see the parameters corresponding to this operation on the console page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 	- and app.id: xxxx@586810fbxxxx19f
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The time is accurate to seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713687833
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetRumDataForPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumDataForPageRequest) GoString() string {
	return s.String()
}

func (s *GetRumDataForPageRequest) GetAppGroup() *string {
	return s.AppGroup
}

func (s *GetRumDataForPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetRumDataForPageRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *GetRumDataForPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRumDataForPageRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRumDataForPageRequest) GetQuery() *string {
	return s.Query
}

func (s *GetRumDataForPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumDataForPageRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *GetRumDataForPageRequest) SetAppGroup(v string) *GetRumDataForPageRequest {
	s.AppGroup = &v
	return s
}

func (s *GetRumDataForPageRequest) SetCurrentPage(v int32) *GetRumDataForPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetRumDataForPageRequest) SetEndTime(v int32) *GetRumDataForPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetRumDataForPageRequest) SetPageSize(v int32) *GetRumDataForPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetRumDataForPageRequest) SetPid(v string) *GetRumDataForPageRequest {
	s.Pid = &v
	return s
}

func (s *GetRumDataForPageRequest) SetQuery(v string) *GetRumDataForPageRequest {
	s.Query = &v
	return s
}

func (s *GetRumDataForPageRequest) SetRegionId(v string) *GetRumDataForPageRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumDataForPageRequest) SetStartTime(v int32) *GetRumDataForPageRequest {
	s.StartTime = &v
	return s
}

func (s *GetRumDataForPageRequest) Validate() error {
	return dara.Validate(s)
}
