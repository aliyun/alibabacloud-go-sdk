// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleNumLimitOfSLARequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *GetRuleNumLimitOfSLARequest
	GetDagId() *int64
	SetTid(v int64) *GetRuleNumLimitOfSLARequest
	GetTid() *int64
}

type GetRuleNumLimitOfSLARequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetRuleNumLimitOfSLARequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuleNumLimitOfSLARequest) GoString() string {
	return s.String()
}

func (s *GetRuleNumLimitOfSLARequest) GetDagId() *int64 {
	return s.DagId
}

func (s *GetRuleNumLimitOfSLARequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetRuleNumLimitOfSLARequest) SetDagId(v int64) *GetRuleNumLimitOfSLARequest {
	s.DagId = &v
	return s
}

func (s *GetRuleNumLimitOfSLARequest) SetTid(v int64) *GetRuleNumLimitOfSLARequest {
	s.Tid = &v
	return s
}

func (s *GetRuleNumLimitOfSLARequest) Validate() error {
	return dara.Validate(s)
}
