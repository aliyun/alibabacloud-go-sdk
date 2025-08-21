// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActionPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanId(v string) *GetActionPlanRequest
	GetActionPlanId() *string
}

type GetActionPlanRequest struct {
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
}

func (s GetActionPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetActionPlanRequest) GoString() string {
	return s.String()
}

func (s *GetActionPlanRequest) GetActionPlanId() *string {
	return s.ActionPlanId
}

func (s *GetActionPlanRequest) SetActionPlanId(v string) *GetActionPlanRequest {
	s.ActionPlanId = &v
	return s
}

func (s *GetActionPlanRequest) Validate() error {
	return dara.Validate(s)
}
