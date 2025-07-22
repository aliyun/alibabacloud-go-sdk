// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlOptimizeAdviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetSqlOptimizeAdviceRequest
	GetConsoleContext() *string
	SetEndDt(v string) *GetSqlOptimizeAdviceRequest
	GetEndDt() *string
	SetEngine(v string) *GetSqlOptimizeAdviceRequest
	GetEngine() *string
	SetInstanceIds(v string) *GetSqlOptimizeAdviceRequest
	GetInstanceIds() *string
	SetRegion(v string) *GetSqlOptimizeAdviceRequest
	GetRegion() *string
	SetStartDt(v string) *GetSqlOptimizeAdviceRequest
	GetStartDt() *string
}

type GetSqlOptimizeAdviceRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The end date of the time range to query. Specify the date in the *yyyyMMdd	- format. The time must be in UTC.
	//
	// 	- The default value of this parameter is one day before the current day.
	//
	// 	- The value must be earlier than the current day. The interval between the start date and the end date cannot exceed 30 days.
	//
	// example:
	//
	// 20210917
	EndDt *string `json:"EndDt,omitempty" xml:"EndDt,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**: ApsaraDB RDS for MySQL.
	//
	// 	- **PolarDBMySQL**: PolarDB for MySQL.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance ID.
	//
	// >  You must specify the instance ID only if your database instance is an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region in which the instance resides. Valid values:
	//
	// 	- **cn-china**: Chinese mainland.
	//
	// 	- **cn-hongkong**: China (Hong Kong).
	//
	// 	- **ap-southeast-1**: Singapore.
	//
	// This parameter takes effect only if **InstanceIds*	- is left empty. If you leave **InstanceIds*	- empty, the system obtains data from the region specified by **Region**. By default, Region is set to **cn-china**. If you specify **InstanceIds**, **Region*	- does not take effect, and the system obtains data from the region in which the first specified instance resides.****
	//
	// >  If your instances reside in the regions inside the Chinese mainland, set this parameter to **cn-china**.
	//
	// example:
	//
	// cn-china
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start date of the time range to query. Specify the date in the *yyyyMMdd	- format. The time must be in UTC.
	//
	// 	- The default value of this parameter is one day before the current day.
	//
	// 	- The value must be earlier than the current day.
	//
	// example:
	//
	// 20210916
	StartDt *string `json:"StartDt,omitempty" xml:"StartDt,omitempty"`
}

func (s GetSqlOptimizeAdviceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlOptimizeAdviceRequest) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetSqlOptimizeAdviceRequest) GetEndDt() *string {
	return s.EndDt
}

func (s *GetSqlOptimizeAdviceRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetSqlOptimizeAdviceRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetSqlOptimizeAdviceRequest) GetRegion() *string {
	return s.Region
}

func (s *GetSqlOptimizeAdviceRequest) GetStartDt() *string {
	return s.StartDt
}

func (s *GetSqlOptimizeAdviceRequest) SetConsoleContext(v string) *GetSqlOptimizeAdviceRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetEndDt(v string) *GetSqlOptimizeAdviceRequest {
	s.EndDt = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetEngine(v string) *GetSqlOptimizeAdviceRequest {
	s.Engine = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetInstanceIds(v string) *GetSqlOptimizeAdviceRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetRegion(v string) *GetSqlOptimizeAdviceRequest {
	s.Region = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetStartDt(v string) *GetSqlOptimizeAdviceRequest {
	s.StartDt = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) Validate() error {
	return dara.Validate(s)
}
