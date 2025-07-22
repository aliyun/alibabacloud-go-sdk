// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *GetQueryOptimizeRuleListRequest
	GetEngine() *string
	SetInstanceIds(v string) *GetQueryOptimizeRuleListRequest
	GetInstanceIds() *string
	SetRegion(v string) *GetQueryOptimizeRuleListRequest
	GetRegion() *string
	SetTagNames(v string) *GetQueryOptimizeRuleListRequest
	GetTagNames() *string
}

type GetQueryOptimizeRuleListRequest struct {
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
	// >  If your instances reside in the regions in the Chinese mainland, set this parameter to **cn-china**.
	//
	// example:
	//
	// cn-china
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	TagNames *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
}

func (s GetQueryOptimizeRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeRuleListRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeRuleListRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetQueryOptimizeRuleListRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQueryOptimizeRuleListRequest) GetTagNames() *string {
	return s.TagNames
}

func (s *GetQueryOptimizeRuleListRequest) SetEngine(v string) *GetQueryOptimizeRuleListRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeRuleListRequest) SetInstanceIds(v string) *GetQueryOptimizeRuleListRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeRuleListRequest) SetRegion(v string) *GetQueryOptimizeRuleListRequest {
	s.Region = &v
	return s
}

func (s *GetQueryOptimizeRuleListRequest) SetTagNames(v string) *GetQueryOptimizeRuleListRequest {
	s.TagNames = &v
	return s
}

func (s *GetQueryOptimizeRuleListRequest) Validate() error {
	return dara.Validate(s)
}
