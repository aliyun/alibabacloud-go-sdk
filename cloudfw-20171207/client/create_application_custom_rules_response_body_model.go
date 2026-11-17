// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationCustomRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCustomRuleIds(v []*string) *CreateApplicationCustomRulesResponseBody
	GetApplicationCustomRuleIds() []*string
	SetRequestId(v string) *CreateApplicationCustomRulesResponseBody
	GetRequestId() *string
}

type CreateApplicationCustomRulesResponseBody struct {
	// The list of IDs for the newly created custom application rules. The order of the IDs corresponds to the order of the rules in the Rules request parameter.
	ApplicationCustomRuleIds []*string `json:"ApplicationCustomRuleIds,omitempty" xml:"ApplicationCustomRuleIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationCustomRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationCustomRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationCustomRulesResponseBody) GetApplicationCustomRuleIds() []*string {
	return s.ApplicationCustomRuleIds
}

func (s *CreateApplicationCustomRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationCustomRulesResponseBody) SetApplicationCustomRuleIds(v []*string) *CreateApplicationCustomRulesResponseBody {
	s.ApplicationCustomRuleIds = v
	return s
}

func (s *CreateApplicationCustomRulesResponseBody) SetRequestId(v string) *CreateApplicationCustomRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationCustomRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
