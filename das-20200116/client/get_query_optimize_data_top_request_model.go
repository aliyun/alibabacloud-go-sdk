// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataTopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *GetQueryOptimizeDataTopRequest
	GetEngine() *string
	SetInstanceIds(v string) *GetQueryOptimizeDataTopRequest
	GetInstanceIds() *string
	SetRegion(v string) *GetQueryOptimizeDataTopRequest
	GetRegion() *string
	SetTagNames(v string) *GetQueryOptimizeDataTopRequest
	GetTagNames() *string
	SetTime(v string) *GetQueryOptimizeDataTopRequest
	GetTime() *string
	SetType(v string) *GetQueryOptimizeDataTopRequest
	GetType() *string
}

type GetQueryOptimizeDataTopRequest struct {
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
	// 	- **cn-china**: Chinese mainland
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// This parameter takes effect only if **InstanceIds*	- is left empty. If you leave **InstanceIds*	- empty, the system obtains data from the region set by **Region**. By default, Region is set to **cn-china**. If you specify **InstanceIds**, **Region*	- does not take effect and the system obtains data from the region in which the first specified instance resides.****
	//
	// >  Set this parameter to **cn-china*	- for all your instances that reside in the regions in the Chinese mainland.
	//
	// example:
	//
	// cn-china
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	TagNames *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
	// The time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1642953600000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of instances that you want to query. Valid values:
	//
	// 	- **RED**: the best-performing instances
	//
	// 	- **BLACK**: the worst-performing instances
	//
	// This parameter is required.
	//
	// example:
	//
	// RED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQueryOptimizeDataTopRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTopRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeDataTopRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetQueryOptimizeDataTopRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQueryOptimizeDataTopRequest) GetTagNames() *string {
	return s.TagNames
}

func (s *GetQueryOptimizeDataTopRequest) GetTime() *string {
	return s.Time
}

func (s *GetQueryOptimizeDataTopRequest) GetType() *string {
	return s.Type
}

func (s *GetQueryOptimizeDataTopRequest) SetEngine(v string) *GetQueryOptimizeDataTopRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetInstanceIds(v string) *GetQueryOptimizeDataTopRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetRegion(v string) *GetQueryOptimizeDataTopRequest {
	s.Region = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetTagNames(v string) *GetQueryOptimizeDataTopRequest {
	s.TagNames = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetTime(v string) *GetQueryOptimizeDataTopRequest {
	s.Time = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetType(v string) *GetQueryOptimizeDataTopRequest {
	s.Type = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) Validate() error {
	return dara.Validate(s)
}
