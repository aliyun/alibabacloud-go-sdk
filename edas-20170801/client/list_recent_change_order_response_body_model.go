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
	AppId                  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BatchCount             *int32  `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	BatchType              *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	ChangeOrderDescription *string `json:"ChangeOrderDescription,omitempty" xml:"ChangeOrderDescription,omitempty"`
	ChangeOrderId          *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	CoType                 *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	CoTypeCode             *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId           *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	FinishTime             *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	GroupId                *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Source                 *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status                 *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
