// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDeliveryTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *ListUserDeliveryTasksRequest
	GetBusinessType() *string
	SetPageNumber(v int64) *ListUserDeliveryTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListUserDeliveryTasksRequest
	GetPageSize() *int64
}

type ListUserDeliveryTasksRequest struct {
	// The log category. Valid values:
	//
	// 1.  dcdn_log_access_l1 (default): access logs.
	//
	// 2.  dcdn_log_er: Edge Routine logs.
	//
	// 3.  dcdn_log_waf: firewall logs.
	//
	// 4.  dcdn_log_ipa: TCP/UDP proxy logs.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserDeliveryTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserDeliveryTasksRequest) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListUserDeliveryTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListUserDeliveryTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserDeliveryTasksRequest) SetBusinessType(v string) *ListUserDeliveryTasksRequest {
	s.BusinessType = &v
	return s
}

func (s *ListUserDeliveryTasksRequest) SetPageNumber(v int64) *ListUserDeliveryTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserDeliveryTasksRequest) SetPageSize(v int64) *ListUserDeliveryTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserDeliveryTasksRequest) Validate() error {
	return dara.Validate(s)
}
