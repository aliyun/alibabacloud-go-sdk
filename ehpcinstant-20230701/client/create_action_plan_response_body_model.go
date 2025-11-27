// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActionPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanId(v string) *CreateActionPlanResponseBody
	GetActionPlanId() *string
	SetRequestId(v string) *CreateActionPlanResponseBody
	GetRequestId() *string
}

type CreateActionPlanResponseBody struct {
	// The ID of the execution plan.
	//
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateActionPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateActionPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateActionPlanResponseBody) GetActionPlanId() *string {
	return s.ActionPlanId
}

func (s *CreateActionPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateActionPlanResponseBody) SetActionPlanId(v string) *CreateActionPlanResponseBody {
	s.ActionPlanId = &v
	return s
}

func (s *CreateActionPlanResponseBody) SetRequestId(v string) *CreateActionPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateActionPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
