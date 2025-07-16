// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFCTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTriggerARN(v string) *DescribeFCTriggerRequest
	GetTriggerARN() *string
}

type DescribeFCTriggerRequest struct {
	// The trigger that corresponds to the Function Compute service.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:cdn:{RegionID}:{AccountID}:{Filter}
	TriggerARN *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
}

func (s DescribeFCTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerRequest) GetTriggerARN() *string {
	return s.TriggerARN
}

func (s *DescribeFCTriggerRequest) SetTriggerARN(v string) *DescribeFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

func (s *DescribeFCTriggerRequest) Validate() error {
	return dara.Validate(s)
}
