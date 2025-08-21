// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSchedulerRuleResponseBody
	GetRequestId() *string
}

type DeleteSchedulerRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchedulerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchedulerRuleResponseBody) SetRequestId(v string) *DeleteSchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchedulerRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
