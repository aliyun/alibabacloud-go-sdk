// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDrawRecordByPkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserDrawRecordByPkResponseBody
	GetCode() *string
	SetData(v []*ListUserDrawRecordByPkResponseBodyData) *ListUserDrawRecordByPkResponseBody
	GetData() []*ListUserDrawRecordByPkResponseBodyData
	SetMessage(v string) *ListUserDrawRecordByPkResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserDrawRecordByPkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserDrawRecordByPkResponseBody
	GetSuccess() *bool
}

type ListUserDrawRecordByPkResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListUserDrawRecordByPkResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListUserDrawRecordByPkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserDrawRecordByPkResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDrawRecordByPkResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserDrawRecordByPkResponseBody) GetData() []*ListUserDrawRecordByPkResponseBodyData {
	return s.Data
}

func (s *ListUserDrawRecordByPkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserDrawRecordByPkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserDrawRecordByPkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserDrawRecordByPkResponseBody) SetCode(v string) *ListUserDrawRecordByPkResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBody) SetData(v []*ListUserDrawRecordByPkResponseBodyData) *ListUserDrawRecordByPkResponseBody {
	s.Data = v
	return s
}

func (s *ListUserDrawRecordByPkResponseBody) SetMessage(v string) *ListUserDrawRecordByPkResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBody) SetRequestId(v string) *ListUserDrawRecordByPkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBody) SetSuccess(v bool) *ListUserDrawRecordByPkResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBody) Validate() error {
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

type ListUserDrawRecordByPkResponseBodyData struct {
	// example:
	//
	// 1401072305438324
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// example:
	//
	// dsadsadsadas
	DrawGroup *string `json:"drawGroup,omitempty" xml:"drawGroup,omitempty"`
	// example:
	//
	// dasdsadasdas
	DrawPoolName *string `json:"drawPoolName,omitempty" xml:"drawPoolName,omitempty"`
	// example:
	//
	// 2
	DrawResult *string `json:"drawResult,omitempty" xml:"drawResult,omitempty"`
	// example:
	//
	// 1545726028000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// null
	TaskGroupId *string `json:"taskGroupId,omitempty" xml:"taskGroupId,omitempty"`
	// example:
	//
	// dsadsadasd
	UccId *string `json:"uccId,omitempty" xml:"uccId,omitempty"`
}

func (s ListUserDrawRecordByPkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserDrawRecordByPkResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserDrawRecordByPkResponseBodyData) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListUserDrawRecordByPkResponseBodyData) GetDrawGroup() *string {
	return s.DrawGroup
}

func (s *ListUserDrawRecordByPkResponseBodyData) GetDrawPoolName() *string {
	return s.DrawPoolName
}

func (s *ListUserDrawRecordByPkResponseBodyData) GetDrawResult() *string {
	return s.DrawResult
}

func (s *ListUserDrawRecordByPkResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListUserDrawRecordByPkResponseBodyData) GetTaskGroupId() *string {
	return s.TaskGroupId
}

func (s *ListUserDrawRecordByPkResponseBodyData) GetUccId() *string {
	return s.UccId
}

func (s *ListUserDrawRecordByPkResponseBodyData) SetAliyunPk(v string) *ListUserDrawRecordByPkResponseBodyData {
	s.AliyunPk = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBodyData) SetDrawGroup(v string) *ListUserDrawRecordByPkResponseBodyData {
	s.DrawGroup = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBodyData) SetDrawPoolName(v string) *ListUserDrawRecordByPkResponseBodyData {
	s.DrawPoolName = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBodyData) SetDrawResult(v string) *ListUserDrawRecordByPkResponseBodyData {
	s.DrawResult = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBodyData) SetGmtCreate(v string) *ListUserDrawRecordByPkResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBodyData) SetTaskGroupId(v string) *ListUserDrawRecordByPkResponseBodyData {
	s.TaskGroupId = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBodyData) SetUccId(v string) *ListUserDrawRecordByPkResponseBodyData {
	s.UccId = &v
	return s
}

func (s *ListUserDrawRecordByPkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
