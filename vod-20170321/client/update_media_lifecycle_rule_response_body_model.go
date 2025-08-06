// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenRuleIds(v []*string) *UpdateMediaLifecycleRuleResponseBody
	GetForbiddenRuleIds() []*string
	SetNonExistRuleIds(v []*string) *UpdateMediaLifecycleRuleResponseBody
	GetNonExistRuleIds() []*string
	SetRequestId(v string) *UpdateMediaLifecycleRuleResponseBody
	GetRequestId() *string
}

type UpdateMediaLifecycleRuleResponseBody struct {
	ForbiddenRuleIds []*string `json:"ForbiddenRuleIds,omitempty" xml:"ForbiddenRuleIds,omitempty" type:"Repeated"`
	NonExistRuleIds  []*string `json:"NonExistRuleIds,omitempty" xml:"NonExistRuleIds,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaLifecycleRuleResponseBody) GetForbiddenRuleIds() []*string {
	return s.ForbiddenRuleIds
}

func (s *UpdateMediaLifecycleRuleResponseBody) GetNonExistRuleIds() []*string {
	return s.NonExistRuleIds
}

func (s *UpdateMediaLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaLifecycleRuleResponseBody) SetForbiddenRuleIds(v []*string) *UpdateMediaLifecycleRuleResponseBody {
	s.ForbiddenRuleIds = v
	return s
}

func (s *UpdateMediaLifecycleRuleResponseBody) SetNonExistRuleIds(v []*string) *UpdateMediaLifecycleRuleResponseBody {
	s.NonExistRuleIds = v
	return s
}

func (s *UpdateMediaLifecycleRuleResponseBody) SetRequestId(v string) *UpdateMediaLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
