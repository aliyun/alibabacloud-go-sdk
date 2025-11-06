// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskByUidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskByUidResponseBody
	GetCode() *int32
	SetData(v *GetTaskByUidResponseBodyData) *GetTaskByUidResponseBody
	GetData() *GetTaskByUidResponseBodyData
	SetMessage(v string) *GetTaskByUidResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskByUidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskByUidResponseBody
	GetSuccess() *bool
}

type GetTaskByUidResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTaskByUidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTaskByUidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUidResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskByUidResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskByUidResponseBody) GetData() *GetTaskByUidResponseBodyData {
	return s.Data
}

func (s *GetTaskByUidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskByUidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskByUidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskByUidResponseBody) SetCode(v int32) *GetTaskByUidResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskByUidResponseBody) SetData(v *GetTaskByUidResponseBodyData) *GetTaskByUidResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskByUidResponseBody) SetMessage(v string) *GetTaskByUidResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskByUidResponseBody) SetRequestId(v string) *GetTaskByUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskByUidResponseBody) SetSuccess(v bool) *GetTaskByUidResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskByUidResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskByUidResponseBodyData struct {
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *GetTaskByUidResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s GetTaskByUidResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUidResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskByUidResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetTaskByUidResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTaskByUidResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTaskByUidResponseBodyData) GetVoList() *GetTaskByUidResponseBodyDataVoList {
	return s.VoList
}

func (s *GetTaskByUidResponseBodyData) SetCurrentPage(v int32) *GetTaskByUidResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetTaskByUidResponseBodyData) SetPageSize(v int32) *GetTaskByUidResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetTaskByUidResponseBodyData) SetTotalCount(v int64) *GetTaskByUidResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetTaskByUidResponseBodyData) SetVoList(v *GetTaskByUidResponseBodyDataVoList) *GetTaskByUidResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetTaskByUidResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskByUidResponseBodyDataVoList struct {
	ImportDefinitionTaskDO []*GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO `json:"ImportDefinitionTaskDO,omitempty" xml:"ImportDefinitionTaskDO,omitempty" type:"Repeated"`
}

func (s GetTaskByUidResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUidResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetTaskByUidResponseBodyDataVoList) GetImportDefinitionTaskDO() []*GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	return s.ImportDefinitionTaskDO
}

func (s *GetTaskByUidResponseBodyDataVoList) SetImportDefinitionTaskDO(v []*GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) *GetTaskByUidResponseBodyDataVoList {
	s.ImportDefinitionTaskDO = v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoList) Validate() error {
	if s.ImportDefinitionTaskDO != nil {
		for _, item := range s.ImportDefinitionTaskDO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO struct {
	BindingNum   *int32  `json:"BindingNum,omitempty" xml:"BindingNum,omitempty"`
	ExchangeNum  *int32  `json:"ExchangeNum,omitempty" xml:"ExchangeNum,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ImportType   *int32  `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	QueueNum     *int32  `json:"QueueNum,omitempty" xml:"QueueNum,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VhostName    *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
	VhostNum     *int32  `json:"VhostNum,omitempty" xml:"VhostNum,omitempty"`
}

func (s GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GoString() string {
	return s.String()
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetBindingNum() *int32 {
	return s.BindingNum
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetExchangeNum() *int32 {
	return s.ExchangeNum
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetId() *int64 {
	return s.Id
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetImportType() *int32 {
	return s.ImportType
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetQueueNum() *int32 {
	return s.QueueNum
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetStatus() *int32 {
	return s.Status
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetUserId() *int64 {
	return s.UserId
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetVhostName() *string {
	return s.VhostName
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) GetVhostNum() *int32 {
	return s.VhostNum
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetBindingNum(v int32) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.BindingNum = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetExchangeNum(v int32) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.ExchangeNum = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetGmtCreate(v string) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.GmtCreate = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetId(v int64) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.Id = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetImportType(v int32) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.ImportType = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetInstanceId(v string) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.InstanceId = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetInstanceName(v string) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.InstanceName = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetQueueNum(v int32) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.QueueNum = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetStatus(v int32) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.Status = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetUserId(v int64) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.UserId = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetVhostName(v string) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.VhostName = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) SetVhostNum(v int32) *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO {
	s.VhostNum = &v
	return s
}

func (s *GetTaskByUidResponseBodyDataVoListImportDefinitionTaskDO) Validate() error {
	return dara.Validate(s)
}
