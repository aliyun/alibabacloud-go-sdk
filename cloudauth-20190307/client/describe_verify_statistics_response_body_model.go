// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifyStatisticsResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVerifyStatisticsResponseBodyResultObject) *DescribeVerifyStatisticsResponseBody
	GetResultObject() *DescribeVerifyStatisticsResponseBodyResultObject
}

type DescribeVerifyStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4E27D502-1287-526A-910C-881A3F023914
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics result.
	ResultObject *DescribeVerifyStatisticsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeVerifyStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyStatisticsResponseBody) GetResultObject() *DescribeVerifyStatisticsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVerifyStatisticsResponseBody) SetRequestId(v string) *DescribeVerifyStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBody) SetResultObject(v *DescribeVerifyStatisticsResponseBodyResultObject) *DescribeVerifyStatisticsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyStatisticsResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyStatisticsResponseBodyResultObject struct {
	// The number of client-side initializations.
	//
	// example:
	//
	// 73
	InitDevice *int64 `json:"InitDevice,omitempty" xml:"InitDevice,omitempty"`
	// The number of client-side initializations for identity deduplication.
	//
	// example:
	//
	// 9
	InitDeviceId *int64 `json:"InitDeviceId,omitempty" xml:"InitDeviceId,omitempty"`
	// The number of successful client-side initializations for identity deduplication.
	//
	// example:
	//
	// 9
	InitDeviceIdSuccess *int64 `json:"InitDeviceIdSuccess,omitempty" xml:"InitDeviceIdSuccess,omitempty"`
	// The number of successful client-side initialization calls.
	//
	// example:
	//
	// 73
	InitDeviceSuccess *int64 `json:"InitDeviceSuccess,omitempty" xml:"InitDeviceSuccess,omitempty"`
	// The number of server-side initializations.
	//
	// example:
	//
	// 73
	InitService *int64 `json:"InitService,omitempty" xml:"InitService,omitempty"`
	// The total number of server-side initialization requests for identity deduplication.
	//
	// example:
	//
	// 9
	InitServiceId *int64 `json:"InitServiceId,omitempty" xml:"InitServiceId,omitempty"`
	// The number of successful server-side initializations for identity deduplication.
	//
	// example:
	//
	// 9
	InitServiceIdSuccess *int64 `json:"InitServiceIdSuccess,omitempty" xml:"InitServiceIdSuccess,omitempty"`
	// The number of successful server-side initialization authentications.
	//
	// example:
	//
	// 134
	InitServiceSuccess *int64 `json:"InitServiceSuccess,omitempty" xml:"InitServiceSuccess,omitempty"`
	// The daily pass/conversion rates (PV).
	Items []*DescribeVerifyStatisticsResponseBodyResultObjectItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of client-side verifications.
	//
	// example:
	//
	// 15
	VerifyDevice *int64 `json:"VerifyDevice,omitempty" xml:"VerifyDevice,omitempty"`
	// The number of client-side authentication attempts for identity deduplication.
	//
	// example:
	//
	// 9
	VerifyDeviceId *int64 `json:"VerifyDeviceId,omitempty" xml:"VerifyDeviceId,omitempty"`
	// The number of successful client-side verifications for identity deduplication.
	//
	// example:
	//
	// 6
	VerifyDeviceIdSuccess *int64 `json:"VerifyDeviceIdSuccess,omitempty" xml:"VerifyDeviceIdSuccess,omitempty"`
	// The number of successful client-side authentications for identity deduplication.
	//
	// example:
	//
	// 3
	VerifyDeviceIdSuccessPassed *int64 `json:"VerifyDeviceIdSuccessPassed,omitempty" xml:"VerifyDeviceIdSuccessPassed,omitempty"`
	// The number of successful client-side authentications.
	//
	// example:
	//
	// 15
	VerifyDeviceSuccess *int64 `json:"VerifyDeviceSuccess,omitempty" xml:"VerifyDeviceSuccess,omitempty"`
	// The number of successful client-side authentications.
	//
	// example:
	//
	// 6
	VerifyDeviceSuccessPassed *int64 `json:"VerifyDeviceSuccessPassed,omitempty" xml:"VerifyDeviceSuccessPassed,omitempty"`
}

func (s DescribeVerifyStatisticsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyStatisticsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitDevice() *int64 {
	return s.InitDevice
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitDeviceId() *int64 {
	return s.InitDeviceId
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitDeviceIdSuccess() *int64 {
	return s.InitDeviceIdSuccess
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitDeviceSuccess() *int64 {
	return s.InitDeviceSuccess
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitService() *int64 {
	return s.InitService
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitServiceId() *int64 {
	return s.InitServiceId
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitServiceIdSuccess() *int64 {
	return s.InitServiceIdSuccess
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetInitServiceSuccess() *int64 {
	return s.InitServiceSuccess
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetItems() []*DescribeVerifyStatisticsResponseBodyResultObjectItems {
	return s.Items
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetVerifyDevice() *int64 {
	return s.VerifyDevice
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetVerifyDeviceId() *int64 {
	return s.VerifyDeviceId
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetVerifyDeviceIdSuccess() *int64 {
	return s.VerifyDeviceIdSuccess
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetVerifyDeviceIdSuccessPassed() *int64 {
	return s.VerifyDeviceIdSuccessPassed
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetVerifyDeviceSuccess() *int64 {
	return s.VerifyDeviceSuccess
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) GetVerifyDeviceSuccessPassed() *int64 {
	return s.VerifyDeviceSuccessPassed
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitDevice(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitDevice = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitDeviceId(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitDeviceId = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitDeviceIdSuccess(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitDeviceIdSuccess = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitDeviceSuccess(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitDeviceSuccess = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitService(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitService = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitServiceId(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitServiceId = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitServiceIdSuccess(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitServiceIdSuccess = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetInitServiceSuccess(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.InitServiceSuccess = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetItems(v []*DescribeVerifyStatisticsResponseBodyResultObjectItems) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.Items = v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetVerifyDevice(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.VerifyDevice = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetVerifyDeviceId(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.VerifyDeviceId = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetVerifyDeviceIdSuccess(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.VerifyDeviceIdSuccess = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetVerifyDeviceIdSuccessPassed(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.VerifyDeviceIdSuccessPassed = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetVerifyDeviceSuccess(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.VerifyDeviceSuccess = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) SetVerifyDeviceSuccessPassed(v int64) *DescribeVerifyStatisticsResponseBodyResultObject {
	s.VerifyDeviceSuccessPassed = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObject) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVerifyStatisticsResponseBodyResultObjectItems struct {
	// The date.
	//
	// example:
	//
	// 2025-10-11
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The client-side initialization pass rate.
	//
	// example:
	//
	// 60
	InitDevicePassRate *string `json:"InitDevicePassRate,omitempty" xml:"InitDevicePassRate,omitempty"`
	// The number of server-side initializations.
	//
	// example:
	//
	// 15
	InitService *int64 `json:"InitService,omitempty" xml:"InitService,omitempty"`
	// The server-side initialization conversion rate.
	//
	// example:
	//
	// 26.67
	InitServiceConversionRate *string `json:"InitServiceConversionRate,omitempty" xml:"InitServiceConversionRate,omitempty"`
	// The server-side initialization pass rate.
	//
	// example:
	//
	// 20
	InitServicePassRate *string `json:"InitServicePassRate,omitempty" xml:"InitServicePassRate,omitempty"`
	// The pass rate.
	//
	// example:
	//
	// 75
	PassRate *string `json:"PassRate,omitempty" xml:"PassRate,omitempty"`
}

func (s DescribeVerifyStatisticsResponseBodyResultObjectItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyStatisticsResponseBodyResultObjectItems) GoString() string {
	return s.String()
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) GetDate() *string {
	return s.Date
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) GetInitDevicePassRate() *string {
	return s.InitDevicePassRate
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) GetInitService() *int64 {
	return s.InitService
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) GetInitServiceConversionRate() *string {
	return s.InitServiceConversionRate
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) GetInitServicePassRate() *string {
	return s.InitServicePassRate
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) GetPassRate() *string {
	return s.PassRate
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) SetDate(v string) *DescribeVerifyStatisticsResponseBodyResultObjectItems {
	s.Date = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) SetInitDevicePassRate(v string) *DescribeVerifyStatisticsResponseBodyResultObjectItems {
	s.InitDevicePassRate = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) SetInitService(v int64) *DescribeVerifyStatisticsResponseBodyResultObjectItems {
	s.InitService = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) SetInitServiceConversionRate(v string) *DescribeVerifyStatisticsResponseBodyResultObjectItems {
	s.InitServiceConversionRate = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) SetInitServicePassRate(v string) *DescribeVerifyStatisticsResponseBodyResultObjectItems {
	s.InitServicePassRate = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) SetPassRate(v string) *DescribeVerifyStatisticsResponseBodyResultObjectItems {
	s.PassRate = &v
	return s
}

func (s *DescribeVerifyStatisticsResponseBodyResultObjectItems) Validate() error {
	return dara.Validate(s)
}
