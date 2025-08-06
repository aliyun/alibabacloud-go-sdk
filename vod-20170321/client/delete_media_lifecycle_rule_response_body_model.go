// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenRuleIds(v []*string) *DeleteMediaLifecycleRuleResponseBody
	GetForbiddenRuleIds() []*string
	SetNonExistRuleIds(v []*string) *DeleteMediaLifecycleRuleResponseBody
	GetNonExistRuleIds() []*string
	SetRequestId(v string) *DeleteMediaLifecycleRuleResponseBody
	GetRequestId() *string
}

type DeleteMediaLifecycleRuleResponseBody struct {
	ForbiddenRuleIds []*string `json:"ForbiddenRuleIds,omitempty" xml:"ForbiddenRuleIds,omitempty" type:"Repeated"`
	NonExistRuleIds  []*string `json:"NonExistRuleIds,omitempty" xml:"NonExistRuleIds,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaLifecycleRuleResponseBody) GetForbiddenRuleIds() []*string {
	return s.ForbiddenRuleIds
}

func (s *DeleteMediaLifecycleRuleResponseBody) GetNonExistRuleIds() []*string {
	return s.NonExistRuleIds
}

func (s *DeleteMediaLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaLifecycleRuleResponseBody) SetForbiddenRuleIds(v []*string) *DeleteMediaLifecycleRuleResponseBody {
	s.ForbiddenRuleIds = v
	return s
}

func (s *DeleteMediaLifecycleRuleResponseBody) SetNonExistRuleIds(v []*string) *DeleteMediaLifecycleRuleResponseBody {
	s.NonExistRuleIds = v
	return s
}

func (s *DeleteMediaLifecycleRuleResponseBody) SetRequestId(v string) *DeleteMediaLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
