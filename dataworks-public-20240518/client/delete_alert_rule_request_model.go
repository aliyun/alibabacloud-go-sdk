// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteAlertRuleRequest
	GetId() *int64
}

type DeleteAlertRuleRequest struct {
	// The rule ID.
	//
	// example:
	//
	// 105412
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteAlertRuleRequest) SetId(v int64) *DeleteAlertRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
