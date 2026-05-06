// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *DisableSubscriptionRequest
	GetBusinessUnitId() *string
}

type DisableSubscriptionRequest struct {
	// example:
	//
	// llm-3pptowd2olrctsvc
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
}

func (s DisableSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *DisableSubscriptionRequest) SetBusinessUnitId(v string) *DisableSubscriptionRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *DisableSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
