// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryPlanShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDeliveryPlanShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateDeliveryPlanShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateDeliveryPlanShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateDeliveryPlanShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDeliveryPlanShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateDeliveryPlanShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateDeliveryPlanShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeliveryPlanShrinkHeaders) SetAccountContextShrink(v string) *CreateDeliveryPlanShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateDeliveryPlanShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
