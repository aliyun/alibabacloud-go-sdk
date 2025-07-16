// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFCTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTriggerARN(v string) *DeleteFCTriggerRequest
	GetTriggerARN() *string
}

type DeleteFCTriggerRequest struct {
	// The trigger that corresponds to the Function Compute service.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:cdn:{RegionID}:{AccountID}:{Filter}
	TriggerARN *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
}

func (s DeleteFCTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteFCTriggerRequest) GetTriggerARN() *string {
	return s.TriggerARN
}

func (s *DeleteFCTriggerRequest) SetTriggerARN(v string) *DeleteFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

func (s *DeleteFCTriggerRequest) Validate() error {
	return dara.Validate(s)
}
