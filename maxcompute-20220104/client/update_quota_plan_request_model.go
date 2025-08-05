// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateQuotaPlanRequest
	GetBody() *string
	SetRegion(v string) *UpdateQuotaPlanRequest
	GetRegion() *string
	SetTenantId(v string) *UpdateQuotaPlanRequest
	GetTenantId() *string
}

type UpdateQuotaPlanRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// { "name": "planA", // The quota is a level-1 quota. You can select only the fields that are related to the quota plan. "quota": { "name": "a", "nickName": "aaa_nick", "tenantId": "10001", "regionId": "cn-hangzhou", "parentId": "0", "cluster": "AT-ODPS-TEST3", "parameter": { "minCU": 40, "maxCU": 40, "adhocCU": 0, "elasticMinCU": 40, "elasticMaxCU": 40, "enablePreemptiveScheduling": false, "forceReservedMin":true, "enablePriority":false, "singleJobCULimit":100, "adhocQuotaBeginTimeInSec": 1345, "adhocQuotaEndTimeInSec": 1234, "ignoreAdhocQuota":false }, "subQuotaInfoList": [ { "nickName": "WlmFuxiSecondaryOnlineQuotaTest", "name": "WlmFuxiSecondaryOnlineQuotaTest", "type": "FUXI_ONLINE", "tenantId": "10001", "regionId": "cn-hangzhou", "cluster": "AT-ODPS-TEST3", "parameter": { "minCU": 40, "maxCU": 40, "adhocCU": 0, "elasticMinCU": 40, "elasticMaxCU": 40, "enablePreemptiveScheduling": false, "forceReservedMin":true, "enablePriority":false, "singleJobCULimit":100, "adhocQuotaBeginTimeInSec": 1345, "adhocQuotaEndTimeInSec": 1234, "ignoreAdhocQuota":false } } ] } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateQuotaPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateQuotaPlanRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateQuotaPlanRequest) SetBody(v string) *UpdateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetRegion(v string) *UpdateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetTenantId(v string) *UpdateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateQuotaPlanRequest) Validate() error {
	return dara.Validate(s)
}
