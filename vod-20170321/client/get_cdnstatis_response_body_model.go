// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCDNStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCDNStatisList(v *GetCDNStatisResponseBodyCDNStatisList) *GetCDNStatisResponseBody
	GetCDNStatisList() *GetCDNStatisResponseBodyCDNStatisList
	SetMaxBpsDataValue(v int64) *GetCDNStatisResponseBody
	GetMaxBpsDataValue() *int64
	SetRequestId(v string) *GetCDNStatisResponseBody
	GetRequestId() *string
	SetSumFlowDataValue(v int64) *GetCDNStatisResponseBody
	GetSumFlowDataValue() *int64
}

type GetCDNStatisResponseBody struct {
	CDNStatisList    *GetCDNStatisResponseBodyCDNStatisList `json:"CDNStatisList,omitempty" xml:"CDNStatisList,omitempty" type:"Struct"`
	MaxBpsDataValue  *int64                                 `json:"MaxBpsDataValue,omitempty" xml:"MaxBpsDataValue,omitempty"`
	RequestId        *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SumFlowDataValue *int64                                 `json:"SumFlowDataValue,omitempty" xml:"SumFlowDataValue,omitempty"`
}

func (s GetCDNStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetCDNStatisResponseBody) GetCDNStatisList() *GetCDNStatisResponseBodyCDNStatisList {
	return s.CDNStatisList
}

func (s *GetCDNStatisResponseBody) GetMaxBpsDataValue() *int64 {
	return s.MaxBpsDataValue
}

func (s *GetCDNStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCDNStatisResponseBody) GetSumFlowDataValue() *int64 {
	return s.SumFlowDataValue
}

func (s *GetCDNStatisResponseBody) SetCDNStatisList(v *GetCDNStatisResponseBodyCDNStatisList) *GetCDNStatisResponseBody {
	s.CDNStatisList = v
	return s
}

func (s *GetCDNStatisResponseBody) SetMaxBpsDataValue(v int64) *GetCDNStatisResponseBody {
	s.MaxBpsDataValue = &v
	return s
}

func (s *GetCDNStatisResponseBody) SetRequestId(v string) *GetCDNStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCDNStatisResponseBody) SetSumFlowDataValue(v int64) *GetCDNStatisResponseBody {
	s.SumFlowDataValue = &v
	return s
}

func (s *GetCDNStatisResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCDNStatisResponseBodyCDNStatisList struct {
	CDNMetric []*GetCDNStatisResponseBodyCDNStatisListCDNMetric `json:"CDNMetric,omitempty" xml:"CDNMetric,omitempty" type:"Repeated"`
}

func (s GetCDNStatisResponseBodyCDNStatisList) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisResponseBodyCDNStatisList) GoString() string {
	return s.String()
}

func (s *GetCDNStatisResponseBodyCDNStatisList) GetCDNMetric() []*GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	return s.CDNMetric
}

func (s *GetCDNStatisResponseBodyCDNStatisList) SetCDNMetric(v []*GetCDNStatisResponseBodyCDNStatisListCDNMetric) *GetCDNStatisResponseBodyCDNStatisList {
	s.CDNMetric = v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisList) Validate() error {
	return dara.Validate(s)
}

type GetCDNStatisResponseBodyCDNStatisListCDNMetric struct {
	BpsDataDomesticValue  *int64  `json:"BpsDataDomesticValue,omitempty" xml:"BpsDataDomesticValue,omitempty"`
	BpsDataOverseasValue  *int64  `json:"BpsDataOverseasValue,omitempty" xml:"BpsDataOverseasValue,omitempty"`
	BpsDataValue          *int64  `json:"BpsDataValue,omitempty" xml:"BpsDataValue,omitempty"`
	CdnDomain             *string `json:"CdnDomain,omitempty" xml:"CdnDomain,omitempty"`
	FlowDataDomesticValue *int64  `json:"FlowDataDomesticValue,omitempty" xml:"FlowDataDomesticValue,omitempty"`
	FlowDataOverseasValue *int64  `json:"FlowDataOverseasValue,omitempty" xml:"FlowDataOverseasValue,omitempty"`
	FlowDataValue         *int64  `json:"FlowDataValue,omitempty" xml:"FlowDataValue,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	StatTime              *string `json:"StatTime,omitempty" xml:"StatTime,omitempty"`
}

func (s GetCDNStatisResponseBodyCDNStatisListCDNMetric) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisResponseBodyCDNStatisListCDNMetric) GoString() string {
	return s.String()
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetBpsDataDomesticValue() *int64 {
	return s.BpsDataDomesticValue
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetBpsDataOverseasValue() *int64 {
	return s.BpsDataOverseasValue
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetBpsDataValue() *int64 {
	return s.BpsDataValue
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetCdnDomain() *string {
	return s.CdnDomain
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetFlowDataDomesticValue() *int64 {
	return s.FlowDataDomesticValue
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetFlowDataOverseasValue() *int64 {
	return s.FlowDataOverseasValue
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetFlowDataValue() *int64 {
	return s.FlowDataValue
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetId() *int64 {
	return s.Id
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) GetStatTime() *string {
	return s.StatTime
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetBpsDataDomesticValue(v int64) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.BpsDataDomesticValue = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetBpsDataOverseasValue(v int64) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.BpsDataOverseasValue = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetBpsDataValue(v int64) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.BpsDataValue = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetCdnDomain(v string) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.CdnDomain = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetFlowDataDomesticValue(v int64) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.FlowDataDomesticValue = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetFlowDataOverseasValue(v int64) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.FlowDataOverseasValue = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetFlowDataValue(v int64) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.FlowDataValue = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetId(v int64) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.Id = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) SetStatTime(v string) *GetCDNStatisResponseBodyCDNStatisListCDNMetric {
	s.StatTime = &v
	return s
}

func (s *GetCDNStatisResponseBodyCDNStatisListCDNMetric) Validate() error {
	return dara.Validate(s)
}
