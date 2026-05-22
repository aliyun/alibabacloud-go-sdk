// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceQuotasRequest
	GetInstanceId() *string
	SetQuotaNames(v string) *ListInstanceQuotasRequest
	GetQuotaNames() *string
	SetSiteId(v int64) *ListInstanceQuotasRequest
	GetSiteId() *int64
}

type ListInstanceQuotasRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	QuotaNames *string `json:"QuotaNames,omitempty" xml:"QuotaNames,omitempty"`
	// if can be null:
	// false
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListInstanceQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceQuotasRequest) GetQuotaNames() *string {
	return s.QuotaNames
}

func (s *ListInstanceQuotasRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListInstanceQuotasRequest) SetInstanceId(v string) *ListInstanceQuotasRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceQuotasRequest) SetQuotaNames(v string) *ListInstanceQuotasRequest {
	s.QuotaNames = &v
	return s
}

func (s *ListInstanceQuotasRequest) SetSiteId(v int64) *ListInstanceQuotasRequest {
	s.SiteId = &v
	return s
}

func (s *ListInstanceQuotasRequest) Validate() error {
	return dara.Validate(s)
}
