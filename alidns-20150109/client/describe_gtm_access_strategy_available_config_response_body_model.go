// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategyAvailableConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPools(v *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) *DescribeGtmAccessStrategyAvailableConfigResponseBody
	GetAddrPools() *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools
	SetLines(v *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) *DescribeGtmAccessStrategyAvailableConfigResponseBody
	GetLines() *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines
	SetRequestId(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBody
	GetRequestId() *string
	SetSuggestSetDefaultLine(v bool) *DescribeGtmAccessStrategyAvailableConfigResponseBody
	GetSuggestSetDefaultLine() *bool
}

type DescribeGtmAccessStrategyAvailableConfigResponseBody struct {
	// The address pools.
	AddrPools *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Struct"`
	// The Domain Name System (DNS) request sources.
	Lines *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C2851BA9-CE56-49AF-8D12-4FC6A49EE688
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the global line is recommended.
	//
	// example:
	//
	// True
	SuggestSetDefaultLine *bool `json:"SuggestSetDefaultLine,omitempty" xml:"SuggestSetDefaultLine,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) GetAddrPools() *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools {
	return s.AddrPools
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) GetLines() *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	return s.Lines
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) GetSuggestSetDefaultLine() *bool {
	return s.SuggestSetDefaultLine
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) SetAddrPools(v *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	s.AddrPools = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) SetLines(v *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	s.Lines = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) SetRequestId(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) SetSuggestSetDefaultLine(v bool) *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	s.SuggestSetDefaultLine = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools struct {
	AddrPool []*DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool `json:"AddrPool,omitempty" xml:"AddrPool,omitempty" type:"Repeated"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) GetAddrPool() []*DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool {
	return s.AddrPool
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) SetAddrPool(v []*DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools {
	s.AddrPool = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPools) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool struct {
	// The ID of the address pool.
	//
	// example:
	//
	// hra0ix
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// test
	AddrPoolName *string `json:"AddrPoolName,omitempty" xml:"AddrPoolName,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) GetAddrPoolName() *string {
	return s.AddrPoolName
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) SetAddrPoolId(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) SetAddrPoolName(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool {
	s.AddrPoolName = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyAddrPoolsAddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategyAvailableConfigResponseBodyLines struct {
	Line []*DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) GetLine() []*DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	return s.Line
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) SetLine(v []*DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.Line = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLines) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine struct {
	// The code of the parent line. No value is returned if no parent line exists.
	//
	// example:
	//
	// telecom
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	// The group number of the DNS request source.
	//
	// example:
	//
	// ISP
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The group name of the DNS request source.
	//
	// example:
	//
	// Mainland China
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The code of the DNS request source.
	//
	// example:
	//
	// cn_telecom_hubei
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The name of the DNS request source.
	//
	// example:
	//
	// South China
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
	// The state of the line. Valid values:
	//
	// 	- **FORBIDDEN**: The line is unavailable.
	//
	// 	- **OPTIONAL**: The line is available.
	//
	// example:
	//
	// FORBIDDEN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetFatherCode() *string {
	return s.FatherCode
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetStatus() *string {
	return s.Status
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetFatherCode(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.FatherCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetGroupCode(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.GroupCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetGroupName(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetLineCode(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.LineCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetLineName(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.LineName = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetStatus(v string) *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.Status = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponseBodyLinesLine) Validate() error {
	return dara.Validate(s)
}
