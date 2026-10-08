// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListScheduleTemplatesCommand(v *ListScheduleTemplatesRequestListScheduleTemplatesCommand) *ListScheduleTemplatesRequest
	GetListScheduleTemplatesCommand() *ListScheduleTemplatesRequestListScheduleTemplatesCommand
	SetOpTenantId(v int64) *ListScheduleTemplatesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListScheduleTemplatesRequest
	GetOpUserId() *string
}

type ListScheduleTemplatesRequest struct {
	// This parameter is required.
	ListScheduleTemplatesCommand *ListScheduleTemplatesRequestListScheduleTemplatesCommand `json:"ListScheduleTemplatesCommand,omitempty" xml:"ListScheduleTemplatesCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListScheduleTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesRequest) GetListScheduleTemplatesCommand() *ListScheduleTemplatesRequestListScheduleTemplatesCommand {
	return s.ListScheduleTemplatesCommand
}

func (s *ListScheduleTemplatesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListScheduleTemplatesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListScheduleTemplatesRequest) SetListScheduleTemplatesCommand(v *ListScheduleTemplatesRequestListScheduleTemplatesCommand) *ListScheduleTemplatesRequest {
	s.ListScheduleTemplatesCommand = v
	return s
}

func (s *ListScheduleTemplatesRequest) SetOpTenantId(v int64) *ListScheduleTemplatesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListScheduleTemplatesRequest) SetOpUserId(v string) *ListScheduleTemplatesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListScheduleTemplatesRequest) Validate() error {
	if s.ListScheduleTemplatesCommand != nil {
		if err := s.ListScheduleTemplatesCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleTemplatesRequestListScheduleTemplatesCommand struct {
	// example:
	//
	// 小时
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BASE_SCHEDULE_TEMPLATE
	ScheduleTemplateType *string `json:"ScheduleTemplateType,omitempty" xml:"ScheduleTemplateType,omitempty"`
}

func (s ListScheduleTemplatesRequestListScheduleTemplatesCommand) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesRequestListScheduleTemplatesCommand) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) GetKeyword() *string {
	return s.Keyword
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) GetScheduleTemplateType() *string {
	return s.ScheduleTemplateType
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) SetKeyword(v string) *ListScheduleTemplatesRequestListScheduleTemplatesCommand {
	s.Keyword = &v
	return s
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) SetPageNumber(v int32) *ListScheduleTemplatesRequestListScheduleTemplatesCommand {
	s.PageNumber = &v
	return s
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) SetPageSize(v int32) *ListScheduleTemplatesRequestListScheduleTemplatesCommand {
	s.PageSize = &v
	return s
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) SetScheduleTemplateType(v string) *ListScheduleTemplatesRequestListScheduleTemplatesCommand {
	s.ScheduleTemplateType = &v
	return s
}

func (s *ListScheduleTemplatesRequestListScheduleTemplatesCommand) Validate() error {
	return dara.Validate(s)
}
