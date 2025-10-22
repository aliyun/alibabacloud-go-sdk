// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProduceOperateLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUserProduceOperateLogsResponseBodyData) *ListUserProduceOperateLogsResponseBody
	GetData() []*ListUserProduceOperateLogsResponseBodyData
	SetPageNum(v int32) *ListUserProduceOperateLogsResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserProduceOperateLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserProduceOperateLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserProduceOperateLogsResponseBody
	GetSuccess() *bool
	SetTotalItemNum(v int32) *ListUserProduceOperateLogsResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListUserProduceOperateLogsResponseBody
	GetTotalPageNum() *int32
}

type ListUserProduceOperateLogsResponseBody struct {
	Data []*ListUserProduceOperateLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0DCBE2FF-2DFC-56DC-9A15-BDF27B7FFB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 6
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 23
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserProduceOperateLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserProduceOperateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsResponseBody) GetData() []*ListUserProduceOperateLogsResponseBodyData {
	return s.Data
}

func (s *ListUserProduceOperateLogsResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserProduceOperateLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserProduceOperateLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserProduceOperateLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserProduceOperateLogsResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListUserProduceOperateLogsResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListUserProduceOperateLogsResponseBody) SetData(v []*ListUserProduceOperateLogsResponseBodyData) *ListUserProduceOperateLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetPageNum(v int32) *ListUserProduceOperateLogsResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetPageSize(v int32) *ListUserProduceOperateLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetRequestId(v string) *ListUserProduceOperateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetSuccess(v bool) *ListUserProduceOperateLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetTotalItemNum(v int32) *ListUserProduceOperateLogsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetTotalPageNum(v int32) *ListUserProduceOperateLogsResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) Validate() error {
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

type ListUserProduceOperateLogsResponseBodyData struct {
	// example:
	//
	// P20210928095324000002
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 10
	BizStatus *int32 `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// example:
	//
	// esp.wangwen
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OperateName *string `json:"OperateName,omitempty" xml:"OperateName,omitempty"`
	// example:
	//
	// 1695324000002
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// example:
	//
	// user
	OperateUserType *string `json:"OperateUserType,omitempty" xml:"OperateUserType,omitempty"`
	// example:
	//
	// 35
	ToBizStatus *int32 `json:"ToBizStatus,omitempty" xml:"ToBizStatus,omitempty"`
}

func (s ListUserProduceOperateLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserProduceOperateLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetBizStatus() *int32 {
	return s.BizStatus
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetOperateName() *string {
	return s.OperateName
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetOperateTime() *int64 {
	return s.OperateTime
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetOperateUserType() *string {
	return s.OperateUserType
}

func (s *ListUserProduceOperateLogsResponseBodyData) GetToBizStatus() *int32 {
	return s.ToBizStatus
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetBizId(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetBizStatus(v int32) *ListUserProduceOperateLogsResponseBodyData {
	s.BizStatus = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetBizType(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetNote(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.Note = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetOperateName(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.OperateName = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetOperateTime(v int64) *ListUserProduceOperateLogsResponseBodyData {
	s.OperateTime = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetOperateUserType(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.OperateUserType = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetToBizStatus(v int32) *ListUserProduceOperateLogsResponseBodyData {
	s.ToBizStatus = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
