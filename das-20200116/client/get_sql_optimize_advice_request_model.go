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
	// A reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The end date of the query. Format: <i>yyyyMMdd</i> (UTC).
	//
	// - If this parameter is left empty, the default value is the day before the current date.
	//
	// - You can only query data from the day before the current date or earlier. The interval between the start date and the end date cannot exceed 30 days.
	//
	// example:
	//
	// 20210917
	EndDt *string `json:"EndDt,omitempty" xml:"EndDt,omitempty"`
	// The database engine. Valid values:
	//
	// - **MySQL**: RDS MySQL.
	//
	// - **PolarDBMySQL**: PolarDB for MySQL.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance ID.
	//
	// >Only RDS MySQL and PolarDB for MySQL instances are supported.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region to which the instance belongs. Valid values:
	//
	// - **cn-china**: the Chinese mainland.
	//
	// - **cn-hongkong**: Hong Kong (China).
	//
	// - **ap-southeast-1**: Singapore.
	//
	// This parameter takes effect only when the **InstanceIds*	- request parameter is left empty. If **InstanceIds*	- is left empty, data is retrieved based on the region specified by the **Region*	- parameter. The default region is **cn-china**. If **InstanceIds*	- is not empty, data is retrieved based on the region of the first instance specified by **InstanceIds**, even if the **Region*	- parameter is set.
	//
	// > For instances created in regions within the Chinese mainland, set this parameter to **cn-china**.
	//
	// example:
	//
	// cn-china
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start date of the query. Format: <i>yyyyMMdd</i> (UTC).
	//
	// - If this parameter is left empty, the default value is the day before the current date.
	//
	// - You can only query data from the day before the current date or earlier.
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
