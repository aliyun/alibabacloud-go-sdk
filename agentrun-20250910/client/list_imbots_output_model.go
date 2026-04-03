// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIMBotsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*IMBotInfo) *ListIMBotsOutput
	GetItems() []*IMBotInfo
	SetPageNumber(v int32) *ListIMBotsOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIMBotsOutput
	GetPageSize() *int32
	SetTotal(v int32) *ListIMBotsOutput
	GetTotal() *int32
}

type ListIMBotsOutput struct {
	Items      []*IMBotInfo `json:"items" xml:"items" type:"Repeated"`
	PageNumber *int32       `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int32       `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListIMBotsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListIMBotsOutput) GoString() string {
	return s.String()
}

func (s *ListIMBotsOutput) GetItems() []*IMBotInfo {
	return s.Items
}

func (s *ListIMBotsOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIMBotsOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIMBotsOutput) GetTotal() *int32 {
	return s.Total
}

func (s *ListIMBotsOutput) SetItems(v []*IMBotInfo) *ListIMBotsOutput {
	s.Items = v
	return s
}

func (s *ListIMBotsOutput) SetPageNumber(v int32) *ListIMBotsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListIMBotsOutput) SetPageSize(v int32) *ListIMBotsOutput {
	s.PageSize = &v
	return s
}

func (s *ListIMBotsOutput) SetTotal(v int32) *ListIMBotsOutput {
	s.Total = &v
	return s
}

func (s *ListIMBotsOutput) Validate() error {
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
