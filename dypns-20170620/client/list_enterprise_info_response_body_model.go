// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEnterpriseInfoResponseBody
	GetCode() *string
	SetData(v []*ListEnterpriseInfoResponseBodyData) *ListEnterpriseInfoResponseBody
	GetData() []*ListEnterpriseInfoResponseBodyData
	SetPageNumber(v int32) *ListEnterpriseInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEnterpriseInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEnterpriseInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEnterpriseInfoResponseBody
	GetTotalCount() *int32
}

type ListEnterpriseInfoResponseBody struct {
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListEnterpriseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEnterpriseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEnterpriseInfoResponseBody) GetData() []*ListEnterpriseInfoResponseBodyData {
	return s.Data
}

func (s *ListEnterpriseInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEnterpriseInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnterpriseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterpriseInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEnterpriseInfoResponseBody) SetCode(v string) *ListEnterpriseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseInfoResponseBody) SetData(v []*ListEnterpriseInfoResponseBodyData) *ListEnterpriseInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterpriseInfoResponseBody) SetPageNumber(v int32) *ListEnterpriseInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEnterpriseInfoResponseBody) SetPageSize(v int32) *ListEnterpriseInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEnterpriseInfoResponseBody) SetRequestId(v string) *ListEnterpriseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseInfoResponseBody) SetTotalCount(v int32) *ListEnterpriseInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEnterpriseInfoResponseBody) Validate() error {
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

type ListEnterpriseInfoResponseBodyData struct {
	AuditDesc            *string `json:"AuditDesc,omitempty" xml:"AuditDesc,omitempty"`
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	EnterpriseName       *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	ManagerContactNumber *string `json:"ManagerContactNumber,omitempty" xml:"ManagerContactNumber,omitempty"`
	ManagerName          *string `json:"ManagerName,omitempty" xml:"ManagerName,omitempty"`
	OrganizationCode     *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEnterpriseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseInfoResponseBodyData) GetAuditDesc() *string {
	return s.AuditDesc
}

func (s *ListEnterpriseInfoResponseBodyData) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListEnterpriseInfoResponseBodyData) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *ListEnterpriseInfoResponseBodyData) GetManagerContactNumber() *string {
	return s.ManagerContactNumber
}

func (s *ListEnterpriseInfoResponseBodyData) GetManagerName() *string {
	return s.ManagerName
}

func (s *ListEnterpriseInfoResponseBodyData) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *ListEnterpriseInfoResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListEnterpriseInfoResponseBodyData) SetAuditDesc(v string) *ListEnterpriseInfoResponseBodyData {
	s.AuditDesc = &v
	return s
}

func (s *ListEnterpriseInfoResponseBodyData) SetEnterpriseId(v int64) *ListEnterpriseInfoResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *ListEnterpriseInfoResponseBodyData) SetEnterpriseName(v string) *ListEnterpriseInfoResponseBodyData {
	s.EnterpriseName = &v
	return s
}

func (s *ListEnterpriseInfoResponseBodyData) SetManagerContactNumber(v string) *ListEnterpriseInfoResponseBodyData {
	s.ManagerContactNumber = &v
	return s
}

func (s *ListEnterpriseInfoResponseBodyData) SetManagerName(v string) *ListEnterpriseInfoResponseBodyData {
	s.ManagerName = &v
	return s
}

func (s *ListEnterpriseInfoResponseBodyData) SetOrganizationCode(v string) *ListEnterpriseInfoResponseBodyData {
	s.OrganizationCode = &v
	return s
}

func (s *ListEnterpriseInfoResponseBodyData) SetStatus(v int32) *ListEnterpriseInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEnterpriseInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
