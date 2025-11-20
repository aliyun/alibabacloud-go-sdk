// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormListInAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetFormListInAppResponseBody
	GetCurrentPage() *int32
	SetData(v []*GetFormListInAppResponseBodyData) *GetFormListInAppResponseBody
	GetData() []*GetFormListInAppResponseBodyData
	SetRequestId(v string) *GetFormListInAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFormListInAppResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetFormListInAppResponseBody
	GetTotalCount() *int32
	SetVendorRequestId(v string) *GetFormListInAppResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetFormListInAppResponseBody
	GetVendorType() *string
}

type GetFormListInAppResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                              `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Data        []*GetFormListInAppResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetFormListInAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetFormListInAppResponseBody) GetData() []*GetFormListInAppResponseBodyData {
	return s.Data
}

func (s *GetFormListInAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFormListInAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFormListInAppResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetFormListInAppResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetFormListInAppResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetFormListInAppResponseBody) SetCurrentPage(v int32) *GetFormListInAppResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetFormListInAppResponseBody) SetData(v []*GetFormListInAppResponseBodyData) *GetFormListInAppResponseBody {
	s.Data = v
	return s
}

func (s *GetFormListInAppResponseBody) SetRequestId(v string) *GetFormListInAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFormListInAppResponseBody) SetSuccess(v bool) *GetFormListInAppResponseBody {
	s.Success = &v
	return s
}

func (s *GetFormListInAppResponseBody) SetTotalCount(v int32) *GetFormListInAppResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetFormListInAppResponseBody) SetVendorRequestId(v string) *GetFormListInAppResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetFormListInAppResponseBody) SetVendorType(v string) *GetFormListInAppResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetFormListInAppResponseBody) Validate() error {
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

type GetFormListInAppResponseBodyData struct {
	// example:
	//
	// 012345
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// receipt
	FormType *string `json:"FormType,omitempty" xml:"FormType,omitempty"`
	// example:
	//
	// FORM-BAxxxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// 2023-02-22 15:27:07
	GmtCreate *string                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Title     *GetFormListInAppResponseBodyDataTitle `json:"Title,omitempty" xml:"Title,omitempty" type:"Struct"`
}

func (s GetFormListInAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetFormListInAppResponseBodyData) GetFormType() *string {
	return s.FormType
}

func (s *GetFormListInAppResponseBodyData) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetFormListInAppResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetFormListInAppResponseBodyData) GetTitle() *GetFormListInAppResponseBodyDataTitle {
	return s.Title
}

func (s *GetFormListInAppResponseBodyData) SetCreator(v string) *GetFormListInAppResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetFormListInAppResponseBodyData) SetFormType(v string) *GetFormListInAppResponseBodyData {
	s.FormType = &v
	return s
}

func (s *GetFormListInAppResponseBodyData) SetFormUuid(v string) *GetFormListInAppResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetFormListInAppResponseBodyData) SetGmtCreate(v string) *GetFormListInAppResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetFormListInAppResponseBodyData) SetTitle(v *GetFormListInAppResponseBodyDataTitle) *GetFormListInAppResponseBodyData {
	s.Title = v
	return s
}

func (s *GetFormListInAppResponseBodyData) Validate() error {
	if s.Title != nil {
		if err := s.Title.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFormListInAppResponseBodyDataTitle struct {
	// example:
	//
	// form
	EnUS *string `json:"EnUS,omitempty" xml:"EnUS,omitempty"`
	ZhCN *string `json:"ZhCN,omitempty" xml:"ZhCN,omitempty"`
}

func (s GetFormListInAppResponseBodyDataTitle) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppResponseBodyDataTitle) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponseBodyDataTitle) GetEnUS() *string {
	return s.EnUS
}

func (s *GetFormListInAppResponseBodyDataTitle) GetZhCN() *string {
	return s.ZhCN
}

func (s *GetFormListInAppResponseBodyDataTitle) SetEnUS(v string) *GetFormListInAppResponseBodyDataTitle {
	s.EnUS = &v
	return s
}

func (s *GetFormListInAppResponseBodyDataTitle) SetZhCN(v string) *GetFormListInAppResponseBodyDataTitle {
	s.ZhCN = &v
	return s
}

func (s *GetFormListInAppResponseBodyDataTitle) Validate() error {
	return dara.Validate(s)
}
