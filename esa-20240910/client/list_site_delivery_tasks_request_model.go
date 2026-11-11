// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteDeliveryTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *ListSiteDeliveryTasksRequest
	GetBusinessType() *string
	SetPageNumber(v int64) *ListSiteDeliveryTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSiteDeliveryTasksRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListSiteDeliveryTasksRequest
	GetSiteId() *int64
}

type ListSiteDeliveryTasksRequest struct {
	// The real-time log type. Valid values:
	//
	// - **dcdn_log_access_l1*	- (default): Access logs.
	//
	// - **dcdn_log_er**: Edge Routine logs.
	//
	// - **dcdn_log_waf**: Security protection logs.
	//
	// - **dcdn_log_ipa**: Layer 4 acceleration logs.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListSiteDeliveryTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSiteDeliveryTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListSiteDeliveryTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSiteDeliveryTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSiteDeliveryTasksRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListSiteDeliveryTasksRequest) SetBusinessType(v string) *ListSiteDeliveryTasksRequest {
	s.BusinessType = &v
	return s
}

func (s *ListSiteDeliveryTasksRequest) SetPageNumber(v int64) *ListSiteDeliveryTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSiteDeliveryTasksRequest) SetPageSize(v int64) *ListSiteDeliveryTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSiteDeliveryTasksRequest) SetSiteId(v int64) *ListSiteDeliveryTasksRequest {
	s.SiteId = &v
	return s
}

func (s *ListSiteDeliveryTasksRequest) Validate() error {
	return dara.Validate(s)
}
