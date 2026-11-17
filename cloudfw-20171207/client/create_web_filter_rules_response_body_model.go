// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFilterRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWebFilterRulesResponseBody
	GetRequestId() *string
	SetWebFilterRuleIds(v []*string) *CreateWebFilterRulesResponseBody
	GetWebFilterRuleIds() []*string
}

type CreateWebFilterRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of IDs of the newly created web filtering rules. The order of the IDs corresponds one-to-one with the order of the input Rules parameter.
	WebFilterRuleIds []*string `json:"WebFilterRuleIds,omitempty" xml:"WebFilterRuleIds,omitempty" type:"Repeated"`
}

func (s CreateWebFilterRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebFilterRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWebFilterRulesResponseBody) GetWebFilterRuleIds() []*string {
	return s.WebFilterRuleIds
}

func (s *CreateWebFilterRulesResponseBody) SetRequestId(v string) *CreateWebFilterRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebFilterRulesResponseBody) SetWebFilterRuleIds(v []*string) *CreateWebFilterRulesResponseBody {
	s.WebFilterRuleIds = v
	return s
}

func (s *CreateWebFilterRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
