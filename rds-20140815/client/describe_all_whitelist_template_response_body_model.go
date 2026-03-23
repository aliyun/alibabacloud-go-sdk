// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAllWhitelistTemplateResponseBody
	GetCode() *string
	SetData(v *DescribeAllWhitelistTemplateResponseBodyData) *DescribeAllWhitelistTemplateResponseBody
	GetData() *DescribeAllWhitelistTemplateResponseBodyData
	SetHttpStatusCode(v int32) *DescribeAllWhitelistTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeAllWhitelistTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAllWhitelistTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAllWhitelistTemplateResponseBody
	GetSuccess() *bool
}

type DescribeAllWhitelistTemplateResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeAllWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAllWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllWhitelistTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAllWhitelistTemplateResponseBody) GetData() *DescribeAllWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *DescribeAllWhitelistTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAllWhitelistTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAllWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllWhitelistTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAllWhitelistTemplateResponseBody) SetCode(v string) *DescribeAllWhitelistTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBody) SetData(v *DescribeAllWhitelistTemplateResponseBodyData) *DescribeAllWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBody) SetHttpStatusCode(v int32) *DescribeAllWhitelistTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBody) SetMessage(v string) *DescribeAllWhitelistTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBody) SetRequestId(v string) *DescribeAllWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBody) SetSuccess(v bool) *DescribeAllWhitelistTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAllWhitelistTemplateResponseBodyData struct {
	CurrPageNumbers   *int32                                                   `json:"CurrPageNumbers,omitempty" xml:"CurrPageNumbers,omitempty"`
	HasNext           *bool                                                    `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	HasPrev           *bool                                                    `json:"HasPrev,omitempty" xml:"HasPrev,omitempty"`
	MaxRecordsPerPage *int32                                                   `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	Templates         []*DescribeAllWhitelistTemplateResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	TotalPageNumbers  *int32                                                   `json:"TotalPageNumbers,omitempty" xml:"TotalPageNumbers,omitempty"`
	TotalRecords      *int32                                                   `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeAllWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) GetCurrPageNumbers() *int32 {
	return s.CurrPageNumbers
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) GetHasNext() *bool {
	return s.HasNext
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) GetHasPrev() *bool {
	return s.HasPrev
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) GetMaxRecordsPerPage() *int32 {
	return s.MaxRecordsPerPage
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) GetTemplates() []*DescribeAllWhitelistTemplateResponseBodyDataTemplates {
	return s.Templates
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) GetTotalPageNumbers() *int32 {
	return s.TotalPageNumbers
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) GetTotalRecords() *int32 {
	return s.TotalRecords
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) SetCurrPageNumbers(v int32) *DescribeAllWhitelistTemplateResponseBodyData {
	s.CurrPageNumbers = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) SetHasNext(v bool) *DescribeAllWhitelistTemplateResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) SetHasPrev(v bool) *DescribeAllWhitelistTemplateResponseBodyData {
	s.HasPrev = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) SetMaxRecordsPerPage(v int32) *DescribeAllWhitelistTemplateResponseBodyData {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) SetTemplates(v []*DescribeAllWhitelistTemplateResponseBodyDataTemplates) *DescribeAllWhitelistTemplateResponseBodyData {
	s.Templates = v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) SetTotalPageNumbers(v int32) *DescribeAllWhitelistTemplateResponseBodyData {
	s.TotalPageNumbers = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) SetTotalRecords(v int32) *DescribeAllWhitelistTemplateResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyData) Validate() error {
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

type DescribeAllWhitelistTemplateResponseBodyDataTemplates struct {
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Ips          *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	TemplateId   *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UserId       *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeAllWhitelistTemplateResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllWhitelistTemplateResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) GetId() *int32 {
	return s.Id
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) GetIps() *string {
	return s.Ips
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) GetUserId() *int32 {
	return s.UserId
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) SetId(v int32) *DescribeAllWhitelistTemplateResponseBodyDataTemplates {
	s.Id = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) SetIps(v string) *DescribeAllWhitelistTemplateResponseBodyDataTemplates {
	s.Ips = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) SetTemplateId(v int32) *DescribeAllWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) SetTemplateName(v string) *DescribeAllWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateName = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) SetUserId(v int32) *DescribeAllWhitelistTemplateResponseBodyDataTemplates {
	s.UserId = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponseBodyDataTemplates) Validate() error {
	return dara.Validate(s)
}
