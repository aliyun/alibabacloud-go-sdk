// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v string) *DescribeGtmAccessStrategyResponseBody
	GetAccessMode() *string
	SetAccessStatus(v string) *DescribeGtmAccessStrategyResponseBody
	GetAccessStatus() *string
	SetDefaultAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategyResponseBody
	GetDefaultAddrPoolMonitorStatus() *string
	SetDefaultAddrPoolName(v string) *DescribeGtmAccessStrategyResponseBody
	GetDefaultAddrPoolName() *string
	SetDefaultAddrPoolStatus(v string) *DescribeGtmAccessStrategyResponseBody
	GetDefaultAddrPoolStatus() *string
	SetDefultAddrPoolId(v string) *DescribeGtmAccessStrategyResponseBody
	GetDefultAddrPoolId() *string
	SetFailoverAddrPoolId(v string) *DescribeGtmAccessStrategyResponseBody
	GetFailoverAddrPoolId() *string
	SetFailoverAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategyResponseBody
	GetFailoverAddrPoolMonitorStatus() *string
	SetFailoverAddrPoolName(v string) *DescribeGtmAccessStrategyResponseBody
	GetFailoverAddrPoolName() *string
	SetFailoverAddrPoolStatus(v string) *DescribeGtmAccessStrategyResponseBody
	GetFailoverAddrPoolStatus() *string
	SetInstanceId(v string) *DescribeGtmAccessStrategyResponseBody
	GetInstanceId() *string
	SetLines(v *DescribeGtmAccessStrategyResponseBodyLines) *DescribeGtmAccessStrategyResponseBody
	GetLines() *DescribeGtmAccessStrategyResponseBodyLines
	SetRequestId(v string) *DescribeGtmAccessStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v string) *DescribeGtmAccessStrategyResponseBody
	GetStrategyId() *string
	SetStrategyMode(v string) *DescribeGtmAccessStrategyResponseBody
	GetStrategyMode() *string
	SetStrategyName(v string) *DescribeGtmAccessStrategyResponseBody
	GetStrategyName() *string
}

type DescribeGtmAccessStrategyResponseBody struct {
	// The access policy.
	//
	// example:
	//
	// AUTO
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The access status. Valid values:
	//
	// 	- **DEFAULT**: Indicates normal when the default address pool is accessed.
	//
	// 	- **FAILOVER**: Indicates an exception when a failover address pool is accessed.
	//
	// example:
	//
	// DEFAULT
	AccessStatus *string `json:"AccessStatus,omitempty" xml:"AccessStatus,omitempty"`
	// Indicates whether health check is enabled for the default address pool.
	//
	// example:
	//
	// OPEN
	DefaultAddrPoolMonitorStatus *string `json:"DefaultAddrPoolMonitorStatus,omitempty" xml:"DefaultAddrPoolMonitorStatus,omitempty"`
	// The name of the default address pool.
	DefaultAddrPoolName *string `json:"DefaultAddrPoolName,omitempty" xml:"DefaultAddrPoolName,omitempty"`
	// The availability status of the default address pool.
	//
	// example:
	//
	// AVAILABLE
	DefaultAddrPoolStatus *string `json:"DefaultAddrPoolStatus,omitempty" xml:"DefaultAddrPoolStatus,omitempty"`
	// The ID of the default address pool.
	//
	// example:
	//
	// hra0i1
	DefultAddrPoolId *string `json:"DefultAddrPoolId,omitempty" xml:"DefultAddrPoolId,omitempty"`
	// The ID of the failover address pool.
	//
	// example:
	//
	// hra0i2
	FailoverAddrPoolId *string `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	// Indicates whether health check is enabled for the failover address pool.
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
	// The ID of the GTM instance whose access policy details you want to query.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The returned lines of access regions.
	Lines *DescribeGtmAccessStrategyResponseBodyLines `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BA1608CA-834C-4E63-8682-8AF0B11ED72D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the access policy queried.
	//
	// example:
	//
	// hra0hs
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The mode of traffic scheduling.
	//
	// example:
	//
	// SELF_DEFINED
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	// The name of the access policy queried.
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s DescribeGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyResponseBody) GetAccessMode() *string {
	return s.AccessMode
}

