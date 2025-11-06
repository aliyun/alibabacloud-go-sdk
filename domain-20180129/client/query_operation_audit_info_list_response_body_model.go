// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOperationAuditInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryOperationAuditInfoListResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*QueryOperationAuditInfoListResponseBodyData) *QueryOperationAuditInfoListResponseBody
	GetData() []*QueryOperationAuditInfoListResponseBodyData
	SetNextPage(v bool) *QueryOperationAuditInfoListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryOperationAuditInfoListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryOperationAuditInfoListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryOperationAuditInfoListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryOperationAuditInfoListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryOperationAuditInfoListResponseBody
	GetTotalPageNum() *int32
}

type QueryOperationAuditInfoListResponseBody struct {
	// example:
	//
	// 2
	CurrentPageNum *int32                                         `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryOperationAuditInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// 9DFCF6F8-243C-40EC-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 199
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 10
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryOperationAuditInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOperationAuditInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryOperationAuditInfoListResponseBody) GetData() []*QueryOperationAuditInfoListResponseBodyData {
	return s.Data
}

func (s *QueryOperationAuditInfoListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryOperationAuditInfoListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOperationAuditInfoListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryOperationAuditInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOperationAuditInfoListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryOperationAuditInfoListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryOperationAuditInfoListResponseBody) SetCurrentPageNum(v int32) *QueryOperationAuditInfoListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetData(v []*QueryOperationAuditInfoListResponseBodyData) *QueryOperationAuditInfoListResponseBody {
	s.Data = v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetNextPage(v bool) *QueryOperationAuditInfoListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetPageSize(v int32) *QueryOperationAuditInfoListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetPrePage(v bool) *QueryOperationAuditInfoListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetRequestId(v string) *QueryOperationAuditInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetTotalItemNum(v int32) *QueryOperationAuditInfoListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetTotalPageNum(v int32) *QueryOperationAuditInfoListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) Validate() error {
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

type QueryOperationAuditInfoListResponseBodyData struct {
	AuditInfo *string `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty"`
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// 1
	AuditType    *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 1581919010101
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example.com,aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1581919010101
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryOperationAuditInfoListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryOperationAuditInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetAuditInfo() *string {
	return s.AuditInfo
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetAuditType() *int32 {
	return s.AuditType
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetBusinessName() *string {
	return s.BusinessName
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *QueryOperationAuditInfoListResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetAuditInfo(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.AuditInfo = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetAuditStatus(v int32) *QueryOperationAuditInfoListResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetAuditType(v int32) *QueryOperationAuditInfoListResponseBodyData {
	s.AuditType = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetBusinessName(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetCreateTime(v int64) *QueryOperationAuditInfoListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetDomainName(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetId(v int64) *QueryOperationAuditInfoListResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetRemark(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetUpdateTime(v int64) *QueryOperationAuditInfoListResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
