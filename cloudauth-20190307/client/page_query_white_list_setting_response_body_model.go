// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryWhiteListSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PageQueryWhiteListSettingResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *PageQueryWhiteListSettingResponseBody
	GetCurrentPage() *int32
	SetMessage(v string) *PageQueryWhiteListSettingResponseBody
	GetMessage() *string
	SetPageSize(v int32) *PageQueryWhiteListSettingResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *PageQueryWhiteListSettingResponseBody
	GetRequestId() *string
	SetResultObject(v []*PageQueryWhiteListSettingResponseBodyResultObject) *PageQueryWhiteListSettingResponseBody
	GetResultObject() []*PageQueryWhiteListSettingResponseBodyResultObject
	SetSuccess(v bool) *PageQueryWhiteListSettingResponseBody
	GetSuccess() *bool
	SetTotalItem(v int32) *PageQueryWhiteListSettingResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *PageQueryWhiteListSettingResponseBody
	GetTotalPage() *int32
}

type PageQueryWhiteListSettingResponseBody struct {
	// The response code. **200*	- indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result.
	ResultObject []*PageQueryWhiteListSettingResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 28
	TotalItem *int32 `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s PageQueryWhiteListSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageQueryWhiteListSettingResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *PageQueryWhiteListSettingResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *PageQueryWhiteListSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PageQueryWhiteListSettingResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageQueryWhiteListSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageQueryWhiteListSettingResponseBody) GetResultObject() []*PageQueryWhiteListSettingResponseBodyResultObject {
	return s.ResultObject
}

func (s *PageQueryWhiteListSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PageQueryWhiteListSettingResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *PageQueryWhiteListSettingResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *PageQueryWhiteListSettingResponseBody) SetCode(v string) *PageQueryWhiteListSettingResponseBody {
	s.Code = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetCurrentPage(v int32) *PageQueryWhiteListSettingResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetMessage(v string) *PageQueryWhiteListSettingResponseBody {
	s.Message = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetPageSize(v int32) *PageQueryWhiteListSettingResponseBody {
	s.PageSize = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetRequestId(v string) *PageQueryWhiteListSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetResultObject(v []*PageQueryWhiteListSettingResponseBodyResultObject) *PageQueryWhiteListSettingResponseBody {
	s.ResultObject = v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetSuccess(v bool) *PageQueryWhiteListSettingResponseBody {
	s.Success = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetTotalItem(v int32) *PageQueryWhiteListSettingResponseBody {
	s.TotalItem = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) SetTotalPage(v int32) *PageQueryWhiteListSettingResponseBody {
	s.TotalPage = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBody) Validate() error {
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

type PageQueryWhiteListSettingResponseBodyResultObject struct {
	// The certificate number.
	//
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// The unique identifier of the ID Verification request.
	//
	// example:
	//
	// sha43d9cabd52d370d9f4cca9468f71e
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-08-30 14:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-08-30 14:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The whitelist ID.
	//
	// example:
	//
	// 234822
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The verification scenario ID.
	//
	// example:
	//
	// 1000000332
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The ServiceCode of the ID Verification cloud service.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The status. Valid values:
	//
	// - DELETE: Deleted.
	//
	// - VALID: Not deleted and within the validity period (valid).
	//
	// - INVALID: Not deleted but outside the validity period (invalid).
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The end date of the validity period.
	//
	// example:
	//
	// 2024-09-02 13:57:51
	ValidEndDate *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// The start date of the validity period.
	//
	// example:
	//
	// 2024-08-30 13:57:51
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s PageQueryWhiteListSettingResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s PageQueryWhiteListSettingResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetCertNo() *string {
	return s.CertNo
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetRemark() *string {
	return s.Remark
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetSceneId() *int64 {
	return s.SceneId
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetValidEndDate() *string {
	return s.ValidEndDate
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) GetValidStartDate() *string {
	return s.ValidStartDate
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetCertNo(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.CertNo = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetCertifyId(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetGmtCreate(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetGmtModified(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetId(v int64) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetRemark(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetSceneId(v int64) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.SceneId = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetServiceCode(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.ServiceCode = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetStatus(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetValidEndDate(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.ValidEndDate = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) SetValidStartDate(v string) *PageQueryWhiteListSettingResponseBodyResultObject {
	s.ValidStartDate = &v
	return s
}

func (s *PageQueryWhiteListSettingResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
