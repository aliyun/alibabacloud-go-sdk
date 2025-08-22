// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaConnDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionDataPerInterval(v *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval) *DescribeDcdnDomainIpaConnDataResponseBody
	GetConnectionDataPerInterval() *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval
	SetEndTime(v string) *DescribeDcdnDomainIpaConnDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainIpaConnDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainIpaConnDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainIpaConnDataResponseBody struct {
	// The number of user connections at each time interval.
	ConnectionDataPerInterval *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval `json:"ConnectionDataPerInterval,omitempty" xml:"ConnectionDataPerInterval,omitempty" type:"Struct"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-02-22T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A2A1EEF8-043E-43A1-807C-BEAC18EA1807
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2015-02-21T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainIpaConnDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) GetConnectionDataPerInterval() *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval {
	return s.ConnectionDataPerInterval
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) SetConnectionDataPerInterval(v *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval) *DescribeDcdnDomainIpaConnDataResponseBody {
	s.ConnectionDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainIpaConnDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainIpaConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainIpaConnDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval struct {
	DataModule []*DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval) GetDataModule() []*DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval) SetDataModule(v []*DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule struct {
	// The number of IPA user connections.
	//
	// example:
	//
	// 189095
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example1.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2022-02-21T15:00:00+08:00
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) GetConnections() *int64 {
	return s.Connections
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) SetConnections(v int64) *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule {
	s.Connections = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) SetDomain(v string) *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponseBodyConnectionDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
