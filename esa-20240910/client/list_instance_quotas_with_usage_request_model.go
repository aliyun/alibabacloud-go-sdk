// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceQuotasWithUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceQuotasWithUsageRequest
	GetInstanceId() *string
	SetQuotaNames(v string) *ListInstanceQuotasWithUsageRequest
	GetQuotaNames() *string
	SetSiteId(v int64) *ListInstanceQuotasWithUsageRequest
	GetSiteId() *int64
}

type ListInstanceQuotasWithUsageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	QuotaNames *string `json:"QuotaNames,omitempty" xml:"QuotaNames,omitempty"`
	// if can be null:
	// false
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListInstanceQuotasWithUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasWithUsageRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasWithUsageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceQuotasWithUsageRequest) GetQuotaNames() *string {
	return s.QuotaNames
}

func (s *ListInstanceQuotasWithUsageRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListInstanceQuotasWithUsageRequest) SetInstanceId(v string) *ListInstanceQuotasWithUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceQuotasWithUsageRequest) SetQuotaNames(v string) *ListInstanceQuotasWithUsageRequest {
	s.QuotaNames = &v
	return s
}

func (s *ListInstanceQuotasWithUsageRequest) SetSiteId(v int64) *ListInstanceQuotasWithUsageRequest {
	s.SiteId = &v
	return s
}

func (s *ListInstanceQuotasWithUsageRequest) Validate() error {
	return dara.Validate(s)
}
