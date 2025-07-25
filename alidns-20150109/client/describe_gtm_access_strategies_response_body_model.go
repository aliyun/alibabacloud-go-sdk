// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeGtmAccessStrategiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmAccessStrategiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGtmAccessStrategiesResponseBody
	GetRequestId() *string
	SetStrategies(v *DescribeGtmAccessStrategiesResponseBodyStrategies) *DescribeGtmAccessStrategiesResponseBody
	GetStrategies() *DescribeGtmAccessStrategiesResponseBodyStrategies
	SetTotalItems(v int32) *DescribeGtmAccessStrategiesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeGtmAccessStrategiesResponseBody
	GetTotalPages() *int32
}

type DescribeGtmAccessStrategiesResponseBody struct {
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0CCC9971-CEC9-4132-824B-4AE611C07623
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned list of access policies of the GTM instance.
	Strategies *DescribeGtmAccessStrategiesResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Struct"`
	// The total number of entries returned on all pages.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeGtmAccessStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmAccessStrategiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmAccessStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmAccessStrategiesResponseBody) GetStrategies() *DescribeGtmAccessStrategiesResponseBodyStrategies {
	return s.Strategies
}

func (s *DescribeGtmAccessStrategiesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeGtmAccessStrategiesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetPageNumber(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetPageSize(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetRequestId(v string) *DescribeGtmAccessStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetStrategies(v *DescribeGtmAccessStrategiesResponseBodyStrategies) *DescribeGtmAccessStrategiesResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetTotalItems(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) SetTotalPages(v int32) *DescribeGtmAccessStrategiesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategiesResponseBodyStrategies struct {
	Strategy []*DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Repeated"`
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) GetStrategy() []*DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	return s.Strategy
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) SetStrategy(v []*DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) *DescribeGtmAccessStrategiesResponseBodyStrategies {
	s.Strategy = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategies) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy struct {
	// The access policy. Valid values:
	//
	// 	- **AUTO**: Automatic switch
	//
	// 	- **DEFAULT**: Default address pool
	//
	// 	- **FAILOVER**: Failover address pool
	//
	// example:
	//
	// DEFAULT
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The access status. Valid values:
	//
	// 	- **DEFAULT**: The default address pool is currently accessed.
	//
	// 	- **FAILOVER**: The failover address pool is currently accessed.
	//
	// example:
	//
	// DEFAULT
	AccessStatus *string `json:"AccessStatus,omitempty" xml:"AccessStatus,omitempty"`
	// The time when the access policy was created.
	//
	// example:
	//
	// 2018-08-09T00:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533773400000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The ID of the default address pool.
	//
	// example:
	//
	// hra0i1
	DefaultAddrPoolId *string `json:"DefaultAddrPoolId,omitempty" xml:"DefaultAddrPoolId,omitempty"`
	// Indicates whether health check was enabled for the default address pool. Valid values:
	//
	// 	- **OPEN**: Enabled
	//
	// 	- **CLOSE**: Disabled
	//
	// 	- **UNCONFIGURED**: Not configured
	//
	// example:
	//
	// OPEN
	DefaultAddrPoolMonitorStatus *string `json:"DefaultAddrPoolMonitorStatus,omitempty" xml:"DefaultAddrPoolMonitorStatus,omitempty"`
	// The name of the default address pool.
	DefaultAddrPoolName *string `json:"DefaultAddrPoolName,omitempty" xml:"DefaultAddrPoolName,omitempty"`
	// The availability status of the default address pool. Valid values:
	//
	// 	- **AVAILABLE**: Available
	//
	// 	- **NOT_AVAILABLE**: Unavailable
	//
	// example:
	//
	// AVAILABLE
	DefaultAddrPoolStatus *string `json:"DefaultAddrPoolStatus,omitempty" xml:"DefaultAddrPoolStatus,omitempty"`
	// The ID of the failover address pool.
	//
	// example:
	//
	// hra0i2
	FailoverAddrPoolId *string `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	// Indicates whether health check was enabled for the failover address pool.
	//
	// example:
	//
	// OPEN
	FailoverAddrPoolMonitorStatus *string `json:"FailoverAddrPoolMonitorStatus,omitempty" xml:"FailoverAddrPoolMonitorStatus,omitempty"`
	// The name of the failover address pool.
	FailoverAddrPoolName *string `json:"FailoverAddrPoolName,omitempty" xml:"FailoverAddrPoolName,omitempty"`
	// The availability status of the failover address pool.
	//
	// example:
	//
	// AVAILABLE
	FailoverAddrPoolStatus *string `json:"FailoverAddrPoolStatus,omitempty" xml:"FailoverAddrPoolStatus,omitempty"`
	// The ID of the GTM instance whose access policies you want to query.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The returned lines of access regions.
	Lines *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Struct"`
	// The ID of the access policy.
	//
	// example:
	//
	// hra0hs
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The mode of the access policy. **SELF_DEFINED*	- indicates that the access policy is user-defined.
	//
	// example:
	//
	// SELF_DEFINED
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	// The name of the access policy.
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetAccessMode() *string {
	return s.AccessMode
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetAccessStatus() *string {
	return s.AccessStatus
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetDefaultAddrPoolId() *string {
	return s.DefaultAddrPoolId
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetDefaultAddrPoolMonitorStatus() *string {
	return s.DefaultAddrPoolMonitorStatus
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetDefaultAddrPoolName() *string {
	return s.DefaultAddrPoolName
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetDefaultAddrPoolStatus() *string {
	return s.DefaultAddrPoolStatus
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetFailoverAddrPoolId() *string {
	return s.FailoverAddrPoolId
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetFailoverAddrPoolMonitorStatus() *string {
	return s.FailoverAddrPoolMonitorStatus
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetFailoverAddrPoolName() *string {
	return s.FailoverAddrPoolName
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetFailoverAddrPoolStatus() *string {
	return s.FailoverAddrPoolStatus
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetLines() *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines {
	return s.Lines
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetAccessMode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.AccessMode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetAccessStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.AccessStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetCreateTime(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetCreateTimestamp(v int64) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetDefaultAddrPoolId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.DefaultAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetDefaultAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.DefaultAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetDefaultAddrPoolName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.DefaultAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetDefaultAddrPoolStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.DefaultAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetFailoverAddrPoolId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetFailoverAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.FailoverAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetFailoverAddrPoolName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.FailoverAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetFailoverAddrPoolStatus(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.FailoverAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetInstanceId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetLines(v *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.Lines = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetStrategyId(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.StrategyId = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetStrategyMode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.StrategyMode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) SetStrategyName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.StrategyName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategy) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines struct {
	Line []*DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines) GetLine() []*DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	return s.Line
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines) SetLine(v []*DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines {
	s.Line = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLines) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine struct {
	// The code of the access region group.
	//
	// example:
	//
	// DEFAULT
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the access region group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The code for the line of the access region.
	//
	// example:
	//
	// default
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The name for the line of the access region.
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetGroupCode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.GroupCode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetGroupName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetLineCode(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.LineCode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetLineName(v string) *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.LineName = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) Validate() error {
	return dara.Validate(s)
}
