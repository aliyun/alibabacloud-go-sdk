// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *GetInstanceSummaryRequest
	GetInstanceType() *string
}

type GetInstanceSummaryRequest struct {
	// example:
	//
	// TEST
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetInstanceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceSummaryRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetInstanceSummaryRequest) SetInstanceType(v string) *GetInstanceSummaryRequest {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
