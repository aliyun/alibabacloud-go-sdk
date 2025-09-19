// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRatioStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceDirectoryAccountId(v int64) *GetClientRatioStatisticRequest
	GetResourceDirectoryAccountId() *int64
	SetStatisticTypes(v []*string) *GetClientRatioStatisticRequest
	GetStatisticTypes() []*string
	SetTimeEnd(v int64) *GetClientRatioStatisticRequest
	GetTimeEnd() *int64
	SetTimeStart(v int64) *GetClientRatioStatisticRequest
	GetTimeStart() *int64
}

type GetClientRatioStatisticRequest struct {
	// The ID of the primary account of the Resource Directory member account.
	//
	// > call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) interface to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// An array that consists of the details of a statistical type.
	StatisticTypes []*string `json:"StatisticTypes,omitempty" xml:"StatisticTypes,omitempty" type:"Repeated"`
	// The timestamp that specifies the end of the time range to collect statistics. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1686412799999
	TimeEnd *int64 `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	// The timestamp that specifies the beginning of the time range to collect statistics. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671382800000
	TimeStart *int64 `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
}

func (s GetClientRatioStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetClientRatioStatisticRequest) GetStatisticTypes() []*string {
	return s.StatisticTypes
}

func (s *GetClientRatioStatisticRequest) GetTimeEnd() *int64 {
	return s.TimeEnd
}

func (s *GetClientRatioStatisticRequest) GetTimeStart() *int64 {
	return s.TimeStart
}

func (s *GetClientRatioStatisticRequest) SetResourceDirectoryAccountId(v int64) *GetClientRatioStatisticRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetClientRatioStatisticRequest) SetStatisticTypes(v []*string) *GetClientRatioStatisticRequest {
	s.StatisticTypes = v
	return s
}

func (s *GetClientRatioStatisticRequest) SetTimeEnd(v int64) *GetClientRatioStatisticRequest {
	s.TimeEnd = &v
	return s
}

func (s *GetClientRatioStatisticRequest) SetTimeStart(v int64) *GetClientRatioStatisticRequest {
	s.TimeStart = &v
	return s
}

func (s *GetClientRatioStatisticRequest) Validate() error {
	return dara.Validate(s)
}
