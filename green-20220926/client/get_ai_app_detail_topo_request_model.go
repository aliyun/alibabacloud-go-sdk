// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailTopoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppDetailTopoRequest
	GetAppId() *string
	SetRegionId(v string) *GetAiAppDetailTopoRequest
	GetRegionId() *string
	SetTimeQuery(v *GetAiAppDetailTopoRequestTimeQuery) *GetAiAppDetailTopoRequest
	GetTimeQuery() *GetAiAppDetailTopoRequestTimeQuery
}

type GetAiAppDetailTopoRequest struct {
	// The application ID that identifies a specific AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time query.
	TimeQuery *GetAiAppDetailTopoRequestTimeQuery `json:"TimeQuery,omitempty" xml:"TimeQuery,omitempty" type:"Struct"`
}

func (s GetAiAppDetailTopoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailTopoRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailTopoRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppDetailTopoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppDetailTopoRequest) GetTimeQuery() *GetAiAppDetailTopoRequestTimeQuery {
	return s.TimeQuery
}

func (s *GetAiAppDetailTopoRequest) SetAppId(v string) *GetAiAppDetailTopoRequest {
	s.AppId = &v
	return s
}

func (s *GetAiAppDetailTopoRequest) SetRegionId(v string) *GetAiAppDetailTopoRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppDetailTopoRequest) SetTimeQuery(v *GetAiAppDetailTopoRequestTimeQuery) *GetAiAppDetailTopoRequest {
	s.TimeQuery = v
	return s
}

func (s *GetAiAppDetailTopoRequest) Validate() error {
	if s.TimeQuery != nil {
		if err := s.TimeQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiAppDetailTopoRequestTimeQuery struct {
	// The dimension.
	//
	// example:
	//
	// DAY
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2025-07-28 17:04:08
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2025-07-22 17:04:08
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAiAppDetailTopoRequestTimeQuery) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailTopoRequestTimeQuery) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailTopoRequestTimeQuery) GetDimension() *string {
	return s.Dimension
}

func (s *GetAiAppDetailTopoRequestTimeQuery) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAiAppDetailTopoRequestTimeQuery) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAiAppDetailTopoRequestTimeQuery) SetDimension(v string) *GetAiAppDetailTopoRequestTimeQuery {
	s.Dimension = &v
	return s
}

func (s *GetAiAppDetailTopoRequestTimeQuery) SetEndTime(v string) *GetAiAppDetailTopoRequestTimeQuery {
	s.EndTime = &v
	return s
}

func (s *GetAiAppDetailTopoRequestTimeQuery) SetStartTime(v string) *GetAiAppDetailTopoRequestTimeQuery {
	s.StartTime = &v
	return s
}

func (s *GetAiAppDetailTopoRequestTimeQuery) Validate() error {
	return dara.Validate(s)
}
