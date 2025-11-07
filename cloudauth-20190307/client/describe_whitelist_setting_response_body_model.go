// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWhitelistSettingResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeWhitelistSettingResponseBodyItems) *DescribeWhitelistSettingResponseBody
	GetItems() []*DescribeWhitelistSettingResponseBodyItems
	SetPageSize(v int32) *DescribeWhitelistSettingResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWhitelistSettingResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWhitelistSettingResponseBody
	GetTotalCount() *int32
}

type DescribeWhitelistSettingResponseBody struct {
	// Pagination parameter: current page number, default value is 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of certification details.
	Items []*DescribeWhitelistSettingResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Number of items per page for pagination.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of this request.
	//
	// example:
	//
	// 0B8ACFD2-C5F0-5F9F-8DD4-E44F93E360E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhitelistSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhitelistSettingResponseBody) GetItems() []*DescribeWhitelistSettingResponseBodyItems {
	return s.Items
}

func (s *DescribeWhitelistSettingResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhitelistSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhitelistSettingResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWhitelistSettingResponseBody) SetCurrentPage(v int32) *DescribeWhitelistSettingResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetItems(v []*DescribeWhitelistSettingResponseBodyItems) *DescribeWhitelistSettingResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetPageSize(v int32) *DescribeWhitelistSettingResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetRequestId(v string) *DescribeWhitelistSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) SetTotalCount(v int32) *DescribeWhitelistSettingResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWhitelistSettingResponseBodyItems struct {
	// Certificate number.
	//
	// example:
	//
	// 320321XXXXXXXX701X
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// Certificate ID.
	//
	// example:
	//
	// shad861465f2aaeeb805b519e1a93ab2
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2025-10-16 17:28:03
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2025-10-16 17:28:03
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Whitelist ID.
	//
	// example:
	//
	// 6372003
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Remark.
	//
	// example:
	//
	// 测试白名单
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000000332
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Whitelist status:
	//
	// - **VALID**: Valid.
	//
	// - **INVALID**: Invalid.
	//
	// - **DELETED**: Deleted.
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Effective end date.
	//
	// example:
	//
	// 2025-10-16 17:28:03
	ValidEndDate *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// Effective start time.
	//
	// example:
	//
	// 2025-10-16 17:28:03
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s DescribeWhitelistSettingResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistSettingResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetCertNo() *string {
	return s.CertNo
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetRemark() *string {
	return s.Remark
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetValidEndDate() *string {
	return s.ValidEndDate
}

func (s *DescribeWhitelistSettingResponseBodyItems) GetValidStartDate() *string {
	return s.ValidStartDate
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetCertNo(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.CertNo = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetCertifyId(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.CertifyId = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetGmtCreate(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetGmtModified(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetId(v int64) *DescribeWhitelistSettingResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetRemark(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.Remark = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetSceneId(v int64) *DescribeWhitelistSettingResponseBodyItems {
	s.SceneId = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetStatus(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetValidEndDate(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.ValidEndDate = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) SetValidStartDate(v string) *DescribeWhitelistSettingResponseBodyItems {
	s.ValidStartDate = &v
	return s
}

func (s *DescribeWhitelistSettingResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
