// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiTemplatesDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *ListApiTemplatesDTO
	GetApiName() *string
	SetContent(v string) *ListApiTemplatesDTO
	GetContent() *string
	SetGmtCreate(v int64) *ListApiTemplatesDTO
	GetGmtCreate() *int64
	SetGmtModified(v int64) *ListApiTemplatesDTO
	GetGmtModified() *int64
	SetId(v int64) *ListApiTemplatesDTO
	GetId() *int64
	SetOperatorId(v string) *ListApiTemplatesDTO
	GetOperatorId() *string
	SetRegionId(v string) *ListApiTemplatesDTO
	GetRegionId() *string
	SetResourceGroupId(v string) *ListApiTemplatesDTO
	GetResourceGroupId() *string
	SetStatus(v string) *ListApiTemplatesDTO
	GetStatus() *string
	SetTemplateId(v string) *ListApiTemplatesDTO
	GetTemplateId() *string
	SetTemplateName(v string) *ListApiTemplatesDTO
	GetTemplateName() *string
	SetUserId(v string) *ListApiTemplatesDTO
	GetUserId() *string
}

type ListApiTemplatesDTO struct {
	ApiName         *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GmtCreate       *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OperatorId      *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListApiTemplatesDTO) String() string {
	return dara.Prettify(s)
}

func (s ListApiTemplatesDTO) GoString() string {
	return s.String()
}

func (s *ListApiTemplatesDTO) GetApiName() *string {
	return s.ApiName
}

func (s *ListApiTemplatesDTO) GetContent() *string {
	return s.Content
}

func (s *ListApiTemplatesDTO) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListApiTemplatesDTO) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListApiTemplatesDTO) GetId() *int64 {
	return s.Id
}

func (s *ListApiTemplatesDTO) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListApiTemplatesDTO) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApiTemplatesDTO) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApiTemplatesDTO) GetStatus() *string {
	return s.Status
}

func (s *ListApiTemplatesDTO) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListApiTemplatesDTO) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListApiTemplatesDTO) GetUserId() *string {
	return s.UserId
}

func (s *ListApiTemplatesDTO) SetApiName(v string) *ListApiTemplatesDTO {
	s.ApiName = &v
	return s
}

func (s *ListApiTemplatesDTO) SetContent(v string) *ListApiTemplatesDTO {
	s.Content = &v
	return s
}

func (s *ListApiTemplatesDTO) SetGmtCreate(v int64) *ListApiTemplatesDTO {
	s.GmtCreate = &v
	return s
}

func (s *ListApiTemplatesDTO) SetGmtModified(v int64) *ListApiTemplatesDTO {
	s.GmtModified = &v
	return s
}

func (s *ListApiTemplatesDTO) SetId(v int64) *ListApiTemplatesDTO {
	s.Id = &v
	return s
}

func (s *ListApiTemplatesDTO) SetOperatorId(v string) *ListApiTemplatesDTO {
	s.OperatorId = &v
	return s
}

func (s *ListApiTemplatesDTO) SetRegionId(v string) *ListApiTemplatesDTO {
	s.RegionId = &v
	return s
}

func (s *ListApiTemplatesDTO) SetResourceGroupId(v string) *ListApiTemplatesDTO {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApiTemplatesDTO) SetStatus(v string) *ListApiTemplatesDTO {
	s.Status = &v
	return s
}

func (s *ListApiTemplatesDTO) SetTemplateId(v string) *ListApiTemplatesDTO {
	s.TemplateId = &v
	return s
}

func (s *ListApiTemplatesDTO) SetTemplateName(v string) *ListApiTemplatesDTO {
	s.TemplateName = &v
	return s
}

func (s *ListApiTemplatesDTO) SetUserId(v string) *ListApiTemplatesDTO {
	s.UserId = &v
	return s
}

func (s *ListApiTemplatesDTO) Validate() error {
	return dara.Validate(s)
}
