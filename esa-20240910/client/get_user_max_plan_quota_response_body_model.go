// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMaxPlanQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetUserMaxPlanQuotaResponseBody
	GetInstanceId() *string
	SetPlanName(v string) *GetUserMaxPlanQuotaResponseBody
	GetPlanName() *string
	SetQuotaValue(v string) *GetUserMaxPlanQuotaResponseBody
	GetQuotaValue() *string
	SetRequestId(v string) *GetUserMaxPlanQuotaResponseBody
	GetRequestId() *string
}

type GetUserMaxPlanQuotaResponseBody struct {
	// The plan instance ID. You can obtain this value by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// esa-site-b09z4sk9pbls
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The plan name.
	//
	// <props="china">
	//
	// - Free Edition: entranceplan
	//
	// - Basic: basic
	//
	// - Standard: medium
	//
	// - Premium Edition: high
	//
	//
	// <props="intl">
	//
	// - Entrance: entranceplan_intl
	//
	// - Pro: basicplan_intl
	//
	// - Premium: vipplan_intl
	//
	// example:
	//
	// entranceplan
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The quota value.
	//
	// example:
	//
	// 44640
	QuotaValue *string `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6599BB6-525D-5CFF-86BC-24068E6FB3EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserMaxPlanQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserMaxPlanQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMaxPlanQuotaResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserMaxPlanQuotaResponseBody) GetPlanName() *string {
	return s.PlanName
}

func (s *GetUserMaxPlanQuotaResponseBody) GetQuotaValue() *string {
	return s.QuotaValue
}

func (s *GetUserMaxPlanQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserMaxPlanQuotaResponseBody) SetInstanceId(v string) *GetUserMaxPlanQuotaResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetUserMaxPlanQuotaResponseBody) SetPlanName(v string) *GetUserMaxPlanQuotaResponseBody {
	s.PlanName = &v
	return s
}

func (s *GetUserMaxPlanQuotaResponseBody) SetQuotaValue(v string) *GetUserMaxPlanQuotaResponseBody {
	s.QuotaValue = &v
	return s
}

func (s *GetUserMaxPlanQuotaResponseBody) SetRequestId(v string) *GetUserMaxPlanQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserMaxPlanQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
