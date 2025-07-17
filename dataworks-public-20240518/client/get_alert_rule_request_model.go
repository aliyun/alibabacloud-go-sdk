// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAlertRuleRequest
	GetId() *string
}

type GetAlertRuleRequest struct {
	// The rule ID.
	//
	// example:
	//
	// 28547072
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAlertRuleRequest) GetId() *string {
	return s.Id
}

func (s *GetAlertRuleRequest) SetId(v string) *GetAlertRuleRequest {
	s.Id = &v
	return s
}

func (s *GetAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
