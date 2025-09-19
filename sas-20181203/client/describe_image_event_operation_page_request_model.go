// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageEventOperationPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageEventOperationPageRequest
	GetCurrentPage() *int32
	SetEventKey(v string) *DescribeImageEventOperationPageRequest
	GetEventKey() *string
	SetEventName(v string) *DescribeImageEventOperationPageRequest
	GetEventName() *string
	SetEventType(v string) *DescribeImageEventOperationPageRequest
	GetEventType() *string
	SetId(v int64) *DescribeImageEventOperationPageRequest
	GetId() *int64
	SetLang(v string) *DescribeImageEventOperationPageRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageEventOperationPageRequest
	GetPageSize() *int32
	SetSource(v string) *DescribeImageEventOperationPageRequest
	GetSource() *string
}

type DescribeImageEventOperationPageRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword of the alert item.
	//
	// example:
	//
	// PEM
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the alert item.
	//
	// example:
	//
	// PEM
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The alert type.
	//
	// 	- Set the value to **sensitiveFile**.
	//
	// example:
	//
	// sensitiveFile
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the alert handling rule.
	//
	// example:
	//
	// 49616
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source of the alert handling rule. Valid values:
	//
	// 	- **default**: image.
	//
	// 	- **agentless**: agentless detection.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeImageEventOperationPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageEventOperationPageRequest) GetEventKey() *string {
	return s.EventKey
}

func (s *DescribeImageEventOperationPageRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeImageEventOperationPageRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeImageEventOperationPageRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeImageEventOperationPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageEventOperationPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageEventOperationPageRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeImageEventOperationPageRequest) SetCurrentPage(v int32) *DescribeImageEventOperationPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) SetEventKey(v string) *DescribeImageEventOperationPageRequest {
	s.EventKey = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) SetEventName(v string) *DescribeImageEventOperationPageRequest {
	s.EventName = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) SetEventType(v string) *DescribeImageEventOperationPageRequest {
	s.EventType = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) SetId(v int64) *DescribeImageEventOperationPageRequest {
	s.Id = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) SetLang(v string) *DescribeImageEventOperationPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) SetPageSize(v int32) *DescribeImageEventOperationPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) SetSource(v string) *DescribeImageEventOperationPageRequest {
	s.Source = &v
	return s
}

func (s *DescribeImageEventOperationPageRequest) Validate() error {
	return dara.Validate(s)
}
