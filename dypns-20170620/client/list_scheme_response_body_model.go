// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSchemeResponseBody
	GetCode() *string
	SetData(v []*ListSchemeResponseBodyData) *ListSchemeResponseBody
	GetData() []*ListSchemeResponseBodyData
	SetPageNumber(v int32) *ListSchemeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSchemeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSchemeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSchemeResponseBody
	GetTotalCount() *int32
}

type ListSchemeResponseBody struct {
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSchemeResponseBody) GetData() []*ListSchemeResponseBodyData {
	return s.Data
}

func (s *ListSchemeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSchemeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchemeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSchemeResponseBody) SetCode(v string) *ListSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *ListSchemeResponseBody) SetData(v []*ListSchemeResponseBodyData) *ListSchemeResponseBody {
	s.Data = v
	return s
}

func (s *ListSchemeResponseBody) SetPageNumber(v int32) *ListSchemeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSchemeResponseBody) SetPageSize(v int32) *ListSchemeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSchemeResponseBody) SetRequestId(v string) *ListSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemeResponseBody) SetTotalCount(v int32) *ListSchemeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSchemeResponseBody) Validate() error {
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

type ListSchemeResponseBodyData struct {
	ApplyTime        *int64   `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	AuditDesc        *string  `json:"AuditDesc,omitempty" xml:"AuditDesc,omitempty"`
	BusinessTypeList []*int32 `json:"BusinessTypeList,omitempty" xml:"BusinessTypeList,omitempty" type:"Repeated"`
	CompanyName      *string  `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	SchemeId         *int64   `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	SchemeName       *string  `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	SchemeType       *int32   `json:"SchemeType,omitempty" xml:"SchemeType,omitempty"`
	Status           *int32   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSchemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSchemeResponseBodyData) GetApplyTime() *int64 {
	return s.ApplyTime
}

func (s *ListSchemeResponseBodyData) GetAuditDesc() *string {
	return s.AuditDesc
}

func (s *ListSchemeResponseBodyData) GetBusinessTypeList() []*int32 {
	return s.BusinessTypeList
}

func (s *ListSchemeResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *ListSchemeResponseBodyData) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *ListSchemeResponseBodyData) GetSchemeName() *string {
	return s.SchemeName
}

func (s *ListSchemeResponseBodyData) GetSchemeType() *int32 {
	return s.SchemeType
}

func (s *ListSchemeResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListSchemeResponseBodyData) SetApplyTime(v int64) *ListSchemeResponseBodyData {
	s.ApplyTime = &v
	return s
}

func (s *ListSchemeResponseBodyData) SetAuditDesc(v string) *ListSchemeResponseBodyData {
	s.AuditDesc = &v
	return s
}

func (s *ListSchemeResponseBodyData) SetBusinessTypeList(v []*int32) *ListSchemeResponseBodyData {
	s.BusinessTypeList = v
	return s
}

func (s *ListSchemeResponseBodyData) SetCompanyName(v string) *ListSchemeResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *ListSchemeResponseBodyData) SetSchemeId(v int64) *ListSchemeResponseBodyData {
	s.SchemeId = &v
	return s
}

func (s *ListSchemeResponseBodyData) SetSchemeName(v string) *ListSchemeResponseBodyData {
	s.SchemeName = &v
	return s
}

func (s *ListSchemeResponseBodyData) SetSchemeType(v int32) *ListSchemeResponseBodyData {
	s.SchemeType = &v
	return s
}

func (s *ListSchemeResponseBodyData) SetStatus(v int32) *ListSchemeResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListSchemeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
