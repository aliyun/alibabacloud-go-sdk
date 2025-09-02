// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDIAlarmRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDIAlarmRuleResponseBody
	GetRequestId() *string
}

type DeleteDIAlarmRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDIAlarmRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDIAlarmRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDIAlarmRuleResponseBody) SetRequestId(v string) *DeleteDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDIAlarmRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
