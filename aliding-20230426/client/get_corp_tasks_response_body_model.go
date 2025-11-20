// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetCorpTasksResponseBodyData) *GetCorpTasksResponseBody
	GetData() []*GetCorpTasksResponseBodyData
	SetPageNumber(v int64) *GetCorpTasksResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *GetCorpTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetCorpTasksResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetCorpTasksResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetCorpTasksResponseBody
	GetVendorType() *string
}

type GetCorpTasksResponseBody struct {
	// example:
	//
	// [{}]
	Data []*GetCorpTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetCorpTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCorpTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpTasksResponseBody) GetData() []*GetCorpTasksResponseBodyData {
	return s.Data
}

func (s *GetCorpTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetCorpTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCorpTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetCorpTasksResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetCorpTasksResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetCorpTasksResponseBody) SetData(v []*GetCorpTasksResponseBodyData) *GetCorpTasksResponseBody {
	s.Data = v
	return s
}

func (s *GetCorpTasksResponseBody) SetPageNumber(v int64) *GetCorpTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetCorpTasksResponseBody) SetRequestId(v string) *GetCorpTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCorpTasksResponseBody) SetTotalCount(v int64) *GetCorpTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetCorpTasksResponseBody) SetVendorRequestId(v string) *GetCorpTasksResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetCorpTasksResponseBody) SetVendorType(v string) *GetCorpTasksResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetCorpTasksResponseBody) Validate() error {
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

type GetCorpTasksResponseBodyData struct {
	// example:
	//
	// 2020-01-01
	ActiveTimeGMT *string `json:"ActiveTimeGMT,omitempty" xml:"ActiveTimeGMT,omitempty"`
	// example:
	//
	// 123456
	ActualActionerId *string `json:"ActualActionerId,omitempty" xml:"ActualActionerId,omitempty"`
	// example:
	//
	// APP_XCxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 2020-01-01
	CreateTimeGMT *string `json:"CreateTimeGMT,omitempty" xml:"CreateTimeGMT,omitempty"`
	// example:
	//
	// 2020-01-01
	FinishTimeGMT *string `json:"FinishTimeGMT,omitempty" xml:"FinishTimeGMT,omitempty"`
	// example:
	//
	// 123456@li.com
	OriginatorEmail *string `json:"OriginatorEmail,omitempty" xml:"OriginatorEmail,omitempty"`
	// example:
	//
	// 123456
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// 名称
	OriginatorName *string `json:"OriginatorName,omitempty" xml:"OriginatorName,omitempty"`
	// example:
	//
	// name
	OriginatorNameInEnglish *string `json:"OriginatorNameInEnglish,omitempty" xml:"OriginatorNameInEnglish,omitempty"`
	// example:
	//
	// 昵称
	OriginatorNickName *string `json:"OriginatorNickName,omitempty" xml:"OriginatorNickName,omitempty"`
	// example:
	//
	// nick en
	OriginatorNickNameEn *string `json:"OriginatorNickNameEn,omitempty" xml:"OriginatorNickNameEn,omitempty"`
	// example:
	//
	// english nick
	OriginatorNickNameInEnglish *string `json:"OriginatorNickNameInEnglish,omitempty" xml:"OriginatorNickNameInEnglish,omitempty"`
	// example:
	//
	// originatorPhotoexample
	OriginatorPhoto *string `json:"OriginatorPhoto,omitempty" xml:"OriginatorPhoto,omitempty"`
	// example:
	//
	// agree
	OutResult *string `json:"OutResult,omitempty" xml:"OutResult,omitempty"`
	// example:
	//
	// 结果名称
	OutResultName *string `json:"OutResultName,omitempty" xml:"OutResultName,omitempty"`
	// example:
	//
	// instancexxxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// taskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// running
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// title
	TitleInEnglish *string `json:"TitleInEnglish,omitempty" xml:"TitleInEnglish,omitempty"`
}

func (s GetCorpTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCorpTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCorpTasksResponseBodyData) GetActiveTimeGMT() *string {
	return s.ActiveTimeGMT
}

func (s *GetCorpTasksResponseBodyData) GetActualActionerId() *string {
	return s.ActualActionerId
}

func (s *GetCorpTasksResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *GetCorpTasksResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetCorpTasksResponseBodyData) GetFinishTimeGMT() *string {
	return s.FinishTimeGMT
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorEmail() *string {
	return s.OriginatorEmail
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorName() *string {
	return s.OriginatorName
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorNameInEnglish() *string {
	return s.OriginatorNameInEnglish
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorNickName() *string {
	return s.OriginatorNickName
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorNickNameEn() *string {
	return s.OriginatorNickNameEn
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorNickNameInEnglish() *string {
	return s.OriginatorNickNameInEnglish
}

func (s *GetCorpTasksResponseBodyData) GetOriginatorPhoto() *string {
	return s.OriginatorPhoto
}

func (s *GetCorpTasksResponseBodyData) GetOutResult() *string {
	return s.OutResult
}

func (s *GetCorpTasksResponseBodyData) GetOutResultName() *string {
	return s.OutResultName
}

func (s *GetCorpTasksResponseBodyData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetCorpTasksResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCorpTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCorpTasksResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetCorpTasksResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetCorpTasksResponseBodyData) GetTitleInEnglish() *string {
	return s.TitleInEnglish
}

func (s *GetCorpTasksResponseBodyData) SetActiveTimeGMT(v string) *GetCorpTasksResponseBodyData {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetActualActionerId(v string) *GetCorpTasksResponseBodyData {
	s.ActualActionerId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetAppType(v string) *GetCorpTasksResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetCreateTimeGMT(v string) *GetCorpTasksResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetFinishTimeGMT(v string) *GetCorpTasksResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorEmail(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorEmail = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorId(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorName(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorName = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorNameInEnglish(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorNameInEnglish = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorNickName(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorNickName = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorNickNameEn(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorNickNameEn = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorNickNameInEnglish(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorNickNameInEnglish = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorPhoto(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorPhoto = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOutResult(v string) *GetCorpTasksResponseBodyData {
	s.OutResult = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOutResultName(v string) *GetCorpTasksResponseBodyData {
	s.OutResultName = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetProcessInstanceId(v string) *GetCorpTasksResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetStatus(v string) *GetCorpTasksResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTaskId(v string) *GetCorpTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTaskType(v string) *GetCorpTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTitle(v string) *GetCorpTasksResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTitleInEnglish(v string) *GetCorpTasksResponseBodyData {
	s.TitleInEnglish = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
