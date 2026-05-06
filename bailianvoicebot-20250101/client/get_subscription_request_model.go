// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *GetSubscriptionRequest
	GetBusinessUnitId() *string
}

type GetSubscriptionRequest struct {
	// example:
	//
	// llm-3pptowd2olrctsvc
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
}

func (s GetSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *GetSubscriptionRequest) SetBusinessUnitId(v string) *GetSubscriptionRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *GetSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
