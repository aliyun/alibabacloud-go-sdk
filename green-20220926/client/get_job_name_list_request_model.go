// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *GetJobNameListRequest
	GetEndDate() *string
	SetQuery(v string) *GetJobNameListRequest
	GetQuery() *string
	SetRegionId(v string) *GetJobNameListRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *GetJobNameListRequest
	GetSort() map[string]*string
	SetStartDate(v string) *GetJobNameListRequest
	GetStartDate() *string
}

type GetJobNameListRequest struct {
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetJobNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobNameListRequest) GoString() string {
	return s.String()
}

func (s *GetJobNameListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetJobNameListRequest) GetQuery() *string {
	return s.Query
}

func (s *GetJobNameListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetJobNameListRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *GetJobNameListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetJobNameListRequest) SetEndDate(v string) *GetJobNameListRequest {
	s.EndDate = &v
	return s
}

func (s *GetJobNameListRequest) SetQuery(v string) *GetJobNameListRequest {
	s.Query = &v
	return s
}

func (s *GetJobNameListRequest) SetRegionId(v string) *GetJobNameListRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobNameListRequest) SetSort(v map[string]*string) *GetJobNameListRequest {
	s.Sort = v
	return s
}

func (s *GetJobNameListRequest) SetStartDate(v string) *GetJobNameListRequest {
	s.StartDate = &v
	return s
}

func (s *GetJobNameListRequest) Validate() error {
	return dara.Validate(s)
}
