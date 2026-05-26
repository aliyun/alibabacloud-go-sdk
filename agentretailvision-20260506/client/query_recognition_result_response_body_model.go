// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognitionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRecognitionResultResponseBody
	GetCode() *string
	SetData(v *QueryRecognitionResultResponseBodyData) *QueryRecognitionResultResponseBody
	GetData() *QueryRecognitionResultResponseBodyData
	SetMessage(v string) *QueryRecognitionResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRecognitionResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRecognitionResultResponseBody
	GetSuccess() *bool
}

type QueryRecognitionResultResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryRecognitionResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecognitionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognitionResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecognitionResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRecognitionResultResponseBody) GetData() *QueryRecognitionResultResponseBodyData {
	return s.Data
}

func (s *QueryRecognitionResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRecognitionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecognitionResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRecognitionResultResponseBody) SetCode(v string) *QueryRecognitionResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecognitionResultResponseBody) SetData(v *QueryRecognitionResultResponseBodyData) *QueryRecognitionResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecognitionResultResponseBody) SetMessage(v string) *QueryRecognitionResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRecognitionResultResponseBody) SetRequestId(v string) *QueryRecognitionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecognitionResultResponseBody) SetSuccess(v bool) *QueryRecognitionResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRecognitionResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRecognitionResultResponseBodyData struct {
	// example:
	//
	// ORDER_001
	OrderUniqueId *string                                       `json:"OrderUniqueId,omitempty" xml:"OrderUniqueId,omitempty"`
	Result        *QueryRecognitionResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// TASK_001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// COMPLETED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s QueryRecognitionResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognitionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecognitionResultResponseBodyData) GetOrderUniqueId() *string {
	return s.OrderUniqueId
}

func (s *QueryRecognitionResultResponseBodyData) GetResult() *QueryRecognitionResultResponseBodyDataResult {
	return s.Result
}

func (s *QueryRecognitionResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryRecognitionResultResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryRecognitionResultResponseBodyData) SetOrderUniqueId(v string) *QueryRecognitionResultResponseBodyData {
	s.OrderUniqueId = &v
	return s
}

func (s *QueryRecognitionResultResponseBodyData) SetResult(v *QueryRecognitionResultResponseBodyDataResult) *QueryRecognitionResultResponseBodyData {
	s.Result = v
	return s
}

func (s *QueryRecognitionResultResponseBodyData) SetTaskId(v string) *QueryRecognitionResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryRecognitionResultResponseBodyData) SetTaskStatus(v string) *QueryRecognitionResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *QueryRecognitionResultResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRecognitionResultResponseBodyDataResult struct {
	CheckoutInfo *QueryRecognitionResultResponseBodyDataResultCheckoutInfo `json:"CheckoutInfo,omitempty" xml:"CheckoutInfo,omitempty" type:"Struct"`
	Items        []*QueryRecognitionResultResponseBodyDataResultItems      `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s QueryRecognitionResultResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognitionResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *QueryRecognitionResultResponseBodyDataResult) GetCheckoutInfo() *QueryRecognitionResultResponseBodyDataResultCheckoutInfo {
	return s.CheckoutInfo
}

func (s *QueryRecognitionResultResponseBodyDataResult) GetItems() []*QueryRecognitionResultResponseBodyDataResultItems {
	return s.Items
}

func (s *QueryRecognitionResultResponseBodyDataResult) SetCheckoutInfo(v *QueryRecognitionResultResponseBodyDataResultCheckoutInfo) *QueryRecognitionResultResponseBodyDataResult {
	s.CheckoutInfo = v
	return s
}

func (s *QueryRecognitionResultResponseBodyDataResult) SetItems(v []*QueryRecognitionResultResponseBodyDataResultItems) *QueryRecognitionResultResponseBodyDataResult {
	s.Items = v
	return s
}

func (s *QueryRecognitionResultResponseBodyDataResult) Validate() error {
	if s.CheckoutInfo != nil {
		if err := s.CheckoutInfo.Validate(); err != nil {
			return err
		}
	}
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

type QueryRecognitionResultResponseBodyDataResultCheckoutInfo struct {
	// example:
	//
	// TRUSTED
	CheckoutStatus *string `json:"CheckoutStatus,omitempty" xml:"CheckoutStatus,omitempty"`
}

func (s QueryRecognitionResultResponseBodyDataResultCheckoutInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognitionResultResponseBodyDataResultCheckoutInfo) GoString() string {
	return s.String()
}

func (s *QueryRecognitionResultResponseBodyDataResultCheckoutInfo) GetCheckoutStatus() *string {
	return s.CheckoutStatus
}

func (s *QueryRecognitionResultResponseBodyDataResultCheckoutInfo) SetCheckoutStatus(v string) *QueryRecognitionResultResponseBodyDataResultCheckoutInfo {
	s.CheckoutStatus = &v
	return s
}

func (s *QueryRecognitionResultResponseBodyDataResultCheckoutInfo) Validate() error {
	return dara.Validate(s)
}

type QueryRecognitionResultResponseBodyDataResultItems struct {
	// example:
	//
	// 690234524880781
	ItemUniqueId *string `json:"ItemUniqueId,omitempty" xml:"ItemUniqueId,omitempty"`
	// example:
	//
	// 535c3daaee3b4b5382db4913413419bc2d
	PlatformItemId *string `json:"PlatformItemId,omitempty" xml:"PlatformItemId,omitempty"`
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s QueryRecognitionResultResponseBodyDataResultItems) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognitionResultResponseBodyDataResultItems) GoString() string {
	return s.String()
}

func (s *QueryRecognitionResultResponseBodyDataResultItems) GetItemUniqueId() *string {
	return s.ItemUniqueId
}

func (s *QueryRecognitionResultResponseBodyDataResultItems) GetPlatformItemId() *string {
	return s.PlatformItemId
}

func (s *QueryRecognitionResultResponseBodyDataResultItems) GetQuantity() *int32 {
	return s.Quantity
}

func (s *QueryRecognitionResultResponseBodyDataResultItems) SetItemUniqueId(v string) *QueryRecognitionResultResponseBodyDataResultItems {
	s.ItemUniqueId = &v
	return s
}

func (s *QueryRecognitionResultResponseBodyDataResultItems) SetPlatformItemId(v string) *QueryRecognitionResultResponseBodyDataResultItems {
	s.PlatformItemId = &v
	return s
}

func (s *QueryRecognitionResultResponseBodyDataResultItems) SetQuantity(v int32) *QueryRecognitionResultResponseBodyDataResultItems {
	s.Quantity = &v
	return s
}

func (s *QueryRecognitionResultResponseBodyDataResultItems) Validate() error {
	return dara.Validate(s)
}
