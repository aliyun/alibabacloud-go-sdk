// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateQuotaPlanRequest
	GetBody() *string
	SetRegion(v string) *CreateQuotaPlanRequest
	GetRegion() *string
	SetTenantId(v string) *CreateQuotaPlanRequest
	GetTenantId() *string
}

type CreateQuotaPlanRequest struct {
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
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 228451885265153
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanRequest) GetBody() *string {
	return s.Body
}

func (s *CreateQuotaPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateQuotaPlanRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateQuotaPlanRequest) SetBody(v string) *CreateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetRegion(v string) *CreateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetTenantId(v string) *CreateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

func (s *CreateQuotaPlanRequest) Validate() error {
	return dara.Validate(s)
}
