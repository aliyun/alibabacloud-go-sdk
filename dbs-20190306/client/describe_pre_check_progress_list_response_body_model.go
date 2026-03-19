// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckProgressListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribePreCheckProgressListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribePreCheckProgressListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribePreCheckProgressListResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribePreCheckProgressListResponseBodyItems) *DescribePreCheckProgressListResponseBody
	GetItems() *DescribePreCheckProgressListResponseBodyItems
	SetProgress(v int32) *DescribePreCheckProgressListResponseBody
	GetProgress() *int32
	SetRequestId(v string) *DescribePreCheckProgressListResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribePreCheckProgressListResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DescribePreCheckProgressListResponseBody
	GetSuccess() *bool
}

type DescribePreCheckProgressListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribePreCheckProgressListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The precheck progress. The value ranges from 0 to 100.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C4A45FE1-A903-470D-B113-F12A4DF942AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The precheck status. Valid values:
	//
	// - **running**: The precheck is in progress.
	//
	// - **failed**: The precheck failed.
	//
	// - **finish**: The precheck is complete.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePreCheckProgressListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckProgressListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribePreCheckProgressListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribePreCheckProgressListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribePreCheckProgressListResponseBody) GetItems() *DescribePreCheckProgressListResponseBodyItems {
	return s.Items
}

func (s *DescribePreCheckProgressListResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribePreCheckProgressListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePreCheckProgressListResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribePreCheckProgressListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePreCheckProgressListResponseBody) SetErrCode(v string) *DescribePreCheckProgressListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetErrMessage(v string) *DescribePreCheckProgressListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetHttpStatusCode(v int32) *DescribePreCheckProgressListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetItems(v *DescribePreCheckProgressListResponseBodyItems) *DescribePreCheckProgressListResponseBody {
	s.Items = v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetProgress(v int32) *DescribePreCheckProgressListResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetRequestId(v string) *DescribePreCheckProgressListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetStatus(v string) *DescribePreCheckProgressListResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetSuccess(v bool) *DescribePreCheckProgressListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePreCheckProgressListResponseBodyItems struct {
	PreCheckProgressDetail []*DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail `json:"PreCheckProgressDetail,omitempty" xml:"PreCheckProgressDetail,omitempty" type:"Repeated"`
}

func (s DescribePreCheckProgressListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckProgressListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponseBodyItems) GetPreCheckProgressDetail() []*DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	return s.PreCheckProgressDetail
}

func (s *DescribePreCheckProgressListResponseBodyItems) SetPreCheckProgressDetail(v []*DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) *DescribePreCheckProgressListResponseBodyItems {
	s.PreCheckProgressDetail = v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItems) Validate() error {
	if s.PreCheckProgressDetail != nil {
		for _, item := range s.PreCheckProgressDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail struct {
	BootTime   *int64  `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	ErrMsg     *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	FinishTime *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Item       *string `json:"Item,omitempty" xml:"Item,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Names      *string `json:"Names,omitempty" xml:"Names,omitempty"`
	OrderNum   *string `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetBootTime() *int64 {
	return s.BootTime
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetItem() *string {
	return s.Item
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetNames() *string {
	return s.Names
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetOrderNum() *string {
	return s.OrderNum
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GetState() *string {
	return s.State
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetBootTime(v int64) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetErrMsg(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetFinishTime(v int64) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetItem(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.Item = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetJobId(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetNames(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.Names = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetOrderNum(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetState(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.State = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) Validate() error {
	return dara.Validate(s)
}
