// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteUrgentDemandPlanRequest
	GetId() *int64
	SetModifier(v string) *DeleteUrgentDemandPlanRequest
	GetModifier() *string
}

type DeleteUrgentDemandPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 222111
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s DeleteUrgentDemandPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteUrgentDemandPlanRequest) GetModifier() *string {
	return s.Modifier
}

func (s *DeleteUrgentDemandPlanRequest) SetId(v int64) *DeleteUrgentDemandPlanRequest {
	s.Id = &v
	return s
}

func (s *DeleteUrgentDemandPlanRequest) SetModifier(v string) *DeleteUrgentDemandPlanRequest {
	s.Modifier = &v
	return s
}

func (s *DeleteUrgentDemandPlanRequest) Validate() error {
	return dara.Validate(s)
}
