// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpAccomplishmentTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetCorpAccomplishmentTasksResponseBodyData) *GetCorpAccomplishmentTasksResponseBody
	GetData() []*GetCorpAccomplishmentTasksResponseBodyData
	SetPageNumber(v int64) *GetCorpAccomplishmentTasksResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *GetCorpAccomplishmentTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetCorpAccomplishmentTasksResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetCorpAccomplishmentTasksResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetCorpAccomplishmentTasksResponseBody
	GetVendorType() *string
}

type GetCorpAccomplishmentTasksResponseBody struct {
	// example:
	//
	// [{}]
	Data []*GetCorpAccomplishmentTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s GetCorpAccomplishmentTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCorpAccomplishmentTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksResponseBody) GetData() []*GetCorpAccomplishmentTasksResponseBodyData {
	return s.Data
}

func (s *GetCorpAccomplishmentTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetCorpAccomplishmentTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCorpAccomplishmentTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetCorpAccomplishmentTasksResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetCorpAccomplishmentTasksResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetData(v []*GetCorpAccomplishmentTasksResponseBodyData) *GetCorpAccomplishmentTasksResponseBody {
	s.Data = v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetPageNumber(v int64) *GetCorpAccomplishmentTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetRequestId(v string) *GetCorpAccomplishmentTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetTotalCount(v int64) *GetCorpAccomplishmentTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetVendorRequestId(v string) *GetCorpAccomplishmentTasksResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetVendorType(v string) *GetCorpAccomplishmentTasksResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) Validate() error {
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

type GetCorpAccomplishmentTasksResponseBodyData struct {
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
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetCorpAccomplishmentTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCorpAccomplishmentTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetActiveTimeGMT() *string {
	return s.ActiveTimeGMT
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetActualActionerId() *string {
	return s.ActualActionerId
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetFinishTimeGMT() *string {
	return s.FinishTimeGMT
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOriginatorEmail() *string {
	return s.OriginatorEmail
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOriginatorName() *string {
	return s.OriginatorName
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOriginatorNameInEnglish() *string {
	return s.OriginatorNameInEnglish
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOriginatorNickName() *string {
	return s.OriginatorNickName
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOriginatorNickNameInEnglish() *string {
	return s.OriginatorNickNameInEnglish
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOriginatorPhoto() *string {
	return s.OriginatorPhoto
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOutResult() *string {
	return s.OutResult
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetOutResultName() *string {
	return s.OutResultName
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetActiveTimeGMT(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetActualActionerId(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.ActualActionerId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetAppType(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetCreateTimeGMT(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetFinishTimeGMT(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorEmail(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorEmail = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorId(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorName(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorName = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorNameInEnglish(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorNameInEnglish = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorNickName(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorNickName = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorNickNameInEnglish(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorNickNameInEnglish = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorPhoto(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorPhoto = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOutResult(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OutResult = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOutResultName(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OutResultName = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetProcessInstanceId(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetTitle(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
