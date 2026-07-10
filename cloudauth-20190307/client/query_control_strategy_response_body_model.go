// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryControlStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryControlStrategyResponseBody
	GetCode() *string
	SetMessage(v string) *QueryControlStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryControlStrategyResponseBody
	GetRequestId() *string
	SetResultObject(v []*QueryControlStrategyResponseBodyResultObject) *QueryControlStrategyResponseBody
	GetResultObject() []*QueryControlStrategyResponseBodyResultObject
}

type QueryControlStrategyResponseBody struct {
	// The response code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B506328A-D84B-4750-82C7-6A207C585CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The processing result.
	ResultObject []*QueryControlStrategyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
}

func (s QueryControlStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryControlStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryControlStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryControlStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryControlStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryControlStrategyResponseBody) GetResultObject() []*QueryControlStrategyResponseBodyResultObject {
	return s.ResultObject
}

func (s *QueryControlStrategyResponseBody) SetCode(v string) *QueryControlStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryControlStrategyResponseBody) SetMessage(v string) *QueryControlStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryControlStrategyResponseBody) SetRequestId(v string) *QueryControlStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryControlStrategyResponseBody) SetResultObject(v []*QueryControlStrategyResponseBodyResultObject) *QueryControlStrategyResponseBody {
	s.ResultObject = v
	return s
}

func (s *QueryControlStrategyResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryControlStrategyResponseBodyResultObject struct {
	// The API name, which is the same as the ProductCode of the verification operation.
	//
	// example:
	//
	// ID_PRO
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The rule configuration type. Valid values:
	//
	// - **QPS**: QPS greater than
	//
	// - **SUCCESS_RATE_5_MIN**: success rate in the last 5 minutes less than
	//
	// - **RESP_TIME_5_MIN**: average response time in the last 5 minutes greater than
	//
	// - **AMOUNT_RISE**: call volume increase compared to the previous period greater than
	//
	// - **AMOUNT_FALL**: call volume decrease compared to the previous period less than
	//
	// - **PASSED_RATE_1_HOUR**: verification consistency rate in the last hour less than
	//
	// - **PARAM_ERROR_RATE_1_HOUR**: parameter error rate in the last hour greater than.
	//
	// example:
	//
	// SUCCESS_RATE_5_MIN
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 234822
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status. Valid values:
	//
	// - **disabled**: Disabled.
	//
	// - **normal**: Enabled.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The alert threshold for configure alerting of the rule.
	//
	// example:
	//
	// 0.9
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 126005125163xxxx
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryControlStrategyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s QueryControlStrategyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *QueryControlStrategyResponseBodyResultObject) GetApiName() *string {
	return s.ApiName
}

func (s *QueryControlStrategyResponseBodyResultObject) GetBizType() *string {
	return s.BizType
}

func (s *QueryControlStrategyResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *QueryControlStrategyResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *QueryControlStrategyResponseBodyResultObject) GetThreshold() *float64 {
	return s.Threshold
}

func (s *QueryControlStrategyResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryControlStrategyResponseBodyResultObject) SetApiName(v string) *QueryControlStrategyResponseBodyResultObject {
	s.ApiName = &v
	return s
}

func (s *QueryControlStrategyResponseBodyResultObject) SetBizType(v string) *QueryControlStrategyResponseBodyResultObject {
	s.BizType = &v
	return s
}

func (s *QueryControlStrategyResponseBodyResultObject) SetId(v int64) *QueryControlStrategyResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *QueryControlStrategyResponseBodyResultObject) SetStatus(v string) *QueryControlStrategyResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *QueryControlStrategyResponseBodyResultObject) SetThreshold(v float64) *QueryControlStrategyResponseBodyResultObject {
	s.Threshold = &v
	return s
}

func (s *QueryControlStrategyResponseBodyResultObject) SetUserId(v int64) *QueryControlStrategyResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *QueryControlStrategyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
