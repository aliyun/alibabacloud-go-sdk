// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v string) *GetQueryOptimizeDataTrendRequest
	GetEnd() *string
	SetEngine(v string) *GetQueryOptimizeDataTrendRequest
	GetEngine() *string
	SetInstanceIds(v string) *GetQueryOptimizeDataTrendRequest
	GetInstanceIds() *string
	SetRegion(v string) *GetQueryOptimizeDataTrendRequest
	GetRegion() *string
	SetStart(v string) *GetQueryOptimizeDataTrendRequest
	GetStart() *string
	SetTagNames(v string) *GetQueryOptimizeDataTrendRequest
	GetTagNames() *string
}

type GetQueryOptimizeDataTrendRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time, but not later than 00:00:00 (UTC+8) on the current day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1643040000000
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PolarDBMySQL**
	//
	// 	- **PostgreSQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region in which the instance resides. Valid values:
	//
	// 	- **cn-china**: Chinese mainland.
	//
	// 	- **cn-hongkong**: China (Hong Kong).
	//
	// 	- **ap-southeast-1**: Singapore.
	//
	// This parameter takes effect only if **InstanceIds*	- is left empty. If you leave **InstanceIds*	- empty, the system obtains data from the region specified by **Region**. By default, Region is set to **cn-china**. If you specify **InstanceIds**, **Region*	- does not take effect and the system obtains data from the region in which the first specified instance resides.****
	//
	// >  If your instances reside in the regions inside the Chinese mainland, set this parameter to **cn-china**.
	//
	// example:
	//
	// cn-china
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  You can specify a start time up to two months earlier than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1642435200000
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	TagNames *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
}

func (s GetQueryOptimizeDataTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTrendRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendRequest) GetEnd() *string {
	return s.End
}

func (s *GetQueryOptimizeDataTrendRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeDataTrendRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetQueryOptimizeDataTrendRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQueryOptimizeDataTrendRequest) GetStart() *string {
	return s.Start
}

func (s *GetQueryOptimizeDataTrendRequest) GetTagNames() *string {
	return s.TagNames
}

func (s *GetQueryOptimizeDataTrendRequest) SetEnd(v string) *GetQueryOptimizeDataTrendRequest {
	s.End = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetEngine(v string) *GetQueryOptimizeDataTrendRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetInstanceIds(v string) *GetQueryOptimizeDataTrendRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetRegion(v string) *GetQueryOptimizeDataTrendRequest {
	s.Region = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetStart(v string) *GetQueryOptimizeDataTrendRequest {
	s.Start = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetTagNames(v string) *GetQueryOptimizeDataTrendRequest {
	s.TagNames = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) Validate() error {
	return dara.Validate(s)
}
