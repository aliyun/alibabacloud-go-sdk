// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRechargeBillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanseInfos(v []*ListRechargeBillsResponseBodyInstanseInfos) *ListRechargeBillsResponseBody
	GetInstanseInfos() []*ListRechargeBillsResponseBodyInstanseInfos
	SetRequestId(v string) *ListRechargeBillsResponseBody
	GetRequestId() *string
	SetResidueAmount(v int32) *ListRechargeBillsResponseBody
	GetResidueAmount() *int32
	SetTotalCount(v int32) *ListRechargeBillsResponseBody
	GetTotalCount() *int32
}

type ListRechargeBillsResponseBody struct {
	InstanseInfos []*ListRechargeBillsResponseBodyInstanseInfos `json:"InstanseInfos,omitempty" xml:"InstanseInfos,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResidueAmount *int32                                        `json:"ResidueAmount,omitempty" xml:"ResidueAmount,omitempty"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRechargeBillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRechargeBillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsResponseBody) GetInstanseInfos() []*ListRechargeBillsResponseBodyInstanseInfos {
	return s.InstanseInfos
}

func (s *ListRechargeBillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRechargeBillsResponseBody) GetResidueAmount() *int32 {
	return s.ResidueAmount
}

func (s *ListRechargeBillsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRechargeBillsResponseBody) SetInstanseInfos(v []*ListRechargeBillsResponseBodyInstanseInfos) *ListRechargeBillsResponseBody {
	s.InstanseInfos = v
	return s
}

func (s *ListRechargeBillsResponseBody) SetRequestId(v string) *ListRechargeBillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRechargeBillsResponseBody) SetResidueAmount(v int32) *ListRechargeBillsResponseBody {
	s.ResidueAmount = &v
	return s
}

func (s *ListRechargeBillsResponseBody) SetTotalCount(v int32) *ListRechargeBillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRechargeBillsResponseBody) Validate() error {
	if s.InstanseInfos != nil {
		for _, item := range s.InstanseInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRechargeBillsResponseBodyInstanseInfos struct {
	CurrTimes  *int32  `json:"CurrTimes,omitempty" xml:"CurrTimes,omitempty"`
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	InitTimes  *int32  `json:"InitTimes,omitempty" xml:"InitTimes,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListRechargeBillsResponseBodyInstanseInfos) String() string {
	return dara.Prettify(s)
}

func (s ListRechargeBillsResponseBodyInstanseInfos) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) GetCurrTimes() *int32 {
	return s.CurrTimes
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) GetInitTimes() *int32 {
	return s.InitTimes
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) GetSource() *string {
	return s.Source
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetCurrTimes(v int32) *ListRechargeBillsResponseBodyInstanseInfos {
	s.CurrTimes = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetGmtEndTime(v string) *ListRechargeBillsResponseBodyInstanseInfos {
	s.GmtEndTime = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetInitTimes(v int32) *ListRechargeBillsResponseBodyInstanseInfos {
	s.InitTimes = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetInstanceId(v string) *ListRechargeBillsResponseBodyInstanseInfos {
	s.InstanceId = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetSource(v string) *ListRechargeBillsResponseBodyInstanseInfos {
	s.Source = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) Validate() error {
	return dara.Validate(s)
}
