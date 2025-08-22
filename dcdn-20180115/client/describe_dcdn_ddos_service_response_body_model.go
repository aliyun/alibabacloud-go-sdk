// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDdosServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangingAffectTime(v string) *DescribeDcdnDdosServiceResponseBody
	GetChangingAffectTime() *string
	SetChangingChargeType(v string) *DescribeDcdnDdosServiceResponseBody
	GetChangingChargeType() *string
	SetChangingDomianNum(v int32) *DescribeDcdnDdosServiceResponseBody
	GetChangingDomianNum() *int32
	SetChangingEdition(v string) *DescribeDcdnDdosServiceResponseBody
	GetChangingEdition() *string
	SetChangingProtectNum(v int32) *DescribeDcdnDdosServiceResponseBody
	GetChangingProtectNum() *int32
	SetChargeType(v string) *DescribeDcdnDdosServiceResponseBody
	GetChargeType() *string
	SetDomianNum(v int32) *DescribeDcdnDdosServiceResponseBody
	GetDomianNum() *int32
	SetEdition(v string) *DescribeDcdnDdosServiceResponseBody
	GetEdition() *string
	SetEnabled(v string) *DescribeDcdnDdosServiceResponseBody
	GetEnabled() *string
	SetEndingTime(v string) *DescribeDcdnDdosServiceResponseBody
	GetEndingTime() *string
	SetInstanceId(v string) *DescribeDcdnDdosServiceResponseBody
	GetInstanceId() *string
	SetOpeningTime(v string) *DescribeDcdnDdosServiceResponseBody
	GetOpeningTime() *string
	SetProtectNum(v int32) *DescribeDcdnDdosServiceResponseBody
	GetProtectNum() *int32
	SetRequestId(v string) *DescribeDcdnDdosServiceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDcdnDdosServiceResponseBody
	GetStatus() *string
}

type DescribeDcdnDdosServiceResponseBody struct {
	// The time when the renewed service takes effect. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-03-31T16:00:00Z
	ChangingAffectTime *string `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	// The metering method after the configuration changes Valid values:
	//
	// 	- **PayByBandwidth**
	//
	// 	- **PayByTraffic**
	//
	// 	- **PayByBandwidth95**
	//
	// example:
	//
	// PayByBandwidth
	ChangingChargeType *string `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	// The number of protected domain names.
	//
	// example:
	//
	// 30
	ChangingDomianNum *int32 `json:"ChangingDomianNum,omitempty" xml:"ChangingDomianNum,omitempty"`
	// The protection edition for which the configuration changes take effect. Valid values:
	//
	// 	- **poc**: POC Edition
	//
	// 	- **basic**: Basic Edition
	//
	// 	- **insurance**: Insurance Edition
	//
	// 	- **unlimited**: Unlimited Edition
	//
	// example:
	//
	// basic
	ChangingEdition *string `json:"ChangingEdition,omitempty" xml:"ChangingEdition,omitempty"`
	// The number of mitigation sessions with configuration changes.
	//
	// example:
	//
	// 100
	ChangingProtectNum *int32 `json:"ChangingProtectNum,omitempty" xml:"ChangingProtectNum,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PayByBandwidth**
	//
	// 	- **PayByTraffic**
	//
	// 	- **PayByBandwidth95**
	//
	// example:
	//
	// PayByTraffic
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of protected domain names.
	//
	// example:
	//
	// 20
	DomianNum *int32 `json:"DomianNum,omitempty" xml:"DomianNum,omitempty"`
	// The protection edition. Valid values:
	//
	// 	- **poc**: POC Edition
	//
	// 	- **basic**: Basic Edition
	//
	// 	- **insurance**: Insurance Edition
	//
	// 	- **unlimited**: Unlimited Edition
	//
	// example:
	//
	// poc
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The activation status of the service. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The service expiration time.
	//
	// example:
	//
	// 2023-09-26T16:00:00Z
	EndingTime *string `json:"EndingTime,omitempty" xml:"EndingTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// xxx-12345
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the service was enabled.
	//
	// example:
	//
	// 2021-09-26T16:00:00Z
	OpeningTime *string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	// The number of mitigation sessions.
	//
	// example:
	//
	// 100
	ProtectNum *int32 `json:"ProtectNum,omitempty" xml:"ProtectNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **WaitForExpire**
	//
	// 	- **expired**
	//
	// 	- **Released**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDdosServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDdosServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDdosServiceResponseBody) GetChangingAffectTime() *string {
	return s.ChangingAffectTime
}

func (s *DescribeDcdnDdosServiceResponseBody) GetChangingChargeType() *string {
	return s.ChangingChargeType
}

func (s *DescribeDcdnDdosServiceResponseBody) GetChangingDomianNum() *int32 {
	return s.ChangingDomianNum
}

func (s *DescribeDcdnDdosServiceResponseBody) GetChangingEdition() *string {
	return s.ChangingEdition
}

func (s *DescribeDcdnDdosServiceResponseBody) GetChangingProtectNum() *int32 {
	return s.ChangingProtectNum
}

func (s *DescribeDcdnDdosServiceResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDcdnDdosServiceResponseBody) GetDomianNum() *int32 {
	return s.DomianNum
}

func (s *DescribeDcdnDdosServiceResponseBody) GetEdition() *string {
	return s.Edition
}

func (s *DescribeDcdnDdosServiceResponseBody) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeDcdnDdosServiceResponseBody) GetEndingTime() *string {
	return s.EndingTime
}

func (s *DescribeDcdnDdosServiceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDcdnDdosServiceResponseBody) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *DescribeDcdnDdosServiceResponseBody) GetProtectNum() *int32 {
	return s.ProtectNum
}

func (s *DescribeDcdnDdosServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDdosServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnDdosServiceResponseBody) SetChangingAffectTime(v string) *DescribeDcdnDdosServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetChangingChargeType(v string) *DescribeDcdnDdosServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetChangingDomianNum(v int32) *DescribeDcdnDdosServiceResponseBody {
	s.ChangingDomianNum = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetChangingEdition(v string) *DescribeDcdnDdosServiceResponseBody {
	s.ChangingEdition = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetChangingProtectNum(v int32) *DescribeDcdnDdosServiceResponseBody {
	s.ChangingProtectNum = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetChargeType(v string) *DescribeDcdnDdosServiceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetDomianNum(v int32) *DescribeDcdnDdosServiceResponseBody {
	s.DomianNum = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetEdition(v string) *DescribeDcdnDdosServiceResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetEnabled(v string) *DescribeDcdnDdosServiceResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetEndingTime(v string) *DescribeDcdnDdosServiceResponseBody {
	s.EndingTime = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetInstanceId(v string) *DescribeDcdnDdosServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetOpeningTime(v string) *DescribeDcdnDdosServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetProtectNum(v int32) *DescribeDcdnDdosServiceResponseBody {
	s.ProtectNum = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetRequestId(v string) *DescribeDcdnDdosServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) SetStatus(v string) *DescribeDcdnDdosServiceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
