// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKafkaClientIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetKafkaClientIpResponseBody
	GetCode() *int64
	SetData(v *GetKafkaClientIpResponseBodyData) *GetKafkaClientIpResponseBody
	GetData() *GetKafkaClientIpResponseBodyData
	SetMessage(v string) *GetKafkaClientIpResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKafkaClientIpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKafkaClientIpResponseBody
	GetSuccess() *bool
}

type GetKafkaClientIpResponseBody struct {
	// The returned status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetKafkaClientIpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E57A8862-DF68-4055-8E55-B80CB4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKafkaClientIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetKafkaClientIpResponseBody) GetData() *GetKafkaClientIpResponseBodyData {
	return s.Data
}

func (s *GetKafkaClientIpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKafkaClientIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKafkaClientIpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKafkaClientIpResponseBody) SetCode(v int64) *GetKafkaClientIpResponseBody {
	s.Code = &v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetData(v *GetKafkaClientIpResponseBodyData) *GetKafkaClientIpResponseBody {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetMessage(v string) *GetKafkaClientIpResponseBody {
	s.Message = &v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetRequestId(v string) *GetKafkaClientIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetSuccess(v bool) *GetKafkaClientIpResponseBody {
	s.Success = &v
	return s
}

func (s *GetKafkaClientIpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKafkaClientIpResponseBodyData struct {
	// The value true indicates that the broker is not of the latest minor version.
	//
	// >  If the broker is not of the latest minor version, the sampled logs may not be accurate. This may cause inaccurate IP information. Therefore, we recommend that you update your broker to the latest version at the earliest opportunity.
	//
	// example:
	//
	// true
	Alert *bool `json:"Alert,omitempty" xml:"Alert,omitempty"`
	// The data returned.
	Data *GetKafkaClientIpResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The end of the date range within which data is queried.
	//
	// example:
	//
	// 1716343502000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The time range within which the client IP addresses are queried.
	//
	// >  The valid value is 1 hour. If the beginning of the time range to query and the end of the time range to query exceeds 1 hour, only data within 1 hour is returned.
	//
	// example:
	//
	// 1
	SearchTimeRange *int32 `json:"SearchTimeRange,omitempty" xml:"SearchTimeRange,omitempty"`
	// The beginning of the date range within which data is queried.
	//
	// example:
	//
	// 1716343501000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The date range within which the client IP addresses are queried.
	//
	// >  The valid value is 7 days. If the beginning of the date range to query and the end of the date range to query exceeds 7 days, only data within 7 days is returned.
	//
	// example:
	//
	// 7
	TimeLimitDay *int32 `json:"TimeLimitDay,omitempty" xml:"TimeLimitDay,omitempty"`
}

func (s GetKafkaClientIpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyData) GetAlert() *bool {
	return s.Alert
}

func (s *GetKafkaClientIpResponseBodyData) GetData() *GetKafkaClientIpResponseBodyDataData {
	return s.Data
}

func (s *GetKafkaClientIpResponseBodyData) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetKafkaClientIpResponseBodyData) GetSearchTimeRange() *int32 {
	return s.SearchTimeRange
}

func (s *GetKafkaClientIpResponseBodyData) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetKafkaClientIpResponseBodyData) GetTimeLimitDay() *int32 {
	return s.TimeLimitDay
}

func (s *GetKafkaClientIpResponseBodyData) SetAlert(v bool) *GetKafkaClientIpResponseBodyData {
	s.Alert = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetData(v *GetKafkaClientIpResponseBodyDataData) *GetKafkaClientIpResponseBodyData {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetEndDate(v int64) *GetKafkaClientIpResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetSearchTimeRange(v int32) *GetKafkaClientIpResponseBodyData {
	s.SearchTimeRange = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetStartDate(v int64) *GetKafkaClientIpResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetTimeLimitDay(v int32) *GetKafkaClientIpResponseBodyData {
	s.TimeLimitDay = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKafkaClientIpResponseBodyDataData struct {
	Data []*GetKafkaClientIpResponseBodyDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetKafkaClientIpResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataData) GetData() []*GetKafkaClientIpResponseBodyDataDataData {
	return s.Data
}

func (s *GetKafkaClientIpResponseBodyDataData) SetData(v []*GetKafkaClientIpResponseBodyDataDataData) *GetKafkaClientIpResponseBodyDataData {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKafkaClientIpResponseBodyDataDataData struct {
	// The response parameters.
	Data *GetKafkaClientIpResponseBodyDataDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request name.
	//
	// >  The value of this parameter indicates the type of request that the client sends to the broker within the specified period of time.
	//
	// example:
	//
	// OFFSET_COMMIT
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetKafkaClientIpResponseBodyDataDataData) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataDataData) GetData() *GetKafkaClientIpResponseBodyDataDataDataData {
	return s.Data
}

func (s *GetKafkaClientIpResponseBodyDataDataData) GetName() *string {
	return s.Name
}

func (s *GetKafkaClientIpResponseBodyDataDataData) SetData(v *GetKafkaClientIpResponseBodyDataDataDataData) *GetKafkaClientIpResponseBodyDataDataData {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataDataData) SetName(v string) *GetKafkaClientIpResponseBodyDataDataData {
	s.Name = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataDataData) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKafkaClientIpResponseBodyDataDataDataData struct {
	Data []*GetKafkaClientIpResponseBodyDataDataDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetKafkaClientIpResponseBodyDataDataDataData) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataDataDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataDataDataData) GetData() []*GetKafkaClientIpResponseBodyDataDataDataDataData {
	return s.Data
}

func (s *GetKafkaClientIpResponseBodyDataDataDataData) SetData(v []*GetKafkaClientIpResponseBodyDataDataDataDataData) *GetKafkaClientIpResponseBodyDataDataDataData {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataDataDataData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKafkaClientIpResponseBodyDataDataDataDataData struct {
	// The IP address of the client.
	//
	// example:
	//
	// 58.210.117.154
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The statistics.
	//
	// >  The value of this parameter indicates the number of connections on different ports of the IP address within the specified period of time.
	//
	// example:
	//
	// 3
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
}

func (s GetKafkaClientIpResponseBodyDataDataDataDataData) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataDataDataDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataDataDataDataData) GetIp() *string {
	return s.Ip
}

func (s *GetKafkaClientIpResponseBodyDataDataDataDataData) GetNum() *int64 {
	return s.Num
}

func (s *GetKafkaClientIpResponseBodyDataDataDataDataData) SetIp(v string) *GetKafkaClientIpResponseBodyDataDataDataDataData {
	s.Ip = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataDataDataDataData) SetNum(v int64) *GetKafkaClientIpResponseBodyDataDataDataDataData {
	s.Num = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataDataDataDataData) Validate() error {
	return dara.Validate(s)
}
