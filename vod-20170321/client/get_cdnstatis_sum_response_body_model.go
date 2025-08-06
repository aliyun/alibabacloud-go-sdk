// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCDNStatisSumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCDNStatisList(v *GetCDNStatisSumResponseBodyCDNStatisList) *GetCDNStatisSumResponseBody
	GetCDNStatisList() *GetCDNStatisSumResponseBodyCDNStatisList
	SetMaxBpsDataValue(v int64) *GetCDNStatisSumResponseBody
	GetMaxBpsDataValue() *int64
	SetRequestId(v string) *GetCDNStatisSumResponseBody
	GetRequestId() *string
	SetSumFlowDataValue(v int64) *GetCDNStatisSumResponseBody
	GetSumFlowDataValue() *int64
}

type GetCDNStatisSumResponseBody struct {
	CDNStatisList    *GetCDNStatisSumResponseBodyCDNStatisList `json:"CDNStatisList,omitempty" xml:"CDNStatisList,omitempty" type:"Struct"`
	MaxBpsDataValue  *int64                                    `json:"MaxBpsDataValue,omitempty" xml:"MaxBpsDataValue,omitempty"`
	RequestId        *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SumFlowDataValue *int64                                    `json:"SumFlowDataValue,omitempty" xml:"SumFlowDataValue,omitempty"`
}

func (s GetCDNStatisSumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisSumResponseBody) GoString() string {
	return s.String()
}

func (s *GetCDNStatisSumResponseBody) GetCDNStatisList() *GetCDNStatisSumResponseBodyCDNStatisList {
	return s.CDNStatisList
}

func (s *GetCDNStatisSumResponseBody) GetMaxBpsDataValue() *int64 {
	return s.MaxBpsDataValue
}

func (s *GetCDNStatisSumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCDNStatisSumResponseBody) GetSumFlowDataValue() *int64 {
	return s.SumFlowDataValue
}

func (s *GetCDNStatisSumResponseBody) SetCDNStatisList(v *GetCDNStatisSumResponseBodyCDNStatisList) *GetCDNStatisSumResponseBody {
	s.CDNStatisList = v
	return s
}

func (s *GetCDNStatisSumResponseBody) SetMaxBpsDataValue(v int64) *GetCDNStatisSumResponseBody {
	s.MaxBpsDataValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBody) SetRequestId(v string) *GetCDNStatisSumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCDNStatisSumResponseBody) SetSumFlowDataValue(v int64) *GetCDNStatisSumResponseBody {
	s.SumFlowDataValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCDNStatisSumResponseBodyCDNStatisList struct {
	CDNMetric []*GetCDNStatisSumResponseBodyCDNStatisListCDNMetric `json:"CDNMetric,omitempty" xml:"CDNMetric,omitempty" type:"Repeated"`
}

func (s GetCDNStatisSumResponseBodyCDNStatisList) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisSumResponseBodyCDNStatisList) GoString() string {
	return s.String()
}

func (s *GetCDNStatisSumResponseBodyCDNStatisList) GetCDNMetric() []*GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	return s.CDNMetric
}

func (s *GetCDNStatisSumResponseBodyCDNStatisList) SetCDNMetric(v []*GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) *GetCDNStatisSumResponseBodyCDNStatisList {
	s.CDNMetric = v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisList) Validate() error {
	return dara.Validate(s)
}

type GetCDNStatisSumResponseBodyCDNStatisListCDNMetric struct {
	BpsDataDomesticValue  *int64  `json:"BpsDataDomesticValue,omitempty" xml:"BpsDataDomesticValue,omitempty"`
	BpsDataOverseasValue  *int64  `json:"BpsDataOverseasValue,omitempty" xml:"BpsDataOverseasValue,omitempty"`
	BpsDataValue          *int64  `json:"BpsDataValue,omitempty" xml:"BpsDataValue,omitempty"`
	FlowDataDomesticValue *int64  `json:"FlowDataDomesticValue,omitempty" xml:"FlowDataDomesticValue,omitempty"`
	FlowDataOverseasValue *int64  `json:"FlowDataOverseasValue,omitempty" xml:"FlowDataOverseasValue,omitempty"`
	FlowDataValue         *int64  `json:"FlowDataValue,omitempty" xml:"FlowDataValue,omitempty"`
	StatTime              *string `json:"StatTime,omitempty" xml:"StatTime,omitempty"`
}

func (s GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GoString() string {
	return s.String()
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GetBpsDataDomesticValue() *int64 {
	return s.BpsDataDomesticValue
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GetBpsDataOverseasValue() *int64 {
	return s.BpsDataOverseasValue
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GetBpsDataValue() *int64 {
	return s.BpsDataValue
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GetFlowDataDomesticValue() *int64 {
	return s.FlowDataDomesticValue
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GetFlowDataOverseasValue() *int64 {
	return s.FlowDataOverseasValue
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GetFlowDataValue() *int64 {
	return s.FlowDataValue
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) GetStatTime() *string {
	return s.StatTime
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) SetBpsDataDomesticValue(v int64) *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	s.BpsDataDomesticValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) SetBpsDataOverseasValue(v int64) *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	s.BpsDataOverseasValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) SetBpsDataValue(v int64) *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	s.BpsDataValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) SetFlowDataDomesticValue(v int64) *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	s.FlowDataDomesticValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) SetFlowDataOverseasValue(v int64) *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	s.FlowDataOverseasValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) SetFlowDataValue(v int64) *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	s.FlowDataValue = &v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) SetStatTime(v string) *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric {
	s.StatTime = &v
	return s
}

func (s *GetCDNStatisSumResponseBodyCDNStatisListCDNMetric) Validate() error {
	return dara.Validate(s)
}
