// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyMeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetNotifyMeResponseBodyData) *GetNotifyMeResponseBody
	GetData() []*GetNotifyMeResponseBodyData
	SetPageNumber(v int64) *GetNotifyMeResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *GetNotifyMeResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetNotifyMeResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetNotifyMeResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetNotifyMeResponseBody
	GetVendorType() *string
}

type GetNotifyMeResponseBody struct {
	// example:
	//
	// [{}]
	Data []*GetNotifyMeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s GetNotifyMeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyMeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotifyMeResponseBody) GetData() []*GetNotifyMeResponseBodyData {
	return s.Data
}

func (s *GetNotifyMeResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetNotifyMeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNotifyMeResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetNotifyMeResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetNotifyMeResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetNotifyMeResponseBody) SetData(v []*GetNotifyMeResponseBodyData) *GetNotifyMeResponseBody {
	s.Data = v
	return s
}

func (s *GetNotifyMeResponseBody) SetPageNumber(v int64) *GetNotifyMeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetNotifyMeResponseBody) SetRequestId(v string) *GetNotifyMeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotifyMeResponseBody) SetTotalCount(v int64) *GetNotifyMeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetNotifyMeResponseBody) SetVendorRequestId(v string) *GetNotifyMeResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetNotifyMeResponseBody) SetVendorType(v string) *GetNotifyMeResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetNotifyMeResponseBody) Validate() error {
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

type GetNotifyMeResponseBodyData struct {
	// example:
	//
	// act-xxaanfaf
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// APP_XCxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// corpIdxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 2020-01-01
	CreateTimeGMT *string `json:"CreateTimeGMT,omitempty" xml:"CreateTimeGMT,omitempty"`
	// example:
	//
	// 123456
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// example:
	//
	// formxxxx
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// example:
	//
	// RUNNING
	InstStatus *string `json:"InstStatus,omitempty" xml:"InstStatus,omitempty"`
	// example:
	//
	// mobileUrlexample
	MobileUrl *string `json:"MobileUrl,omitempty" xml:"MobileUrl,omitempty"`
	// example:
	//
	// 2020-01-01
	ModifiedTimeGMT *string `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
}

func (s GetNotifyMeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyMeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNotifyMeResponseBodyData) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetNotifyMeResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *GetNotifyMeResponseBodyData) GetCorpId() *string {
	return s.CorpId
}

func (s *GetNotifyMeResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetNotifyMeResponseBodyData) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *GetNotifyMeResponseBodyData) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *GetNotifyMeResponseBodyData) GetInstStatus() *string {
	return s.InstStatus
}

func (s *GetNotifyMeResponseBodyData) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *GetNotifyMeResponseBodyData) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *GetNotifyMeResponseBodyData) SetActivityId(v string) *GetNotifyMeResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetAppType(v string) *GetNotifyMeResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetCorpId(v string) *GetNotifyMeResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetCreateTimeGMT(v string) *GetNotifyMeResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetCreatorUserId(v string) *GetNotifyMeResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetFormInstanceId(v string) *GetNotifyMeResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetInstStatus(v string) *GetNotifyMeResponseBodyData {
	s.InstStatus = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetMobileUrl(v string) *GetNotifyMeResponseBodyData {
	s.MobileUrl = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetModifiedTimeGMT(v string) *GetNotifyMeResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
