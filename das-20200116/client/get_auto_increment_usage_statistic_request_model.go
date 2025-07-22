// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoIncrementUsageStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbNames(v string) *GetAutoIncrementUsageStatisticRequest
	GetDbNames() *string
	SetInstanceId(v string) *GetAutoIncrementUsageStatisticRequest
	GetInstanceId() *string
	SetRatioFilter(v float64) *GetAutoIncrementUsageStatisticRequest
	GetRatioFilter() *float64
	SetRealTime(v bool) *GetAutoIncrementUsageStatisticRequest
	GetRealTime() *bool
}

type GetAutoIncrementUsageStatisticRequest struct {
	// The database name. If you specify a database, the operation queries the usage of auto-increment table IDs in the specified database. Otherwise, the operation queries the usage of auto-increment table IDs in all databases on the instance.
	//
	// >  Specify the parameter value as a JSON array, such as [\\"db1\\",\\"db2\\"]. Separate multiple database names with commas (,).
	//
	// example:
	//
	// [\\"db1\\",\\"db2\\"]
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The usage threshold of auto-increment IDs. Only usage that exceeds the threshold can be returned. Valid values are decimals that range from 0 to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.9
	RatioFilter *float64 `json:"RatioFilter,omitempty" xml:"RatioFilter,omitempty"`
	// Specifies whether to query real-time data. Valid values:
	//
	// 	- **true**: queries data in real time except for data generated in the last 10 minutes.****
	//
	// 	- **false**: queries data generated in the last 2 hours. If no such data exists, queries the latest data.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	RealTime *bool `json:"RealTime,omitempty" xml:"RealTime,omitempty"`
}

func (s GetAutoIncrementUsageStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoIncrementUsageStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetAutoIncrementUsageStatisticRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *GetAutoIncrementUsageStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoIncrementUsageStatisticRequest) GetRatioFilter() *float64 {
	return s.RatioFilter
}

func (s *GetAutoIncrementUsageStatisticRequest) GetRealTime() *bool {
	return s.RealTime
}

func (s *GetAutoIncrementUsageStatisticRequest) SetDbNames(v string) *GetAutoIncrementUsageStatisticRequest {
	s.DbNames = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticRequest) SetInstanceId(v string) *GetAutoIncrementUsageStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticRequest) SetRatioFilter(v float64) *GetAutoIncrementUsageStatisticRequest {
	s.RatioFilter = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticRequest) SetRealTime(v bool) *GetAutoIncrementUsageStatisticRequest {
	s.RealTime = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticRequest) Validate() error {
	return dara.Validate(s)
}
