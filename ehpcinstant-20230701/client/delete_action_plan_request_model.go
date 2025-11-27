// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteActionPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanId(v string) *DeleteActionPlanRequest
	GetActionPlanId() *string
}

type DeleteActionPlanRequest struct {
	// The ID of the execution plan.
	//
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
}

func (s DeleteActionPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteActionPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteActionPlanRequest) GetActionPlanId() *string {
	return s.ActionPlanId
}

func (s *DeleteActionPlanRequest) SetActionPlanId(v string) *DeleteActionPlanRequest {
	s.ActionPlanId = &v
	return s
}

func (s *DeleteActionPlanRequest) Validate() error {
	return dara.Validate(s)
}
