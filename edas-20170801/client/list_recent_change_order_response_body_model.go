// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentChangeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderList(v *ListRecentChangeOrderResponseBodyChangeOrderList) *ListRecentChangeOrderResponseBody
	GetChangeOrderList() *ListRecentChangeOrderResponseBodyChangeOrderList
	SetCode(v int32) *ListRecentChangeOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *ListRecentChangeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRecentChangeOrderResponseBody
	GetRequestId() *string
}

type ListRecentChangeOrderResponseBody struct {
	// The information about change processes.
	ChangeOrderList *ListRecentChangeOrderResponseBodyChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecentChangeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecentChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponseBody) GetChangeOrderList() *ListRecentChangeOrderResponseBodyChangeOrderList {
	return s.ChangeOrderList
}

func (s *ListRecentChangeOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListRecentChangeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRecentChangeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecentChangeOrderResponseBody) SetChangeOrderList(v *ListRecentChangeOrderResponseBodyChangeOrderList) *ListRecentChangeOrderResponseBody {
	s.ChangeOrderList = v
	return s
}

func (s *ListRecentChangeOrderResponseBody) SetCode(v int32) *ListRecentChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListRecentChangeOrderResponseBody) SetMessage(v string) *ListRecentChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ListRecentChangeOrderResponseBody) SetRequestId(v string) *ListRecentChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBody) Validate() error {
	if s.ChangeOrderList != nil {
		if err := s.ChangeOrderList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRecentChangeOrderResponseBodyChangeOrderList struct {
	ChangeOrder []*ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder `json:"ChangeOrder,omitempty" xml:"ChangeOrder,omitempty" type:"Repeated"`
}

func (s ListRecentChangeOrderResponseBodyChangeOrderList) String() string {
	return dara.Prettify(s)
}

func (s ListRecentChangeOrderResponseBodyChangeOrderList) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderList) GetChangeOrder() []*ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	return s.ChangeOrder
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderList) SetChangeOrder(v []*ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) *ListRecentChangeOrderResponseBodyChangeOrderList {
	s.ChangeOrder = v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderList) Validate() error {
	if s.ChangeOrder != nil {
		for _, item := range s.ChangeOrder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder struct {
	// The ID of the application.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of batches for the change. Valid values: 1 to 5.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// The way in which the next batch is triggered during a phased release. Valid values:
	//
	// 	- Automatic
	//
	// 	- Manual
	//
	// example:
	//
	// Automatic
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The description of the change process.
	//
	// example:
	//
	// Version: 2020-05-14 20:02:33 | Deployment Package: hsf-pandora-boot-provider-1.0.jar | Deploy to: all groups
	ChangeOrderDescription *string `json:"ChangeOrderDescription,omitempty" xml:"ChangeOrderDescription,omitempty"`
	// The unique ID of the change process.
	//
	// example:
	//
	// 1074f3e2-e974-4a0e-****-************
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The type of the change process.
	//
	// example:
	//
	// Application Scale Out
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The type of the change process.
	//
	// example:
	//
	// CoDeploy
	CoTypeCode *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	// The time when the change process was created.
	//
	// example:
	//
	// 2019-11-13 14:23:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user who created the change process.
	//
	// example:
	//
	// edas_test1@aliyun-test.com
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The time when the change process ended.
	//
	// example:
	//
	// 2019-11-13 14:24:02
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the application instance group on which the change was performed.
	//
	// example:
	//
	// 8123db90-880f-486f-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The source of the change. Valid values:
	//
	// 	- console: the Enterprise Distributed Application Service (EDAS) console
	//
	// 	- pop: the POP API or tool
	//
	// example:
	//
	// pop
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the change process. Valid values:
	//
	// 	- 0: ready to start execution
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: waiting for manual confirmation (You can see the state when you manually confirm the execution of the next batch of the change.)
	//
	// 	- 9: waiting for automatic execution
	//
	// 	- 10: failed due to a system error
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user who created the change process.
	//
	// example:
	//
	// 1432536****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) String() string {
	return dara.Prettify(s)
}

func (s ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetAppId() *string {
	return s.AppId
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetBatchCount() *int32 {
	return s.BatchCount
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetBatchType() *string {
	return s.BatchType
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetChangeOrderDescription() *string {
	return s.ChangeOrderDescription
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetCoType() *string {
	return s.CoType
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetCoTypeCode() *string {
	return s.CoTypeCode
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetGroupId() *string {
	return s.GroupId
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetSource() *string {
	return s.Source
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetStatus() *int32 {
	return s.Status
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GetUserId() *string {
	return s.UserId
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetAppId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.AppId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetBatchCount(v int32) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.BatchCount = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetBatchType(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.BatchType = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetChangeOrderDescription(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.ChangeOrderDescription = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetChangeOrderId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.ChangeOrderId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCoType(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CoType = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCoTypeCode(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CoTypeCode = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCreateTime(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CreateTime = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCreateUserId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CreateUserId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetFinishTime(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.FinishTime = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetGroupId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.GroupId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetSource(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.Source = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetStatus(v int32) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.Status = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetUserId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.UserId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) Validate() error {
	return dara.Validate(s)
}
