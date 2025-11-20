// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDeliveryPlanHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateDeliveryPlanHeadersAccountContext) *CreateDeliveryPlanHeaders
	GetAccountContext() *CreateDeliveryPlanHeadersAccountContext
}

type CreateDeliveryPlanHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateDeliveryPlanHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateDeliveryPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDeliveryPlanHeaders) GetAccountContext() *CreateDeliveryPlanHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateDeliveryPlanHeaders) SetCommonHeaders(v map[string]*string) *CreateDeliveryPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeliveryPlanHeaders) SetAccountContext(v *CreateDeliveryPlanHeadersAccountContext) *CreateDeliveryPlanHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateDeliveryPlanHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDeliveryPlanHeadersAccountContext struct {
	// example:
	//
	// ba3a9b612345678d8fedf544ef69d19e
	UserToken *string `json:"userToken,omitempty" xml:"userToken,omitempty"`
}

func (s CreateDeliveryPlanHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanHeadersAccountContext) GetUserToken() *string {
	return s.UserToken
}

func (s *CreateDeliveryPlanHeadersAccountContext) SetUserToken(v string) *CreateDeliveryPlanHeadersAccountContext {
	s.UserToken = &v
	return s
}

func (s *CreateDeliveryPlanHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
