// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomizeFlowStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCustomizeFlowStrategyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *QueryCustomizeFlowStrategyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryCustomizeFlowStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCustomizeFlowStrategyResponseBody
	GetRequestId() *string
	SetResultObject(v []*QueryCustomizeFlowStrategyResponseBodyResultObject) *QueryCustomizeFlowStrategyResponseBody
	GetResultObject() []*QueryCustomizeFlowStrategyResponseBodyResultObject
	SetSuccess(v bool) *QueryCustomizeFlowStrategyResponseBody
	GetSuccess() *bool
}

type QueryCustomizeFlowStrategyResponseBody struct {
	// The return code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The processing result.
	ResultObject []*QueryCustomizeFlowStrategyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// Indicates whether the response is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomizeFlowStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomizeFlowStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomizeFlowStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCustomizeFlowStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryCustomizeFlowStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCustomizeFlowStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCustomizeFlowStrategyResponseBody) GetResultObject() []*QueryCustomizeFlowStrategyResponseBodyResultObject {
	return s.ResultObject
}

func (s *QueryCustomizeFlowStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCustomizeFlowStrategyResponseBody) SetCode(v string) *QueryCustomizeFlowStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBody) SetHttpStatusCode(v int32) *QueryCustomizeFlowStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBody) SetMessage(v string) *QueryCustomizeFlowStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBody) SetRequestId(v string) *QueryCustomizeFlowStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBody) SetResultObject(v []*QueryCustomizeFlowStrategyResponseBodyResultObject) *QueryCustomizeFlowStrategyResponseBody {
	s.ResultObject = v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBody) SetSuccess(v bool) *QueryCustomizeFlowStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBody) Validate() error {
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

type QueryCustomizeFlowStrategyResponseBodyResultObject struct {
	// AccumulateKey
	//
	// example:
	//
	// -
	AccumulateKey *string `json:"AccumulateKey,omitempty" xml:"AccumulateKey,omitempty"`
	// The rate limiting statistical window. Unit: **minutes**.
	//
	// example:
	//
	// 60
	AccumulateWindow *string `json:"AccumulateWindow,omitempty" xml:"AccumulateWindow,omitempty"`
	// The API name, which is the same as **ProductCode**.
	//
	// example:
	//
	// ID_PRO
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The flow type. Valid values:
	//
	// - **ACCUMULATE**: repeated occurrence of an ID card number.
	//
	// - **PASSED_RATE**: pass rate less than the threshold.
	//
	// - **SUB_CODE_205**: authentication failed and the proportion of liveness attack 205 is greater than the threshold.
	//
	// - **SUB_CODE_206**: authentication failed and the proportion of liveness attack 206 is greater than the threshold.
	//
	// example:
	//
	// ACCUMULATE
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 234822
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The operation.
	//
	// example:
	//
	// match
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The status. Valid values:
	//
	// - **disabled**: Disabled.
	//
	// - **normal**: Enabled.
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The threshold.
	//
	// example:
	//
	// 10
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 126005125163xxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryCustomizeFlowStrategyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomizeFlowStrategyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetAccumulateKey() *string {
	return s.AccumulateKey
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetAccumulateWindow() *string {
	return s.AccumulateWindow
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetApiName() *string {
	return s.ApiName
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetFlowType() *string {
	return s.FlowType
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetId() *string {
	return s.Id
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetOperation() *string {
	return s.Operation
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetThreshold() *string {
	return s.Threshold
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) GetUserId() *string {
	return s.UserId
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetAccumulateKey(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.AccumulateKey = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetAccumulateWindow(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.AccumulateWindow = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetApiName(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.ApiName = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetFlowType(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.FlowType = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetId(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetOperation(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.Operation = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetStatus(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetThreshold(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.Threshold = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) SetUserId(v string) *QueryCustomizeFlowStrategyResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *QueryCustomizeFlowStrategyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
