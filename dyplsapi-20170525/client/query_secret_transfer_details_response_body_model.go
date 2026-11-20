// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretTransferDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySecretTransferDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySecretTransferDetailsResponseBody
	GetCode() *string
	SetData(v *QuerySecretTransferDetailsResponseBodyData) *QuerySecretTransferDetailsResponseBody
	GetData() *QuerySecretTransferDetailsResponseBodyData
	SetMessage(v string) *QuerySecretTransferDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySecretTransferDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySecretTransferDetailsResponseBody
	GetSuccess() *bool
}

type QuerySecretTransferDetailsResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QuerySecretTransferDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 123456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySecretTransferDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretTransferDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretTransferDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySecretTransferDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySecretTransferDetailsResponseBody) GetData() *QuerySecretTransferDetailsResponseBodyData {
	return s.Data
}

func (s *QuerySecretTransferDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySecretTransferDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySecretTransferDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySecretTransferDetailsResponseBody) SetAccessDeniedDetail(v string) *QuerySecretTransferDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBody) SetCode(v string) *QuerySecretTransferDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBody) SetData(v *QuerySecretTransferDetailsResponseBodyData) *QuerySecretTransferDetailsResponseBody {
	s.Data = v
	return s
}

func (s *QuerySecretTransferDetailsResponseBody) SetMessage(v string) *QuerySecretTransferDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBody) SetRequestId(v string) *QuerySecretTransferDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBody) SetSuccess(v bool) *QuerySecretTransferDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySecretTransferDetailsResponseBodyData struct {
	List []*QuerySecretTransferDetailsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 87
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QuerySecretTransferDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretTransferDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySecretTransferDetailsResponseBodyData) GetList() []*QuerySecretTransferDetailsResponseBodyDataList {
	return s.List
}

func (s *QuerySecretTransferDetailsResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QuerySecretTransferDetailsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QuerySecretTransferDetailsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QuerySecretTransferDetailsResponseBodyData) SetList(v []*QuerySecretTransferDetailsResponseBodyDataList) *QuerySecretTransferDetailsResponseBodyData {
	s.List = v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyData) SetPageNo(v int64) *QuerySecretTransferDetailsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyData) SetPageSize(v int64) *QuerySecretTransferDetailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyData) SetTotal(v int64) *QuerySecretTransferDetailsResponseBodyData {
	s.Total = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySecretTransferDetailsResponseBodyDataList struct {
	// example:
	//
	// 杭州
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 号码不存在
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// example:
	//
	// "2026-11-20 00:00:00"
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// "2026-11-20 00:00:00"
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 90
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 46
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 138XXXXXX
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuerySecretTransferDetailsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretTransferDetailsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetCity() *string {
	return s.City
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetFailReason() *string {
	return s.FailReason
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetRecordId() *int64 {
	return s.RecordId
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetSecretNo() *string {
	return s.SecretNo
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) GetStatus() *int64 {
	return s.Status
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetCity(v string) *QuerySecretTransferDetailsResponseBodyDataList {
	s.City = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetFailReason(v string) *QuerySecretTransferDetailsResponseBodyDataList {
	s.FailReason = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetGmtCreate(v string) *QuerySecretTransferDetailsResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetGmtModified(v string) *QuerySecretTransferDetailsResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetId(v int64) *QuerySecretTransferDetailsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetRecordId(v int64) *QuerySecretTransferDetailsResponseBodyDataList {
	s.RecordId = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetSecretNo(v string) *QuerySecretTransferDetailsResponseBodyDataList {
	s.SecretNo = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) SetStatus(v int64) *QuerySecretTransferDetailsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QuerySecretTransferDetailsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
