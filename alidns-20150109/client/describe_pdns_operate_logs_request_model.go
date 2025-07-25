// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsOperateLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *DescribePdnsOperateLogsRequest
	GetActionType() *string
	SetEndDate(v string) *DescribePdnsOperateLogsRequest
	GetEndDate() *string
	SetKeyword(v string) *DescribePdnsOperateLogsRequest
	GetKeyword() *string
	SetLang(v string) *DescribePdnsOperateLogsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribePdnsOperateLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsOperateLogsRequest
	GetPageSize() *int64
	SetResourceType(v string) *DescribePdnsOperateLogsRequest
	GetResourceType() *string
	SetStartDate(v string) *DescribePdnsOperateLogsRequest
	GetStartDate() *string
}

type DescribePdnsOperateLogsRequest struct {
	ActionType   *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribePdnsOperateLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsOperateLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsOperateLogsRequest) GetActionType() *string {
	return s.ActionType
}

func (s *DescribePdnsOperateLogsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePdnsOperateLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribePdnsOperateLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsOperateLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsOperateLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsOperateLogsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePdnsOperateLogsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePdnsOperateLogsRequest) SetActionType(v string) *DescribePdnsOperateLogsRequest {
	s.ActionType = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) SetEndDate(v string) *DescribePdnsOperateLogsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) SetKeyword(v string) *DescribePdnsOperateLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) SetLang(v string) *DescribePdnsOperateLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) SetPageNumber(v int64) *DescribePdnsOperateLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) SetPageSize(v int64) *DescribePdnsOperateLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) SetResourceType(v string) *DescribePdnsOperateLogsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) SetStartDate(v string) *DescribePdnsOperateLogsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePdnsOperateLogsRequest) Validate() error {
	return dara.Validate(s)
}
