// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotAttackStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHoneypotAttackStatisticsResponseBody
	GetCode() *string
	SetData(v *GetHoneypotAttackStatisticsResponseBodyData) *GetHoneypotAttackStatisticsResponseBody
	GetData() *GetHoneypotAttackStatisticsResponseBodyData
	SetHttpStatusCode(v int32) *GetHoneypotAttackStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHoneypotAttackStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHoneypotAttackStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoneypotAttackStatisticsResponseBody
	GetSuccess() *bool
}

type GetHoneypotAttackStatisticsResponseBody struct {
	// The status code that is returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The statistics.
	Data *GetHoneypotAttackStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FACC60A-3FE4-5F49-9184-50730C8B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHoneypotAttackStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotAttackStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoneypotAttackStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHoneypotAttackStatisticsResponseBody) GetData() *GetHoneypotAttackStatisticsResponseBodyData {
	return s.Data
}

func (s *GetHoneypotAttackStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHoneypotAttackStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHoneypotAttackStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoneypotAttackStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoneypotAttackStatisticsResponseBody) SetCode(v string) *GetHoneypotAttackStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBody) SetData(v *GetHoneypotAttackStatisticsResponseBodyData) *GetHoneypotAttackStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBody) SetHttpStatusCode(v int32) *GetHoneypotAttackStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBody) SetMessage(v string) *GetHoneypotAttackStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBody) SetRequestId(v string) *GetHoneypotAttackStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBody) SetSuccess(v bool) *GetHoneypotAttackStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHoneypotAttackStatisticsResponseBodyData struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The statistics details.
	HoneypotAttackStatistics []*GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics `json:"HoneypotAttackStatistics,omitempty" xml:"HoneypotAttackStatistics,omitempty" type:"Repeated"`
	// The type of the attack source statistics. Valid values:
	//
	// 	- **TOP_ATTACKED_AGENT**: the top five probes that are attacked the most frequently.
	//
	// 	- **TOP_ATTACKED_IP**: the top five IP addresses that are attacked the most frequently.
	//
	// 	- **ATTACK_EVENT_TYPE**: the type of the intrusion event.
	//
	// 	- **ATTACK_HONEYPOT_TYPE**: the type of the attacked honeypot.
	//
	// example:
	//
	// TOP_ATTACKED_IP
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
}

func (s GetHoneypotAttackStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotAttackStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHoneypotAttackStatisticsResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetHoneypotAttackStatisticsResponseBodyData) GetHoneypotAttackStatistics() []*GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics {
	return s.HoneypotAttackStatistics
}

func (s *GetHoneypotAttackStatisticsResponseBodyData) GetStatisticsType() *string {
	return s.StatisticsType
}

func (s *GetHoneypotAttackStatisticsResponseBodyData) SetCount(v int32) *GetHoneypotAttackStatisticsResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBodyData) SetHoneypotAttackStatistics(v []*GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) *GetHoneypotAttackStatisticsResponseBodyData {
	s.HoneypotAttackStatistics = v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBodyData) SetStatisticsType(v string) *GetHoneypotAttackStatisticsResponseBodyData {
	s.StatisticsType = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics struct {
	// The number of times the value is counted.
	//
	// example:
	//
	// 10
	StatisticsCount *int32 `json:"StatisticsCount,omitempty" xml:"StatisticsCount,omitempty"`
	// The statistical value.
	//
	// example:
	//
	// 112.168.1.**
	StatisticsValue *string `json:"StatisticsValue,omitempty" xml:"StatisticsValue,omitempty"`
}

func (s GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) GoString() string {
	return s.String()
}

func (s *GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) GetStatisticsCount() *int32 {
	return s.StatisticsCount
}

func (s *GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) GetStatisticsValue() *string {
	return s.StatisticsValue
}

func (s *GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) SetStatisticsCount(v int32) *GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics {
	s.StatisticsCount = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) SetStatisticsValue(v string) *GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics {
	s.StatisticsValue = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponseBodyDataHoneypotAttackStatistics) Validate() error {
	return dara.Validate(s)
}