func (s *DescribeGtmAccessStrategyResponseBody) GetAccessStatus() *string {
	return s.AccessStatus
}

func (s *DescribeGtmAccessStrategyResponseBody) GetDefaultAddrPoolMonitorStatus() *string {
	return s.DefaultAddrPoolMonitorStatus
}

func (s *DescribeGtmAccessStrategyResponseBody) GetDefaultAddrPoolName() *string {
	return s.DefaultAddrPoolName
}

func (s *DescribeGtmAccessStrategyResponseBody) GetDefaultAddrPoolStatus() *string {
	return s.DefaultAddrPoolStatus
}

func (s *DescribeGtmAccessStrategyResponseBody) GetDefultAddrPoolId() *string {
	return s.DefultAddrPoolId
}

func (s *DescribeGtmAccessStrategyResponseBody) GetFailoverAddrPoolId() *string {
	return s.FailoverAddrPoolId
}

func (s *DescribeGtmAccessStrategyResponseBody) GetFailoverAddrPoolMonitorStatus() *string {
	return s.FailoverAddrPoolMonitorStatus
}

func (s *DescribeGtmAccessStrategyResponseBody) GetFailoverAddrPoolName() *string {
	return s.FailoverAddrPoolName
}

func (s *DescribeGtmAccessStrategyResponseBody) GetFailoverAddrPoolStatus() *string {
	return s.FailoverAddrPoolStatus
}

func (s *DescribeGtmAccessStrategyResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmAccessStrategyResponseBody) GetLines() *DescribeGtmAccessStrategyResponseBodyLines {
	return s.Lines
}

func (s *DescribeGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmAccessStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeGtmAccessStrategyResponseBody) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *DescribeGtmAccessStrategyResponseBody) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeGtmAccessStrategyResponseBody) SetAccessMode(v string) *DescribeGtmAccessStrategyResponseBody {
	s.AccessMode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetAccessStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.AccessStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefaultAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefaultAddrPoolName(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefaultAddrPoolStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetDefultAddrPoolId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.DefultAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolMonitorStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolMonitorStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolName(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetFailoverAddrPoolStatus(v string) *DescribeGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolStatus = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetInstanceId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetLines(v *DescribeGtmAccessStrategyResponseBodyLines) *DescribeGtmAccessStrategyResponseBody {
	s.Lines = v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetRequestId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetStrategyId(v string) *DescribeGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetStrategyMode(v string) *DescribeGtmAccessStrategyResponseBody {
	s.StrategyMode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) SetStrategyName(v string) *DescribeGtmAccessStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBody) Validate() error {
	if s.Lines != nil {
		if err := s.Lines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGtmAccessStrategyResponseBodyLines struct {
	Line []*DescribeGtmAccessStrategyResponseBodyLinesLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s DescribeGtmAccessStrategyResponseBodyLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyResponseBodyLines) GetLine() []*DescribeGtmAccessStrategyResponseBodyLinesLine {
	return s.Line
}

func (s *DescribeGtmAccessStrategyResponseBodyLines) SetLine(v []*DescribeGtmAccessStrategyResponseBodyLinesLine) *DescribeGtmAccessStrategyResponseBodyLines {
	s.Line = v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLines) Validate() error {
	if s.Line != nil {
		for _, item := range s.Line {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGtmAccessStrategyResponseBodyLinesLine struct {
	// The code of the access region group.
	//
	// example:
	//
	// DEFAULT
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the access region group.
	//
	// example:
	//
	// Global
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The code for the line of the access region.
	//
	// example:
	//
	// default
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The name for the line of the access region.
	//
	// example:
	//
	// Global
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeGtmAccessStrategyResponseBodyLinesLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyResponseBodyLinesLine) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) SetGroupCode(v string) *DescribeGtmAccessStrategyResponseBodyLinesLine {
	s.GroupCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) SetGroupName(v string) *DescribeGtmAccessStrategyResponseBodyLinesLine {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) SetLineCode(v string) *DescribeGtmAccessStrategyResponseBodyLinesLine {
	s.LineCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) SetLineName(v string) *DescribeGtmAccessStrategyResponseBodyLinesLine {
	s.LineName = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponseBodyLinesLine) Validate() error {
	return dara.Validate(s)
}
