// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarSecurityEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSimilarSecurityEventsRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeSimilarSecurityEventsRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeSimilarSecurityEventsRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeSimilarSecurityEventsRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeSimilarSecurityEventsRequest
	GetSourceIp() *string
	SetTaskId(v int64) *DescribeSimilarSecurityEventsRequest
	GetTaskId() *int64
}

type DescribeSimilarSecurityEventsRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The task ID. You can call the [CreateSimilarSecurityEventsQueryTask](~~CreateSimilarSecurityEventsQueryTask~~) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1689135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSimilarSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSimilarSecurityEventsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSimilarSecurityEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSimilarSecurityEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSimilarSecurityEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSimilarSecurityEventsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeSimilarSecurityEventsRequest) SetCurrentPage(v int32) *DescribeSimilarSecurityEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetLang(v string) *DescribeSimilarSecurityEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetPageSize(v int32) *DescribeSimilarSecurityEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetResourceOwnerId(v int64) *DescribeSimilarSecurityEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetSourceIp(v string) *DescribeSimilarSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetTaskId(v int64) *DescribeSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}
