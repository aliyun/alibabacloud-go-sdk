// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *DeleteAlertRuleRequest
	GetAlertId() *int64
}

type DeleteAlertRuleRequest struct {
	// The alert rule ID.
	//
	// For more information about how to obtain the ID of an alert rule, see [GetAlertRules](https://help.aliyun.com/document_detail/2612348.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s DeleteAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *DeleteAlertRuleRequest) SetAlertId(v int64) *DeleteAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *DeleteAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
