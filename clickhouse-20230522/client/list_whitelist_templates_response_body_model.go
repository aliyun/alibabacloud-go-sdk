// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhitelistTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListWhitelistTemplatesResponseBodyData) *ListWhitelistTemplatesResponseBody
	GetData() *ListWhitelistTemplatesResponseBodyData
	SetRequestId(v string) *ListWhitelistTemplatesResponseBody
	GetRequestId() *string
}

type ListWhitelistTemplatesResponseBody struct {
	Data *ListWhitelistTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWhitelistTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWhitelistTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWhitelistTemplatesResponseBody) GetData() *ListWhitelistTemplatesResponseBodyData {
	return s.Data
}

func (s *ListWhitelistTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWhitelistTemplatesResponseBody) SetData(v *ListWhitelistTemplatesResponseBodyData) *ListWhitelistTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListWhitelistTemplatesResponseBody) SetRequestId(v string) *ListWhitelistTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWhitelistTemplatesResponseBodyData struct {
	// example:
	//
	// 1
	CurrPageNumbers *int32 `json:"CurrPageNumbers,omitempty" xml:"CurrPageNumbers,omitempty"`
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// false
	HasPrev *bool `json:"HasPrev,omitempty" xml:"HasPrev,omitempty"`
	// example:
	//
	// 20
	PageSize  *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Templates []*ListWhitelistTemplatesResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPageNumbers *int32 `json:"TotalPageNumbers,omitempty" xml:"TotalPageNumbers,omitempty"`
}

func (s ListWhitelistTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWhitelistTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWhitelistTemplatesResponseBodyData) GetCurrPageNumbers() *int32 {
	return s.CurrPageNumbers
}

func (s *ListWhitelistTemplatesResponseBodyData) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListWhitelistTemplatesResponseBodyData) GetHasPrev() *bool {
	return s.HasPrev
}

func (s *ListWhitelistTemplatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWhitelistTemplatesResponseBodyData) GetTemplates() []*ListWhitelistTemplatesResponseBodyDataTemplates {
	return s.Templates
}

func (s *ListWhitelistTemplatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWhitelistTemplatesResponseBodyData) GetTotalPageNumbers() *int32 {
	return s.TotalPageNumbers
}

func (s *ListWhitelistTemplatesResponseBodyData) SetCurrPageNumbers(v int32) *ListWhitelistTemplatesResponseBodyData {
	s.CurrPageNumbers = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyData) SetHasNext(v bool) *ListWhitelistTemplatesResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyData) SetHasPrev(v bool) *ListWhitelistTemplatesResponseBodyData {
	s.HasPrev = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyData) SetPageSize(v int32) *ListWhitelistTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyData) SetTemplates(v []*ListWhitelistTemplatesResponseBodyDataTemplates) *ListWhitelistTemplatesResponseBodyData {
	s.Templates = v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyData) SetTotalCount(v int32) *ListWhitelistTemplatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyData) SetTotalPageNumbers(v int32) *ListWhitelistTemplatesResponseBodyData {
	s.TotalPageNumbers = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyData) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWhitelistTemplatesResponseBodyDataTemplates struct {
	DbInstances []*ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 192.168.1.1,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// g-asdfwem
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// g-asdfwem
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListWhitelistTemplatesResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListWhitelistTemplatesResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) GetDbInstances() []*ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances {
	return s.DbInstances
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) SetDbInstances(v []*ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances) *ListWhitelistTemplatesResponseBodyDataTemplates {
	s.DbInstances = v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) SetSecurityIPList(v string) *ListWhitelistTemplatesResponseBodyDataTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) SetTemplateId(v string) *ListWhitelistTemplatesResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) SetTemplateName(v string) *ListWhitelistTemplatesResponseBodyDataTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplates) Validate() error {
	if s.DbInstances != nil {
		for _, item := range s.DbInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances struct {
	// example:
	//
	// cc-asdfwef
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances) String() string {
	return dara.Prettify(s)
}

func (s ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances) GoString() string {
	return s.String()
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances) SetDBInstanceId(v string) *ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances {
	s.DBInstanceId = &v
	return s
}

func (s *ListWhitelistTemplatesResponseBodyDataTemplatesDbInstances) Validate() error {
	return dara.Validate(s)
}
