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
	// The plan ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The quota names in the plan. Separate the quota names with commas (,). You can query up to 10 quota names at a time. Valid values:
	//
	// 	- **customHttpCert**: the custom certificates.
	//
	// 	- **transition_rule**: the transform rules.
	//
	// 	- **waiting_room**: the waiting rooms.
	//
	// 	- **https|rule_quota**: the SSL/TLS rules.
	//
	// 	- **cache_rules|rule_quota**: the cache rules.
	//
	// 	- **configuration_rules|rule_quota**: the configuration rules.
	//
	// 	- **redirect_rules|rule_quota**: the redirect rules.
	//
	// 	- **compression_rules|rule_quota**: the compression rules.
	//
	// 	- **origin_rules|rule_quota**: the origin rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// customHttpCert
	QuotaNames *string `json:"QuotaNames,omitempty" xml:"QuotaNames,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1232223****
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
